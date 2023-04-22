package main

import (
	"backend/api"
	"backend/configs"
)

func main() {
	configs.DB.GetConnection()
	api.Run()
}
