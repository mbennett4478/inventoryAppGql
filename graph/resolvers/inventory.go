package resolvers

import (
	"context"
	"github.com/jinzhu/gorm"
	"inventoryGql/graph/model"
	"log"
)

func (q *queryResolver) Inventories(ctx context.Context) ([]model.Inventory, error) {
	var inventories []model.Inventory
	q.Db.Set("gorm:auto_preload", true).Find(&inventories)
	return inventories, nil
}

func (m *mutationResolver) CreateInventory(ctx context.Context, input *model.NewInventory) (*model.Inventory, error) {
	inventory := model.Inventory {
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

func (m *mutationResolver) AddInventoryItem(ctx context.Context, input *model.NewInventoryItem) (*model.InventoryItem, error) {
	item := model.Item{}
	m.Db.Where("id = ?", input.ItemID).Find(&item)
	inventoryItem := model.InventoryItem {
		Quantity: input.Quantity,
		Item: &item,
		InventoryID: input.InventoryID,
	}

	err := m.Db.Transaction(func(tx *gorm.DB) error {
		if e := tx.Create(&inventoryItem).Error; e != nil {
			log.Println("error creating `InventoryItem`")
			return e
		}
		return nil
	})

	return &inventoryItem, err
}

func (m *mutationResolver) RemoveInventoryItem(ctx context.Context, id *string) (string, error) {
	err := m.Db.Transaction(func(tx *gorm.DB) error {
		if e := tx.Delete(&model.InventoryItem{}, id).Error; e != nil {
			log.Println("error deleting `InventoryItem`")
			return e
		}
		return nil
	})
	return *id, err
}

func (m *mutationResolver) DeleteInventory(ctx context.Context, id *string) (*model.Inventory, error) {
	inventory := model.Inventory{}
	err := m.Db.Transaction(func(tx *gorm.DB) error {
		if e := tx.Delete(&inventory, id).Error; e != nil {
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

