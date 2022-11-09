package branch_request

type BranchStore struct {
	  ID int64 `json:"id" `
	  UUID string `json:"uuid" `
	  CompanyID int64 `json:"company_id" `
	  Name string `json:"name" validate:"required"`
	  Description *string `json:"description" `
}