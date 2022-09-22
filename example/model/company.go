package model

type Company struct {
	BaseModel
	Name        string   `json:"name"`  //@meta validate_store:true
	Phone       string   `json:"phone"` //@meta validate_store:true
	Email       string   `json:"email"` //@meta validate_store:true
	Description *string  `json:"description"`
	LogoID      *int64   `json:"logo_id"` //MEDIA
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
}

/*
	@model Company
	@service company
*/
