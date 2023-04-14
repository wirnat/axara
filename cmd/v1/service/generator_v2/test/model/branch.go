package model

type Branch struct {
	BaseModel
	CompanyID   int64   `json:"company_id"`
	Name        string  `json:"name"` //@meta validate_store:true
	Description *string `json:"description"`
}

//@Register Branch
