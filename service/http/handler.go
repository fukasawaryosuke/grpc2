package http

import (
	"net/http"

	"github.com/fukasawaryosuke/serve_streaming_grpc_app/service/usage/usecase"

	"github.com/labstack/echo/v4"
)

type IUsageHandler interface {
	SampleGrpc(c echo.Context) error
}

type handler struct {
	uu usecase.IUsageUsecase
}

func NewUsageHandler(uu usecase.IUsageUsecase) IUsageHandler {
	return &handler{uu}
}

func (tc *handler) SampleGrpc(c echo.Context) error {
	// usecaseを呼び出す
	tc.uu.GetDessertStream()
	return c.JSON(http.StatusOK, "Success!")
}
