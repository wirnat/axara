package company_request

import (
    "time"
)

type CompanyUpdate struct {
    ID int64 `json:"id" `
    UUID string `json:"uuid" validate:"required"`
    CreatedAt time.Time `json:"created_at" `
    UpdatedAt time.Time `json:"updated_at" `
    DeletedAt time.Time `json:"deleted_at" `
    Name string `json:"name" `
}