package branch_http_1

import (
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/request/branch_request"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/usecase/branch_usecase"
	"github.com/labstack/echo/v4"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/infrastructure/utils"
)

type branchRest struct {
	BranchFetch  branch_usecase.BranchFetch
	BranchGet    branch_usecase.BranchGet
	BranchDelete branch_usecase.BranchDelete
	BranchUpdate branch_usecase.BranchUpdate
	BranchStore  branch_usecase.BranchStore
}

func NewBranchRest(branchFetch branch_usecase.BranchFetch, branchGet branch_usecase.BranchGet, branchDelete branch_usecase.BranchDelete, branchUpdate branch_usecase.BranchUpdate, branchStore branch_usecase.BranchStore) *branchRest {
	return &branchRest{BranchFetch: branchFetch, BranchGet: branchGet, BranchDelete: branchDelete, BranchUpdate: branchUpdate, BranchStore: branchStore}
}

func (r branchRest) Expose(e echo.Echo) {
	branch := e.Group("/branch")
	branch.GET("/:uuid", r.Show)
	branch.GET("", r.Paginate)
	branch.POST("", r.Store)
	branch.PUT("/:uuid", r.Update)
	branch.DELETE("/:uuid", r.Delete)
}

// Show godoc
//	@Summary		Show Branch
//	@Description	get detail of branch
//	@Tags			branch
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string	false	"uuid branch"
//	@Success		200	{object}		model.Branch
//	@Router			/branch/{uuid} [get]
func (r branchRest) Show(ctx echo.Context) error {
	param := branch_request.BranchParam{
	    UUID:  utils.StringP(ctx.Param("uuid")),
	}

	res, err := r.BranchGet.Get(ctx.Request().Context(), param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	return ctx.JSON(200, res)
}

func (r branchRest) Find(ctx echo.Context) error {
	param := branch_request.BranchParam{}
	err := ctx.Bind(&param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	res, err := r.BranchFetch.Fetch(ctx.Request().Context(), param)
	if err != nil {
		return ctx.JSON(404, err.Error())
	}

	return ctx.JSON(200, res)
}

// Store godoc
//	@Summary		Store Branch
//	@Description	create a branch
//	@Tags			branch
//	@Accept			json
//	@Produce		json
//	@Param			payload	body	branch_request.BranchStore	true	"payload creates branch"
//	@Success		200	{array}		model.Branch
//	@Router			/branch [post]
func (r branchRest) Store(ctx echo.Context) error {
	param := new(branch_request.BranchStore)
	err := ctx.Bind(param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	err = ctx.Validate(param)
	if err != nil {
		return err
	}

	res, err := r.BranchStore.Store(ctx.Request().Context(), *param)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, res)
}


// Update godoc
//	@Summary		Update Branch
//	@Description	update a branch
//	@Tags			branch
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string	false	"uuid branch"
//	@Param			payload	body branch_request.BranchUpdate	true	"payload updates branch"
//	@Success		200	{array}		model.Branch
//	@Router			/branch/{uuid} [put]
func (r branchRest) Update(ctx echo.Context) error {
	param := new(branch_request.BranchUpdate)
	err := ctx.Bind(param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	param.UUID = ctx.Param("uuid")

	err = ctx.Validate(param)
	if err != nil {
		return err
	}

	res, err := r.BranchUpdate.Update(ctx.Request().Context(), *param)
	if err != nil {
		return err
	}

	return ctx.JSON(200, res)
}

// Delete godoc
//	@Summary		Delete Branch
//	@Description	delete a branch by UUID
//	@Tags			branch
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string	false	"uuid branch"
//	@Success		200	{array}		string
//	@Router			/branch/{uuid} [delete]
func (r branchRest) Delete(ctx echo.Context) error {
	uuid := ctx.Param("uuid")
	err := r.BranchDelete.Delete(ctx.Request().Context(), uuid)
	if err != nil {
		return err
	}

	return ctx.JSON(200, uuid)
}

// Paginate godoc
//	@Summary		Fetch Branch with Paginator
//	@Description	fetch branch with paginator param and response
//	@Tags			branch
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
//	@Param			parent_branch_uuid   query		string	false	"search by parent_branch_uuid"
//	@Success		200	{array}		model.Branch
//	@Router			/branch [get]
func (r branchRest) Paginate(ctx echo.Context) error {
	param := branch_request.BranchParam{}
	err := ctx.Bind(&param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	res, err := r.BranchFetch.Paginate(ctx.Request().Context(), param)
	if err != nil {
		return ctx.JSON(404, err.Error())
	}

	return ctx.JSON(200, res)
}
