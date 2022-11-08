package model

type Branch struct {
	BaseModel
	CompanyID   int64   `json:"company_id"`
	Name        string  `json:"name"` //@meta validate_store:true
	Description *string `json:"description"`
}

/*
	~model~ Branch
	~module branch
	~end_model~
*/

type CheckDoubled struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
