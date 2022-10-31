package company_request

type CompanyStore struct {
	ID   int64  `json:"id" `
	UUID string `json:"uuid" `
	Name string `json:"name" validate:"required"`
}
