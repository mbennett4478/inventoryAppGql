package transformations

import (
	"github.com/gofrs/uuid"
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
		BaseModelSoftDelete: models.BaseModelSoftDelete{
			BaseModel: models.BaseModel{
				ID: uuid.FromStringOrNil(gqlIn.ID),
				CreatedAt: gqlIn.CreatedAt,
				UpdatedAt: gqlIn.UpdatedAt,
			},
		},
		Name:  gqlIn.Name,
		Items: GqlInventoryItemsToDbInventoryItems(gqlIn.Items),
	}
}

func GqlInventoryItemsToDbInventoryItems(gqlIn []model.InventoryItem) []models.InventoryItem {
	var dbItems []models.InventoryItem
	for _, item := range gqlIn {
		dbItems = append(dbItems, GqlInventoryItemToDbInventoryItem(&item))
	}
	return dbItems
}

func GqlInventoryItemToDbInventoryItem(gqlIn *model.InventoryItem) models.InventoryItem {
	return models.InventoryItem{
		BaseModelSoftDelete: models.BaseModelSoftDelete{
			BaseModel: models.BaseModel{
				ID: uuid.FromStringOrNil(gqlIn.ID),
			},
		},
		InventoryID: uuid.FromStringOrNil(gqlIn.InventoryID),
		Quantity: gqlIn.Quantity,
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
	return model.InventoryItem{}
}
