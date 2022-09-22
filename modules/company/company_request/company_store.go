package company_request

type CompanyStore struct {
    ID int64 `json:"id" `
    UUID string `json:"uuid" `
    CreatedAt time.Time `json:"created_at" `
    UpdatedAt time.Time `json:"updated_at" `
    DeletedAt *time.Time `json:"deleted_at" `
    Name string `json:"name" validate:"required"`
    Phone string `json:"phone" validate:"required"`
    Email string `json:"email" validate:"required"`
    Description *string `json:"description" `
    LogoID *int64 `json:"logo_id" `
    Latitude *float64 `json:"latitude" `
    Longitude *float64 `json:"longitude" `
}