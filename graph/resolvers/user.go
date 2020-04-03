package resolvers

import (
	"context"
	"errors"
	"inventoryGql/graph/model"
	"inventoryGql/graph/models"
)

func (m *mutationResolver) RegisterUser(ctx context.Context, input *model.NewUser) (bool, error) {
	if input.Passord != input.ConfirmationPassword {
		return false, errors.New("passwords don't match")
	}

	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Passord,
	}

	err := m.Db.Create(&user).Error

	return err == nil, err
}