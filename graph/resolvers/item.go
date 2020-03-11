package resolvers

import (
	"context"
	"github.com/jinzhu/gorm"
	"inventoryGql/graph/model"
	"inventoryGql/graph/models"
	"log"

	//"log"
)

func (m *mutationResolver) CreateItem(ctx context.Context, input *model.NewItem) (*models.Item, error) {
	item := models.Item{
		Name: input.Name,
	}

	err := m.Db.Transaction(func(tx *gorm.DB) error {
		if e := tx.Create(&item).Error; e != nil {
			return e
		}
		return nil
	})

	return &item, err
}

func (q *queryResolver) Item(ctx context.Context, id *int) (*models.Item, error) {
	var item = models.Item{}
	var items []models.Item
	q.Db.Find(&items)
	err := q.Db.Where("id = ?", *id).Find(&item).Error
	if err != nil {
		log.Println(err)
	}
	return &item, nil
}
