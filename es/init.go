package es

import (
	"context"

	"github.com/ajdinahmetovic/go-rest/logger"
	"gopkg.in/olivere/elastic.v5"
)

const url = "http://127.0.0.1:9200"
const itemMapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"item":{
			"properties":{
				"id":{
					"type":"integer"
				},
				"title":{
					"type":"text"
				},
				"description":{
					"type":"text"
				},
				"userid":{
					"type":"integer"
				}
			}
		}
	}
}`

var elasticClient *elastic.Client

//Init func, inits elasticsearch
func Init(ctx context.Context) error {
	var err error
	// init Elastic client
	elasticClient, err = elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	//Ping elasticsearch
	result, status, err := elasticClient.Ping(url).Do(ctx)
	if err != nil {
		return err
	}
	logger.Info("Elasticsearch inited", "Result", result, "code", status)

	//Check for item index
	exists, err := elasticClient.IndexExists("items").Do(ctx)
	if err != nil {
		return err
	}

	//If index is not present create new one
	if !exists {
		logger.Info("Index not present creating new one")
		_, err := elasticClient.CreateIndex("items").BodyString(itemMapping).Do(ctx)
		if err != nil {
			logger.Error("Failed to create elastic index", "err", err)
			return err
		}
	}

	return nil
}
