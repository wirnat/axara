package model

type Company struct {
	BaseModel
	Name string `json:"name"` //@meta validate_store:true
}

//@Register Company

type XX struct {
	BaseModel
	Ix int `json:"ix"`
}
