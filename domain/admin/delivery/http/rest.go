package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/halo-suster/domain/admin/interfaces"
	"github.com/mrakhaf/halo-suster/models/request"
	"github.com/mrakhaf/halo-suster/shared/common"
	"github.com/mrakhaf/halo-suster/shared/common/jwt"
	"github.com/mrakhaf/halo-suster/shared/utils"
)

type handlerAdmin struct {
	usecase    interfaces.Usecase
	repository interfaces.Repository
	Json       common.JSON
	jwtAccess  *jwt.JWT
}

func AdminHandler(publicRoute *echo.Group, restrictedRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, Json common.JSON, JwtAccess *jwt.JWT) {
	handler := handlerAdmin{
		usecase:    usecase,
		repository: repository,
		Json:       Json,
		jwtAccess:  JwtAccess,
	}

	publicRoute.POST("/admin/register", handler.Register)
	publicRoute.POST("/admin/login", handler.Login)

	//merchant
	restrictedRoute.POST("/admin/merchants", handler.CreateMerchant)
	restrictedRoute.GET("/admin/merchants", handler.GetMerchants)

}

func (h *handlerAdmin) Register(c echo.Context) error {

	var req request.Register

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	data, err := h.usecase.Register(req)
	if err != nil && (err.Error() == "user already exist" || err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"") {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})

	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusCreated, "Register success", data)
}

func (h *handlerAdmin) Login(c echo.Context) error {
	var req request.Login

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	data, err := h.usecase.Login(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusOK, "Login success", data)

}

func (h *handlerAdmin) CreateMerchant(c echo.Context) error {

	var req request.MerchantRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	isImage := utils.CheckImageType(req.ImageUrl)

	if !isImage {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "image not valid"})
	}

	//check token
	_, role, err := h.jwtAccess.GetUserIdFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "user not admin"})
	}

	//create merchant
	data, err := h.usecase.CreateMerchant(req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusCreated, "Create merchant success", data)
}

func (h *handlerAdmin) GetMerchants(c echo.Context) error {

	//check token
	_, role, err := h.jwtAccess.GetUserIdFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "user not admin"})
	}

	var req request.GetMerchants

	//bind query param
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	data, err := h.usecase.GetMerchants(req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, data)

}
