package resolvers

import (
	"context"
	"github.com/jinzhu/gorm"
	"inventoryGql/graph/model"
	"inventoryGql/graph/models"
	"log"
)

func (q *queryResolver) Inventories(ctx context.Context) ([]models.Inventory, error) {
	var inventories []models.Inventory
	q.Db.Set("gorm:auto_preload", true).Find(&inventories)
	return inventories, nil
}

func (m *mutationResolver) CreateInventory(ctx context.Context, input *model.NewInventory) (*models.Inventory, error) {
	inventory := models.Inventory {
		Name: input.Name,
	}
	err := m.Db.Transaction(func(tx *gorm.DB) error {
		if e := tx.Create(&inventory).Error; e != nil {
			log.Println("error creating `Inventory`")
			return e
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (m *mutationResolver) AddInventoryItem(ctx context.Context, input *model.NewInventoryItem) (*models.InventoryItem, error) {
	item := models.Item{}
	m.Db.Where("id = ?", input.ItemID).Find(&item)
	inventoryItem := models.InventoryItem {
		Quantity: input.Quantity,
		Item: item,
		InventoryID: input.InventoryID,
	}

	err := m.Db.Transaction(func(tx *gorm.DB) error {
		if e := tx.Preload("Item").Create(&inventoryItem).Error; e != nil {
			log.Println("error creating `InventoryItem`")
			return e
		}
		return nil
	})
	//item := models.Item{}
	//m.Db.Find(&item).Where("id = ?", input.ItemID)
	//inventoryItem.Item = &item

	return &inventoryItem, err
}

func (m *mutationResolver) RemoveInventoryItem(ctx context.Context, id *int) (*models.InventoryItem, error) {
	panic("implement me")
}

func (m *mutationResolver) DeleteInventory(ctx context.Context, id *int) (*models.Inventory, error) {
	inventory := models.Inventory{ID: *id}
	err := m.Db.Transaction(func(tx *gorm.DB) error {
		if e := tx.Delete(&inventory).Error; e != nil {
			log.Println("error deleting `Inventory`")
			return e
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

