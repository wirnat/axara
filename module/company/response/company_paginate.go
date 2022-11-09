package response

import (
	"github.com/wirnat/axara/example/model"
	"github.com/wirnat/axara/example/infrastructure/paginator"
)

type CompanyPaginate struct {
	Pagination *paginator.Pagination `json:"paginator,omitempty"`
	Result     []model.Company       `json:"result"`
}
