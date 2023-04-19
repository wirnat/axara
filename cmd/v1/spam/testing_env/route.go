package testing_env

func Route() {
	//@Generate app

//@Generate route
//Company Route
company := e.Group("company")
company.GET("", func(c echo.Context) error {
	return nil
})
company.GET("", func(c echo.Context) error {
	return nil
})
//Branch Route
branch := e.Group("branch")
branch.GET("", func(c echo.Context) error {
	return nil
})
branch.GET("", func(c echo.Context) error {
	return nil
})

	//@Generate endApp
}
