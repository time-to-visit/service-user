package handler

import (
	"service-user/cmd/entry"
	"service-user/internal/domain/usecase"
	validator "service-user/internal/domain/validator"
	"service-user/internal/infra/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewHandlerUser(e *echo.Echo, userUseCase usecase.UserUseCase, progressUseCase usecase.ProgressUseCase) *echo.Echo {
	userEntry := entry.NewUserEntry(userUseCase)
	progressEntry := entry.NewProgressEntry(progressUseCase)
	jwtConf := jwt.NewJwtClient()
	e.POST("/register", userEntry.Register, validator.ValidateUser)
	e.POST("/verify", userEntry.Verify, middleware.JWTWithConfig(jwtConf.GetConfig()))
	e.POST("/auth", userEntry.Auth, validator.ValidateVerify)
	e.POST("/verify-email", userEntry.VerifyEmail)
	e.GET("/verify-email", userEntry.VerifyEmailView)
	e.POST("/send-email", userEntry.SendEmail)
	e.POST("/progress", progressEntry.Register, validator.ValidateProgress)

	return e
}
