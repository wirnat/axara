package company_http_1

import (
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/request/company_request"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/usecase/company_usecase"
	"github.com/labstack/echo/v4"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/infrastructure/utils"
)

type companyRest struct {
	CompanyFetch  company_usecase.CompanyFetch
	CompanyGet    company_usecase.CompanyGet
	CompanyDelete company_usecase.CompanyDelete
	CompanyUpdate company_usecase.CompanyUpdate
	CompanyStore  company_usecase.CompanyStore
}

func NewCompanyRest(companyFetch company_usecase.CompanyFetch, companyGet company_usecase.CompanyGet, companyDelete company_usecase.CompanyDelete, companyUpdate company_usecase.CompanyUpdate, companyStore company_usecase.CompanyStore) *companyRest {
	return &companyRest{CompanyFetch: companyFetch, CompanyGet: companyGet, CompanyDelete: companyDelete, CompanyUpdate: companyUpdate, CompanyStore: companyStore}
}

func (r companyRest) Expose(e echo.Echo) {
	company := e.Group("/company")
	company.GET("/:uuid", r.Show)
	company.GET("", r.Paginate)
	company.POST("", r.Store)
	company.PUT("/:uuid", r.Update)
	company.DELETE("/:uuid", r.Delete)
}

// Show godoc
//	@Summary		Show Company
//	@Description	get detail of company
//	@Tags			company
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string	false	"uuid company"
//	@Success		200	{object}		model.Company
//	@Router			/company/{uuid} [get]
func (r companyRest) Show(ctx echo.Context) error {
	param := company_request.CompanyParam{
	    UUID:  utils.StringP(ctx.Param("uuid")),
	}

	res, err := r.CompanyGet.Get(ctx.Request().Context(), param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	return ctx.JSON(200, res)
}

func (r companyRest) Find(ctx echo.Context) error {
	param := company_request.CompanyParam{}
	err := ctx.Bind(&param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	res, err := r.CompanyFetch.Fetch(ctx.Request().Context(), param)
	if err != nil {
		return ctx.JSON(404, err.Error())
	}

	return ctx.JSON(200, res)
}

// Store godoc
//	@Summary		Store Company
//	@Description	create a company
//	@Tags			company
//	@Accept			json
//	@Produce		json
//	@Param			payload	body	company_request.CompanyStore	true	"payload creates company"
//	@Success		200	{array}		model.Company
//	@Router			/company [post]
func (r companyRest) Store(ctx echo.Context) error {
	param := new(company_request.CompanyStore)
	err := ctx.Bind(param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	err = ctx.Validate(param)
	if err != nil {
		return err
	}

	res, err := r.CompanyStore.Store(ctx.Request().Context(), *param)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, res)
}


// Update godoc
//	@Summary		Update Company
//	@Description	update a company
//	@Tags			company
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string	false	"uuid company"
//	@Param			payload	body company_request.CompanyUpdate	true	"payload updates company"
//	@Success		200	{array}		model.Company
//	@Router			/company/{uuid} [put]
func (r companyRest) Update(ctx echo.Context) error {
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

	res, err := r.CompanyUpdate.Update(ctx.Request().Context(), *param)
	if err != nil {
		return err
	}

	return ctx.JSON(200, res)
}

// Delete godoc
//	@Summary		Delete Company
//	@Description	delete a company by UUID
//	@Tags			company
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string	false	"uuid company"
//	@Success		200	{array}		string
//	@Router			/company/{uuid} [delete]
func (r companyRest) Delete(ctx echo.Context) error {
	uuid := ctx.Param("uuid")
	err := r.CompanyDelete.Delete(ctx.Request().Context(), uuid)
	if err != nil {
		return err
	}

	return ctx.JSON(200, uuid)
}

// Paginate godoc
//	@Summary		Fetch Company with Paginator
//	@Description	fetch company with paginator param and response
//	@Tags			company
//	@Accept			json
//	@Produce		json
//	@Param			id   query		string	false	"search by id"
//	@Param			uuid   query		string	false	"search by uuid"
//	@Param			created_at   query		string	false	"search by created_at"
//	@Param			created_by   query		string	false	"search by created_by"
//	@Param			updated_at   query		string	false	"search by updated_at"
//	@Param			updated_by   query		string	false	"search by updated_by"
//	@Param			deleted_by   query		string	false	"search by deleted_by"
//	@Param			text   query		string	false	"search by text"
//	@Param			parent_company_uuid   query		string	false	"search by parent_company_uuid"
//	@Success		200	{array}		model.Company
//	@Router			/company [get]
func (r companyRest) Paginate(ctx echo.Context) error {
	param := company_request.CompanyParam{}
	err := ctx.Bind(&param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	res, err := r.CompanyFetch.Paginate(ctx.Request().Context(), param)
	if err != nil {
		return ctx.JSON(404, err.Error())
	}

	return ctx.JSON(200, res)
}
