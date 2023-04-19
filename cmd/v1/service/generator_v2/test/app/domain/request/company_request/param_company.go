package company_request

import (
    "time"
"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/infrastructure/paginator"
)

type CompanyParam struct {
    ID *int64 `json:"id" form:"id" query:"id"`
    UUID *string `json:"uuid" form:"uuid" query:"uuid"`
    CreatedAt *time.Time `json:"created_at" form:"created_at" query:"created_at"`
    UpdatedAt *time.Time `json:"updated_at" form:"updated_at" query:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at" form:"deleted_at" query:"deleted_at"`
    Name *string `json:"name" form:"name" query:"name"`
    paginator.PaginateReq
}
