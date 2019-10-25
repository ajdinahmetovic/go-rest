package es

import (
	"context"

	"gopkg.in/olivere/elastic.v5"
)

//GetItem func resturns items from es
func GetItem(ctx context.Context, field string, term string) (res *elastic.SearchResult, err error) {

	var query elastic.Query
	if len(term) < 1 {
		query = elastic.MatchAllQuery{}
	} else {
		query = elastic.NewMatchPhrasePrefixQuery(field, term)
	}

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
