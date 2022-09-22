package query_reader

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"ms-glow-erp/barberque/domain/request"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_queryReader_Bind(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		e := echo.New()
		uu := uuid.New()
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/?uuid=%v&id=2&latitude=-41290", uu), strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		reqx := request.BranchParam{}

		err := Bind(ctx, &reqx)

		assert.Nil(t, err)
		assert.Equal(t, uu.String(), *reqx.UUID)
	})
}
