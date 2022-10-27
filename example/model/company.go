package model

type Company struct {
	BaseModel
	Name string `json:"name"`
}

/*
	@model Company
	@service company
	@snakecase company
*/
