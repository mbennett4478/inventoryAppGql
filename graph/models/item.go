package models

type Item struct {
	BaseModelSoftDelete
	Name      string `json:"name"`
}
