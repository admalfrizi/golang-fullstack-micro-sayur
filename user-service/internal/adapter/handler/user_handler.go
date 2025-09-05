package handler

import (
	"context"
	"net/http"
	"user-service/internal/adapter/handler/request"
	"user-service/internal/adapter/handler/response"
	"user-service/internal/core/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserHandlerInterface interface {
	SignIn(c echo.Context) error
}

type userHandler struct {
	userService service.UserServiceInterface
}

// SignIn implements UserHandlerInterface.
func (u *userHandler) SignIn(c echo.Context) error {
	var (
		req = request.SignInRequest{}
		resp = response.DefaultResponse{}
		respSignIn = response.SignInResponse{}
		ctx = c.Request().Context()
	)

	if err = c.Bind(&req); err != nil {
		log.Errorf("[UserHandler-1] SignIn: %v", err)
		resp.Message = err.Error()
		resp.Data = nil
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}


}

var err error

func NewUserHandler(userService service.UserServiceInterface) UserHandlerInterface {
	return &userHandler{userService: userService}
}
