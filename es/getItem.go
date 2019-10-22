package es

import (
	"context"

	"github.com/ajdinahmetovic/go-rest/logger"

	"gopkg.in/olivere/elastic.v5"
)

//GetItem func resturns items from es
func GetItem(ctx context.Context, terms ...string) (res *elastic.SearchResult, err error) {

	//termQuery := elastic.NewTermQuery("John", "")
	//Execute search on elasticsearch with options

	query := elastic.NewMatchQuery("title", "Hello")
	logger.Info("Query started", "query", *query)
	searchResult, err := elasticClient.Search().
		Index("items").
		Query(query).
		Pretty(true).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}
