package company_response

import (
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/domain/model"
    "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/infrastructure/paginator"
)

type CompanyPaginate struct {
	Pagination *paginator.Pagination `json:"pagination,omitempty"`
	Result     []model.Company      `json:"result"`
}
