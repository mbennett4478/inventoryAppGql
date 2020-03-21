package transformations

import (
	"inventoryGql/graph/model"
	"inventoryGql/graph/models"
)

func DbInventoryToGqlInventory(dbIn *models.Inventory) *model.Inventory {
	return &model.Inventory{
		ID: dbIn.ID.String(),
		CreatedAt: dbIn.CreatedAt,
		UpdatedAt: dbIn.UpdatedAt,
		Items: DbInventoryItemsToGqlInventoryItems(dbIn.Items),
	}
}

func GqlInventoryToDbInventory(gqlIn *model.Inventory) *models.Inventory {
	return &models.Inventory{
		Name:                "",
		Items:               nil,
	}
}

func DbInventoryItemsToGqlInventoryItems(dbItems []models.InventoryItem) []model.InventoryItem{
	var gqlItems []model.InventoryItem
	for _, item := range dbItems {
		gqlItems = append(gqlItems, DbInventoryItemToGqlInventoryItem(&item))
	}
	return gqlItems
}

func DbInventoryItemToGqlInventoryItem(dbItem *models.InventoryItem) model.InventoryItem {
	return model.Item{}
}
