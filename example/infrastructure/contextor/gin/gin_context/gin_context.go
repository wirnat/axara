package gin_context

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.com/aksaratech/barber-backend/infrastructure/contextor"
	"mime/multipart"
	"net/http"
)

type ginContext struct {
	Gin *gin.Context
}

func (g *ginContext) Request() *http.Request {
	return g.Gin.Request
}

func (g *ginContext) Validate(i interface{}) error {
	return g.Gin.ShouldBindJSON(i)
}

func (g *ginContext) FormFile(name string) (*multipart.FileHeader, error) {
	return g.Gin.FormFile(name)
}

func (g *ginContext) ToContext() (c context.Context) {
	user, _ := g.Gin.Get(contextor.UserID)
	c = context.WithValue(g.Gin.Request.Context(), contextor.UserID, user)

	company, _ := g.Gin.Get(contextor.Company)
	c = context.WithValue(c, contextor.Company, company)

	token, _ := g.Gin.Get(contextor.Token)
	c = context.WithValue(c, contextor.Token, token)
	return g.Gin.Request.Context()
}

func (g *ginContext) QueryParam(name string) string {
	return g.Gin.Query(name)
}

func (g *ginContext) JSON(code int, i interface{}) error {
	g.Gin.JSON(code, i)
	return nil
}

func (g *ginContext) Param(name string) string {
	return g.Gin.Param(name)
}

func (g *ginContext) Bind(dest interface{}) error {
	return g.Gin.Bind(dest)
}

func NewGinContext(gin *gin.Context) *ginContext {
	return &ginContext{Gin: gin}
}
