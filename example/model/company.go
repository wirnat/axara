package model

type Company struct {
	BaseModel
	Name string `json:"name"` //@meta validate_store:true
}

/*
	~model~ Company
	~module~ company
	~end_model~
*/
