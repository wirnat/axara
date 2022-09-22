package branch_usecase_v1

import (
	"context"
	"github.com/wirnat/aksara-cli/example/model"
	"github.com/wirnat/aksara-cli/modules/branch/branch_repository"
	"github.com/wirnat/aksara-cli/modules/branch/branch_request"
)

type branchUsecase struct {
	branchStore  branch_repository.BranchStore
	branchUpdate branch_repository.BranchUpdate
}

func NewBranchUsecase(branchStore branch_repository.BranchStore, branchUpdate branch_repository.BranchUpdate) *branchUsecase {
	return &branchUsecase{branchStore: branchStore, branchUpdate: branchUpdate}
}

func (u branchUsecase) Store(ctx context.Context, param branch_request.BranchStore) (r model.Branch, err error) {
	r.ID = param.ID
	r.UUID = param.UUID
	r.CreatedAt = param.CreatedAt
	r.UpdatedAt = param.UpdatedAt
	r.DeletedAt = param.DeletedAt
	r.CompanyID = param.CompanyID
	r.Name = param.Name
	r.Description = param.Description
	r.Email = param.Email
	r.Phone = param.Phone
	r.PicName = param.PicName
	r.PicPhone = param.PicPhone
	r.PicEmail = param.PicEmail
	r.Address = param.Address
	r.Status = param.Status
	r.VerifiedStatus = param.VerifiedStatus
	r.OpenStatus = param.OpenStatus
	r.ProfileImageID = param.ProfileImageID
	r.OpenedAt = param.OpenedAt
	r.ClosedAt = param.ClosedAt
	r.Latitude = param.Latitude
	r.Longitude = param.Longitude

	err = u.branchStore.Store(ctx, &r)
	return
}

func (u branchUsecase) Update(ctx context.Context, param branch_request.BranchStore) (r model.Branch, err error) {
	r.ID = param.ID
	r.UUID = param.UUID
	r.CreatedAt = param.CreatedAt
	r.UpdatedAt = param.UpdatedAt
	r.DeletedAt = param.DeletedAt
	r.CompanyID = param.CompanyID
	r.Name = param.Name
	r.Description = param.Description
	r.Email = param.Email
	r.Phone = param.Phone
	r.PicName = param.PicName
	r.PicPhone = param.PicPhone
	r.PicEmail = param.PicEmail
	r.Address = param.Address
	r.Status = param.Status
	r.VerifiedStatus = param.VerifiedStatus
	r.OpenStatus = param.OpenStatus
	r.ProfileImageID = param.ProfileImageID
	r.OpenedAt = param.OpenedAt
	r.ClosedAt = param.ClosedAt
	r.Latitude = param.Latitude
	r.Longitude = param.Longitude

	err = u.branchStore.Store(ctx, &r)
	return
}
