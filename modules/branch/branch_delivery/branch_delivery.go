package branch_rest

import (
	"github.com/labstack/echo/v4"
	"github.com/wirnat/aksara-cli/modules/branch/branch_repository"
	"github.com/wirnat/aksara-cli/modules/branch/branch_request"
	"github.com/wirnat/aksara-cli/modules/branch/branch_usecase"
	"github.com/wirnat/aksara-cli/util/query_reader"
)

type BranchRest struct {
	BranchInteractor branch_usecase.BranchUsecase
	BranchFetchRepo  branch_repository.BranchFetch
	BranchGetRepo    branch_repository.BranchGet
	BranchDeleteRepo branch_repository.BranchDelete
}

func (r BranchRest) Show(ctx echo.Context) error {
	param := branch_request.BranchParam{}
	_uuid := ctx.Param("uuid")
	param.UUID = &_uuid

	res, err := r.BranchGetRepo.Get(ctx.Request().Context(), param)
	if err != nil {
		return err
	}

	return ctx.JSON(200, res)
}

func (r BranchRest) Find(ctx echo.Context) error {
	param := branch_request.BranchParam{}
	err := query_reader.Bind(ctx, &param)
	if err != nil {
		return err
	}

	res, err := r.BranchFetchRepo.Fetch(ctx.Request().Context(), param)
	if err != nil {
		return err
	}

	return ctx.JSON(200, res)
}

func (r BranchRest) Store(ctx echo.Context) error {
	param := new(branch_request.BranchStore)
	err := ctx.Bind(param)
	if err != nil {
		return err
	}

	err = ctx.Validate(param)
	if err != nil {
		return err
	}

	res, err := r.BranchInteractor.Store(ctx.Request().Context(), *param)
	if err != nil {
		return err
	}
	return ctx.JSON(200, res)
}
