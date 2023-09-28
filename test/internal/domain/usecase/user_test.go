package usecase_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"service-user/internal/domain/entity"
	objectValues "service-user/internal/domain/object_values"
	"service-user/internal/domain/usecase"
	"service-user/internal/infra/cache"
	"service-user/internal/infra/jwt"
	mocks "service-user/test/mock"
	"testing"

	jwtC "github.com/golang-jwt/jwt"

	"github.com/go-playground/assert/v2"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

func deleteFolder() error {
	err := filepath.Walk("./data/cache", func(ruta string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			err = os.Remove(ruta)
			if err != nil {
				return err
			}
			fmt.Println("Archivo borrado:", ruta)
		}
		return nil
	})
	return err
}
func Test_RegisterUser(t *testing.T) {
	deleteFolder()
	userRepo := mocks.NewIRepositoryUser(t)
	userUseCase := usecase.NewUserUserCase(userRepo, jwt.NewJwtClient(), cache.GetCacheProvider())
	user := entity.User{
		Model: entity.Model{
			ID: 1,
		},
	}
	userRepo.On("RegisterUser", mock.Anything).Return(&user, nil)
	_, status := userUseCase.RegisterUser(user)
	assert.Equal(t, status, http.StatusCreated)
}

func Test_RegisterUserErr(t *testing.T) {
	deleteFolder()
	userRepo := mocks.NewIRepositoryUser(t)
	userUseCase := usecase.NewUserUserCase(userRepo, jwt.NewJwtClient(), cache.GetCacheProvider())
	user := entity.User{
		Model: entity.Model{
			ID: 1,
		},
	}
	userRepo.On("RegisterUser", mock.Anything).Return(&user, errors.New("err"))
	_, status := userUseCase.RegisterUser(entity.User{})
	assert.Equal(t, status, http.StatusBadRequest)
}

func TestVerifyUser(t *testing.T) {
	deleteFolder()
	cacheProvider := cache.GetCacheProvider()
	cacheProvider.Set("token-usuario", "asdadadadsasdasd")
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/some-route", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", &jwtC.Token{
		Valid: true,
		Claims: &objectValues.JwtCustomClaims{
			Correo: "usuario",
		},
	})
	userRepo := mocks.NewIRepositoryUser(t)
	userUseCase := usecase.NewUserUserCase(userRepo, jwt.NewJwtClient(), cacheProvider)
	_, status := userUseCase.VerifyUser(c)
	assert.Equal(t, status, http.StatusOK)
}

func TestVerifyErr(t *testing.T) {
	deleteFolder()
	cacheProvider := cache.GetCacheProvider()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/some-route", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", &jwtC.Token{
		Valid: false,
		Claims: &objectValues.JwtCustomClaims{
			Correo: "usuario",
		},
	})
	userRepo := mocks.NewIRepositoryUser(t)
	userUseCase := usecase.NewUserUserCase(userRepo, jwt.NewJwtClient(), cacheProvider)
	_, status := userUseCase.VerifyUser(c)
	assert.Equal(t, status, http.StatusUnauthorized)
}
