package cmd

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type myApp struct {
	echoHTTP echo.Echo
	gormDB   gorm.DB
}

func NewMyApp(echoHTTP echo.Echo, gormDB gorm.DB) *myApp {
	return &myApp{echoHTTP: echoHTTP, gormDB: gormDB}
}

func (a myApp) Init() {
	//@Generate dependencies
	companyRepo := company_gorm.New(a.gormDB)
	companyFetchU := company_usecase_v1.NewCompanyFetchUsecase(companyRepo, companyRepo)
	companyGetU := company_usecase_v1.NewCompanyGetUsecase(companyRepo)
	companyStoreU := company_usecase_v1.NewCompanyStoreUsecase(companyRepo)
	companyDeleteU := company_usecase_v1.NewCompanyDeleteUsecase(companyRepo)
	companyUpdateU := company_usecase_v1.NewCompanyUpdateUsecase(companyRepo, companyRepo)
	companyRest := company_http_1.NewCompanyRest(companyFetchU, companyGetU, companyDeleteU, companyUpdateU, companyStoreU)
	companyRest.Expose(a.echoHTTP)

	branchRepo := branch_gorm.New(a.gormDB)
	branchFetchU := branch_usecase_v1.NewBranchFetchUsecase(branchRepo, branchRepo)
	branchGetU := branch_usecase_v1.NewBranchGetUsecase(branchRepo)
	branchStoreU := branch_usecase_v1.NewBranchStoreUsecase(branchRepo)
	branchDeleteU := branch_usecase_v1.NewBranchDeleteUsecase(branchRepo)
	branchUpdateU := branch_usecase_v1.NewBranchUpdateUsecase(branchRepo, branchRepo)
	branchRest := branch_http_1.NewBranchRest(branchFetchU, branchGetU, branchDeleteU, branchUpdateU, branchStoreU)
	branchRest.Expose(a.echoHTTP)

}
