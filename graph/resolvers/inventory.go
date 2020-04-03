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
	err := m.Db.Create(&inventory).Error
	return &inventory, err
}

func (m *mutationResolver) AddInventoryItem(ctx context.Context, input *model.NewInventoryItem) (*models.InventoryItem, error) {
	inventoryItem := models.InventoryItem{
		InventoryID:	input.InventoryID,
		Quantity:		input.Quantity,
		ItemID:     	input.ItemID,
	}

	err := m.Db.Transaction(func(tx *gorm.DB) error {
		if e := tx.Create(&inventoryItem).Error; e != nil {
			log.Println("error creating `InventoryItem`")
			return e
		}
		e := m.Db.Where("id = ?", input.ItemID).Find(&inventoryItem.Item).Error
		return e
	})
	return &inventoryItem, err
}

func (m *mutationResolver) RemoveInventoryItem(ctx context.Context, id *string) (string, error) {
	err := m.Db.Delete(&models.InventoryItem{}, id).Error
	return *id, err
}

func (m *mutationResolver) DeleteInventory(ctx context.Context, id *string) (*models.Inventory, error) {
	var inventory models.Inventory
	err := m.Db.Delete(&inventory, id).Error
	return &inventory, err
}

