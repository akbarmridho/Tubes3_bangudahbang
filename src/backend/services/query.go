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
	similarity float32
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
	var match models.Query
	i := 0
	for i < len(queries) && match != (models.Query{}) {
		query := queries[i]
		var matchIdxs []int

		if isKMP {
			matchIdxs = stringmatcher.KMP(input, query.Query)
		} else {
			matchIdxs = stringmatcher.BM(input, query.Query)
		}

		if len(matchIdxs) != 0 {
			match = query
		}
		i++
	}

	var closestSimilar []models.Query

	// if not found, find by similarity
	if match == (models.Query{}) {
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

	if match != (models.Query{}) {
		return match.Response, nil
	}

	response := "Pertanyaan tidak ditemukan di database\n"

	if len(closestSimilar) > 0 {
		response = response + "Apakah maksud anda\n"

		for idx, similar := range closestSimilar {
			response = response + strconv.Itoa(idx+1) + ". " + similar.Query + "\n"
		}
	}

	return response, nil
}

func DeleteQuery(input string) (string, error) {
	if err := refreshQuery(); err != nil {
		return "", err
	}

	// find first exact match
	var match models.Query
	i := 0
	for i < len(queries) && match != (models.Query{}) {
		query := queries[i]
		var matchIdxs []int

		matchIdxs = stringmatcher.KMP(input, query.Query)

		if len(matchIdxs) != 0 {
			match = query
		}
		i++
	}

	if match == (models.Query{}) {
		return "Tidak ada pertanyaan " + input + " pada database!", nil
	}

	db := configs.DB.GetConnection()

	lock.Lock()

	if err := db.Delete(&models.Query{}, match.ID); err != nil {
		lock.Unlock()
		return "Tidak dapat menghapus query", nil
	}

	isDirty = true
	lock.Unlock()

	return "Pertanyaan " + match.Query + " telah dihapus", nil
}

func AddQuery(input string, answer string) (string, error) {
	return "Not implemented yet", nil
}
