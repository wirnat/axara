package branch_request

type BranchParam struct {
    ID *int64 `json:"id"`
    UUID *string `json:"uuid"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at"`
    CompanyID *int64 `json:"company_id"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Email *string `json:"email"`
    Phone *string `json:"phone"`
    PicName *string `json:"pic_name"`
    PicPhone *string `json:"pic_phone"`
    PicEmail *string `json:"pic_email"`
    Address *string `json:"address"`
    Status *string `json:"status"`
    VerifiedStatus *string `json:"verified_status"`
    OpenStatus *string `json:"open_status"`
    ProfileImageID *int64 `json:"profile_image_id"`
    OpenedAt *time.Time `json:"opened_at"`
    ClosedAt *time.Time `json:"closed_at"`
    Latitude *float64 `json:"latitude"`
    Longitude *float64 `json:"longitude"`
}