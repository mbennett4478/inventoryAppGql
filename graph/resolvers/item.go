package resolvers

import (
	"context"
	"github.com/jinzhu/gorm"
	"inventoryGql/graph/model"
	"log"

	//"log"
)

func (m *mutationResolver) CreateItem(ctx context.Context, input *model.NewItem) (*model.Item, error) {
	item := model.Item{
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

func (q *queryResolver) Item(ctx context.Context, id *string) (*model.Item, error) {
	var item = model.Item{}
	var items []model.Item
	q.Db.Find(&items)
	err := q.Db.Where("id = ?", id).Find(&item).Error
	if err != nil {
		log.Println(err)
	}
	return &item, nil
}
