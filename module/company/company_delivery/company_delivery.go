package company_delivery

import (
    "github.com/wirnat/axara/example/infrastructure/contextor"
    "github.com/wirnat/axara/module/company/company_usecase"
    "github.com/wirnat/axara/module/company/repository/company_repository"
    "github.com/wirnat/axara/module/company/request/company_request"
)

type CompanyRest struct {
	CompanyInteractor company_usecase.CompanyUsecase
	CompanyFetchRepo  company_repository.CompanyFetch
	CompanyGetRepo    company_repository.CompanyGet
	CompanyDeleteRepo company_repository.CompanyDelete
	CompanyPaginate   company_repository.CompanyPaginate
}

func NewCompanyRest() *CompanyRest {
	return &CompanyRest{}
}

func (r CompanyRest) Show(ctx *contextor.Contextor) error {
	param := company_request.CompanyParam{}
	err := ctx.BindQuery(&param)
	if err != nil {
		return err
	}

	res, err := r.CompanyGetRepo.Get(ctx.ToContext(), param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	return ctx.JSON(200, res)
}

func (r CompanyRest) Find(ctx *contextor.Contextor) error {
	param := company_request.CompanyParam{}
	err := ctx.BindQuery(&param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	res, err := r.CompanyFetchRepo.Fetch(ctx.ToContext(), param)
	if err != nil {
		return ctx.JSON(404, err.Error())
	}

	return ctx.JSON(200, res)
}

func (r CompanyRest) Store(ctx *contextor.Contextor) error {
	param := new(company_request.CompanyStore)
	err := ctx.Bind(param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	err = ctx.Validate(param)
	if err != nil {
		return err
	}

	res, err := r.CompanyInteractor.Store(ctx.ToContext(), *param)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, res)
}

func (r CompanyRest) Update(ctx *contextor.Contextor) error {
	param := new(company_request.CompanyUpdate)
	err := ctx.Bind(param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	param.UUID = ctx.Param("uuid")

	err = ctx.Validate(param)
	if err != nil {
		return err
	}

	res, err := r.CompanyInteractor.Update(ctx.ToContext(), *param)
	if err != nil {
		return err
	}

	return ctx.JSON(200, res)
}

func (r CompanyRest) Delete(ctx *contextor.Contextor) error {
	uuid := ctx.Param("uuid")
	err := r.CompanyDeleteRepo.Delete(ctx.ToContext(), uuid)
	if err != nil {
		return err
	}

	return ctx.JSON(200, uuid)
}

func (r CompanyRest) Paginate(ctx *contextor.Contextor) error {
	param := company_request.CompanyParam{}
	err := ctx.Bind(&param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	res, err := r.CompanyPaginate.Paginate(ctx.ToContext(), param)
	if err != nil {
		return ctx.JSON(404, err.Error())
	}

	return ctx.JSON(200, res)
}