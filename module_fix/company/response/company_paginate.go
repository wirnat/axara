package response

import (
	"github.com/wirnat/aksara-cli/example/infrastructure/paginator"
	"github.com/wirnat/aksara-cli/example/model"
)

type CompanyPaginate struct {
	Pagination *paginator.Pagination `json:"paginator,omitempty"`
	Result     []model.Company       `json:"result"`
}
