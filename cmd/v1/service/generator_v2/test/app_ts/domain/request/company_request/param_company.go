package company_request

import (
    "time"
"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/infrastructure/paginator"
)

type CompanyParam struct {
    ID *bigint `json:"id" form:"id" query:"id"`
    UUID *string `json:"uuid" form:"uuid" query:"uuid"`
    createdAt *string `json:"created_at" form:"created_at" query:"created_at"`
    createdBy *string `json:"created_by" form:"created_by" query:"created_by"`
    updatedAt *string `json:"updated_at" form:"updated_at" query:"updated_at"`
    updatedBy *string `json:"updated_by" form:"updated_by" query:"updated_by"`
    deletedAt *string `json:"deleted_at" form:"deleted_at" query:"deleted_at"`
    deletedBy *string `json:"deleted_by" form:"deleted_by" query:"deleted_by"`
    name *string `json:"name" form:"name" query:"name"`
    parent *string `json:"parent" form:"parent" query:"parent"`
    parentUUID *string `json:"parent_uuid" form:"parent_uuid" query:"parent_uuid"`
    subCompany *CompanyModel `json:"sub_company" form:"sub_company" query:"sub_company"`
    type *string `json:"type" form:"type" query:"type"`
    paginator.PaginateReq
}
