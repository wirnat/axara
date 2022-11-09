package branch_delivery

import (
	"github.com/wirnat/axara/example/infrastructure/contextor"
	"github.com/wirnat/axara/module/branch/branch_usecase"
	"github.com/wirnat/axara/module/branch/repository/branch_repository"
	"github.com/wirnat/axara/module/branch/request/branch_request"
)

type BranchRest struct {
	BranchInteractor branch_usecase.BranchUsecase
	BranchFetchRepo  branch_repository.BranchFetch
	BranchGetRepo    branch_repository.BranchGet
	BranchDeleteRepo branch_repository.BranchDelete
	BranchPaginate   branch_repository.BranchPaginate
}

func NewBranchRest() *BranchRest {
	return &BranchRest{}
}

func (r BranchRest) Show(ctx *contextor.Contextor) error {
	param := branch_request.BranchParam{}
	err := ctx.BindQuery(&param)
	if err != nil {
		return err
	}

	res, err := r.BranchGetRepo.Get(ctx.ToContext(), param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	return ctx.JSON(200, res)
}

func (r BranchRest) Find(ctx *contextor.Contextor) error {
	param := branch_request.BranchParam{}
	err := ctx.BindQuery(&param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	res, err := r.BranchFetchRepo.Fetch(ctx.ToContext(), param)
	if err != nil {
		return ctx.JSON(404, err.Error())
	}

	return ctx.JSON(200, res)
}

func (r BranchRest) Store(ctx *contextor.Contextor) error {
	param := new(branch_request.BranchStore)
	err := ctx.Bind(param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	err = ctx.Validate(param)
	if err != nil {
		return err
	}

	res, err := r.BranchInteractor.Store(ctx.ToContext(), *param)
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, res)
}

func (r BranchRest) Update(ctx *contextor.Contextor) error {
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

	res, err := r.BranchInteractor.Update(ctx.ToContext(), *param)
	if err != nil {
		return err
	}

	return ctx.JSON(200, res)
}

func (r BranchRest) Delete(ctx *contextor.Contextor) error {
	uuid := ctx.Param("uuid")
	err := r.BranchDeleteRepo.Delete(ctx.ToContext(), uuid)
	if err != nil {
		return err
	}

	return ctx.JSON(200, uuid)
}

func (r BranchRest) Paginate(ctx *contextor.Contextor) error {
	param := branch_request.BranchParam{}
	err := ctx.Bind(&param)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	res, err := r.BranchPaginate.Paginate(ctx.ToContext(), param)
	if err != nil {
		return ctx.JSON(404, err.Error())
	}

	return ctx.JSON(200, res)
}