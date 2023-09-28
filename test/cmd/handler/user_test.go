package handler_test

import (
	"service-user/cmd/handler"
	"service-user/internal/domain/usecase"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_Handler(t *testing.T) {
	e := handler.NewHandlerUser(echo.New(), usecase.UserUseCase{})
	assert.NotNil(t, e)
}
