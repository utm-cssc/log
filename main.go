package main

import (
	"github.com/utm-cssc/log/app"
	"github.com/utm-cssc/log/config"
)

func main() {
	dbConfig := config.GetDBConfig()
	logger := app.App{}
	logger.Init(dbConfig)
	logger.Run(":8080")
}
