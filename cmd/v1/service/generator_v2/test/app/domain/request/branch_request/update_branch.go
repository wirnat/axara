package branch_request

import (
    "time"
)

type BranchUpdate struct {
    ID int64 `json:"id" `
    UUID string `json:"uuid" validate:"required"`
    CreatedAt time.Time `json:"created_at" `
    UpdatedAt time.Time `json:"updated_at" `
    DeletedAt time.Time `json:"deleted_at" `
    CompanyID int64 `json:"company_id" `
    Name string `json:"name" `
    Description string `json:"description" `
}
