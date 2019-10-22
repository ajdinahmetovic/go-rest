package es

import (
	"context"
	"strconv"

	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/item-service/proto/v1"
)

//InsertItem func inserts data to es
func InsertItem(ctx context.Context, item *proto.Item) error {

	req := db.Item{
		ID:          int(item.ID),
		Title:       item.Title,
		Description: item.Description,
		UserID:      int(item.UserID),
	}

	//Save item to es Index
	_, err := elasticClient.Index().Index("items").Type("item").Id(strconv.Itoa(int(item.ID))).BodyJson(req).Do(ctx)
	if err != nil {
		return err
	}

	elasticClient.Flush().Index("items").Do(ctx)
	return nil
}
