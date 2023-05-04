package services

import (
	"backend/configs"
	"backend/models"
	"backend/utils"
	stringmatcher "backend/utils/string-matcher"
	"sort"
	"strconv"
	"sync"
)

const SimilarityOffset = 0.9

var queries []models.Query
var isDirty = true
var lock sync.Mutex

type QuerySimilarity struct {
	idx        int
	similarity float64
}

func refreshQuery() error {
	lock.Lock()
	if isDirty {
		db := configs.DB.GetConnection()
		var result []models.Query
		if err := db.Find(&result).Error; err != nil {
			lock.Unlock()
			return err
		}
		queries = result
		isDirty = false
	}
	lock.Unlock()
	return nil
}

func MatchQuery(input string, isKMP bool) (string, error) {
	if err := refreshQuery(); err != nil {
		return "", err
	}
	// find first exact match
	match := models.Query{
		Response: "",
	}
	i := 0
	for i < len(queries) && match.Response == "" {
		query := queries[i]
		var matchIdxs []int = make([]int, 0)

		if isKMP {
			matchIdxs = stringmatcher.KMP(input, query.Query)
		} else {
			matchIdxs = stringmatcher.BM(input, query.Query)
		}

		if len(matchIdxs) > 0 {
			match = query
		}
		i++
	}

	var closestSimilar []models.Query

	// if not found, find by similarity
	if match.Response == "" {
		var querySimilarities []QuerySimilarity
		for idx, query := range queries {
			querySimilarities = append(querySimilarities, QuerySimilarity{
				idx:        idx,
				similarity: utils.MeasureSimilarity(input, query.Query),
			})
		}

		// sort
		sort.Slice(querySimilarities, func(i, j int) bool {
			return querySimilarities[i].similarity > querySimilarities[j].similarity
		})

		if querySimilarities[0].similarity >= SimilarityOffset {
			match = queries[querySimilarities[0].idx]
		} else {
			var closestSimilarLength int

			if len(querySimilarities) < 3 {
				closestSimilarLength = len(querySimilarities)
			} else {
				closestSimilarLength = 3
			}

			for j := 0; j < closestSimilarLength; j++ {
				closestSimilar = append(closestSimilar, queries[querySimilarities[j].idx])
			}
		}
	}

	if match.Response != "" {
		return match.Response, nil
	}

	response := "Question not found in database\n"

	if len(closestSimilar) > 0 {
		response = response + "Do you mean\n"

		for idx, similar := range closestSimilar {
			response = response + strconv.Itoa(idx+1) + ". " + similar.Query + "\n"
		}
	}

	return response, nil
}

func DeleteQuery(input string, isKMP bool) (string, error) {
	if err := refreshQuery(); err != nil {
		return "", err
	}

	// find first exact match
	match := models.Query{
		Response: "",
	}
	i := 0
	for i < len(queries) && match.Response == "" {
		query := queries[i]
		var matchIdxs []int = make([]int, 0)
		if isKMP {
			matchIdxs = stringmatcher.KMP(input, query.Query)
		} else {
			matchIdxs = stringmatcher.BM(input, query.Query)
		}

		if utils.Comparator(matchIdxs, query.Query, input) {
			match = query
		}
		i++
	}

	if match.Response == "" {
		return "The question " + input + " is not found in the database!", nil
	}

	db := configs.DB.GetConnection()

	lock.Lock()

	if err := db.Delete(&models.Query{}, match.ID).Error; err != nil {
		lock.Unlock()
		return "Failed to delete query", nil
	}

	isDirty = true
	lock.Unlock()

	return "The question " + match.Query + " has successfully been deleted", nil
}

func AddQuery(question string, answer string, isKMP bool) (string, error) {
	if err := refreshQuery(); err != nil {
		return "", err
	}

	// find first exact match
	match := models.Query{
		Response: "",
	}
	for i := 0; i < len(queries) && match.Response == ""; i++ {
		query := queries[i]
		var matchIdxs []int = make([]int, 0)
		if isKMP {
			matchIdxs = stringmatcher.KMP(question, query.Query)
		} else {
			matchIdxs = stringmatcher.BM(question, query.Query)
		}

		if utils.Comparator(matchIdxs, query.Query, question) {
			match = query
		}
	}

	db := configs.DB.GetConnection()

	lock.Lock()

	// if exists
	if match.Response != "" {
		// update the response
		if err := db.Model(&models.Query{}).Where("id = ?", match.ID).Update("response", answer).Error; err != nil {
			lock.Unlock()
			return "Failed to update the answer to that question!", nil
		}
	} else {
		// create new query and ans
		newQuery := &models.Query{
			Query:    question,
			Response: answer,
		}
		if err := db.Create(newQuery).Error; err != nil {
			return "Failed to add question!", nil
		}
	}

	isDirty = true
	lock.Unlock()

	return "Successfully added " + question + " to database with anwer " + answer, nil
}
