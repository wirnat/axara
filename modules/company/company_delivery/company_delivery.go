package company_rest

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wirawirw/aksara-cli/modules/company/company_repository"
	"gitlab.com/wirawirw/aksara-cli/modules/company/company_request"
	"gitlab.com/wirawirw/aksara-cli/modules/company/company_usecase"
	"gitlab.com/wirawirw/aksara-cli/util/query_reader"
)

type CompanyRest struct {
	CompanyInteractor company_usecase.CompanyUsecase
	CompanyFetchRepo  company_repository.CompanyFetch
	CompanyGetRepo    company_repository.CompanyGet
	CompanyDeleteRepo company_repository.CompanyDelete
}

func (r CompanyRest) Show(ctx echo.Context) error {
	param := company_request.CompanyParam{}
	_uuid := ctx.Param("uuid")
	param.UUID = &_uuid

	res, err := r.CompanyGetRepo.Get(ctx.Request().Context(), param)
	if err != nil {
		return err
	}

	return ctx.JSON(200, res)
}

func (r CompanyRest) Find(ctx echo.Context) error {
	param := company_request.CompanyParam{}
	err := query_reader.Bind(ctx, &param)
	if err != nil {
		return err
	}

	res, err := r.CompanyFetchRepo.Fetch(ctx.Request().Context(), param)
	if err != nil {
		return err
	}

	return ctx.JSON(200, res)
}

func (r CompanyRest) Store(ctx echo.Context) error {
	param := new(company_request.CompanyStore)
	err := ctx.Bind(param)
	if err != nil {
		return err
	}

	err = ctx.Validate(param)
	if err != nil {
		return err
	}

	res, err := r.CompanyInteractor.Store(ctx.Request().Context(), *param)
	if err != nil {
		return err
	}
	return ctx.JSON(200, res)
}
