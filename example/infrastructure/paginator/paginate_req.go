package paginator

type PaginateReq struct {
	Limit  int    `json:"limit" form:"limit"`
	Page   int    `json:"page" form:"page"`
	Sort   string `json:"sort" form:"sort"`
	Search string `json:"search" form:"search"`
}
