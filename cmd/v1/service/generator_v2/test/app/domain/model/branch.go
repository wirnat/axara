package model

type Branch struct {
	BaseModel
	CompanyID   int64   `json:"company_id"` //@meta validate_store:true
	Name        string  `json:"name"`       //@meta validate_store:true
	Description *string `json:"description"`
}

//@Register Branch
