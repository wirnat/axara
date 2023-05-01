package company_request
import(
    "time"
)

type CompanyStore struct {
	  ID bigint `json:"id" `
	  UUID string `json:"uuid" `
	  createdAt string `json:"created_at" `
	  createdBy string `json:"created_by" `
	  updatedAt string `json:"updated_at" `
	  updatedBy string `json:"updated_by" `
	  deletedAt string `json:"deleted_at" `
	  deletedBy string `json:"deleted_by" `
	  name string `json:"name" `
	  parent string `json:"parent" `
	  parentUUID string `json:"parent_uuid" `
	  subCompany CompanyModel `json:"sub_company" `
	  type string `json:"type" `
}
