package response

import (
	"github.com/wirnat/aksara-cli/example/model"
    "github.com/wirnat/aksara-cli/example/infrastructure/paginator"
)

type BranchPaginate struct {
	Pagination *paginator.Pagination `json:"paginator,omitempty"`
	Result     []model.Branch      `json:"result"`
}
