package paginator

type PaginateReq struct {
	Limit  int    `json:"limit" form:"limit" query:"limit"`
	Page   int    `json:"page" form:"page" query:"page"`
	Sort   string `json:"sort" form:"sort" query:"sort"`
	Search string `json:"search" form:"search" query:"search"`
}

func Filter(limit, page *int, sort, direction *string) *Pagination {
	var p Pagination
	var withPage bool
	if limit != nil {
		p.PaginateReq.Limit = *limit
		withPage = true
	}
	if page != nil {
		p.PaginateReq.Page = *page
		withPage = true
	}
	if sort != nil {
		p.PaginateReq.Sort = *sort
		withPage = true
	}
	if direction != nil {
		p.PaginateReq.Sort = p.PaginateReq.Sort + " " + *direction
		withPage = true
	} else if sort != nil {
		p.PaginateReq.Sort = p.PaginateReq.Sort + " DESC"
	}

	if withPage {
		return &p
	}
	return nil
}
