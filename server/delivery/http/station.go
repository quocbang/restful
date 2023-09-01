package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type stationHandler struct {
}

func NewStationHandler(e *echo.Echo) {
	handler := &stationHandler{}
	g := e.Group("/api")
	g.POST("/station/upload", handler.UploadStation)
}

func (s *stationHandler) UploadStation(ctx echo.Context) error {
	body := ctx.Request().Body

	_, err := io.ReadAll(body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to read body, error: %s", err.Error()))
	}

	return ctx.JSON(http.StatusOK, body)
}
