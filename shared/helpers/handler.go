package helpers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SuccessResponse(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   data,
	})
}
