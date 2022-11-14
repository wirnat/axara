package company_request

import (
    "time"
"github.com/wirnat/axara/example/infrastructure/paginator"
)

type CompanyParam struct {
    ID *int64 `json:"id"`
    UUID *string `json:"uuid"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at"`
    Name *string `json:"name"`
    paginator.PaginateReq
}