package branch_response

import (
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/model"
    "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/infrastructure/paginator"
)

type BranchPaginate struct {
	Pagination *paginator.Pagination `json:"pagination,omitempty"`
	Result     []model.Branch      `json:"result"`
}
