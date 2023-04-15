package branch_request
import(
    "time"
)

type BranchStore struct {
	  ID int64 `json:"id" `
	  UUID string `json:"uuid" `
	  CreatedAt time.Time `json:"created_at" `
	  UpdatedAt time.Time `json:"updated_at" `
	  DeletedAt *time.Time `json:"deleted_at" `
	  CompanyID int64 `json:"company_id" `
	  Name string `json:"name" validate:"required"`
	  Description *string `json:"description" `
}
