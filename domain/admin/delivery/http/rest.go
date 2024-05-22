package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/halo-suster/domain/admin/interfaces"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/shared/common"
)

type handlerAdmin struct {
	usecase    interfaces.Usecase
	repository interfaces.Repository
	Json       common.JSON
}

func AdminHandler(authRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, Json common.JSON) {
	handler := handlerAdmin{
		usecase:    usecase,
		repository: repository,
		Json:       Json,
	}

	authRoute.POST("/user/it/register", handler.Register)
	authRoute.POST("/user/it/login", handler.Login)
}

func (h *handlerAdmin) Register(c echo.Context) error {

	var req request.Register

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return nil
}

func (h *handlerAdmin) Login(c echo.Context) error {
	var req request.Login

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusOK, "Login success", nil)

}
