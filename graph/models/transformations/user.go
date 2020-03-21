package transformations

import (
	"inventoryGql/graph/model"
	"inventoryGql/graph/models"
)

func DbUserToGql(dbUser *models.User) *model.User {
	return &model.User{}
}

func GqlUserToDbUser(dbUser *model.User) *models.User {
	return &models.User{}
}
