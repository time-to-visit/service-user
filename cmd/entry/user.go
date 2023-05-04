package entry

import (
	"service-user/internal/domain/entity"
	objectValues "service-user/internal/domain/object_values"
	"service-user/internal/domain/usecase"

	"github.com/labstack/echo/v4"
)

type userEntry struct {
	userCaseuse usecase.UserUseCase
}

func NewUserEntry(userCaseuse usecase.UserUseCase) *userEntry {
	return &userEntry{
		userCaseuse,
	}
}

func (u *userEntry) Register(context echo.Context) error {
	user := context.Get("user-data").(*entity.User)
	data, status := u.userCaseuse.RegisterUser(*user)
	return context.JSON(status, data)
}

func (u *userEntry) Auth(context echo.Context) error {
	verify := context.Get("verify").(*objectValues.Verify)
	data, status := u.userCaseuse.LoginUser(*verify)
	return context.JSON(status, data)
}

func (u *userEntry) Verify(context echo.Context) error {
	data, status := u.userCaseuse.VerifyUser(context)
	return context.JSON(status, data)
}

func (u *userEntry) SendEmail(context echo.Context) error {
	return nil
}

func (u *userEntry) VerifyEmail(context echo.Context) error {
	return nil
}

func (u *userEntry) VerifyEmailView(context echo.Context) error {
	return nil
}
