package usecase

import (
	"net/http"
	"service-user/internal/domain/entity"
	objectValues "service-user/internal/domain/object_values"
	"service-user/internal/domain/repository"
	"service-user/internal/infra/cache"
	"service-user/internal/infra/jwt"

	jwtC "github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
)

type UserUseCase struct {
	repoUser repository.IRepositoryUser
	jwt      jwt.JwtClient
	cache    *cache.CacheProvider
}

func NewUserUserCase(repoUser repository.IRepositoryUser, jwt jwt.JwtClient, cache *cache.CacheProvider) UserUseCase {
	return UserUseCase{
		repoUser,
		jwt,
		cache,
	}
}

func (u *UserUseCase) RegisterUser(user entity.User) (interface{}, int) {
	newUser, err := u.repoUser.RegisterUser(user)
	newUser.Clave = ""
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al registrar el usuario", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusCreated, "ok", "usuario registrado sastifactoriamente", newUser), http.StatusCreated
}

func (u *UserUseCase) LoginUser(verify objectValues.Verify) (interface{}, int) {
	user := u.repoUser.FindUserByEmailAndPassword(verify.Correo, verify.Clave)
	if user != nil && user.Correo == verify.Correo {
		token, err := u.jwt.GenerateToken(objectValues.JwtEntity{
			Nombres:   user.Nombres,
			Apellidos: user.Apellidos,
			Id:        user.ID,
			Correo:    user.Correo,
		})
		if err != nil {
			return objectValues.NewResponseWithData(http.StatusUnauthorized, "no authorization", "el usuario no es valido", nil), http.StatusUnauthorized
		}
		u.cache.Set("token-"+user.Correo, token)
		return objectValues.NewResponseWithData(http.StatusOK, "ok", "usuario autorizado", map[string]string{
			"token": token,
		}), http.StatusOK
	}
	return objectValues.NewResponseWithData(http.StatusUnauthorized, "no authorization", "el usuario no es valido", nil), http.StatusUnauthorized
}

func (u *UserUseCase) VerifyUser(c echo.Context) (map[string]interface{}, int) {
	user := c.Get("user").(*jwtC.Token)
	if user.Valid {

		claims := user.Claims.(*objectValues.JwtCustomClaims)
		byteToken := u.cache.Get("token-" + claims.Correo)
		if byteToken != nil {
			return map[string]interface{}{
				"valid": true,
				"data":  claims,
			}, http.StatusOK
		}
	}
	return map[string]interface{}{
		"valid": false,
	}, http.StatusUnauthorized
}

func (u *UserUseCase) ChangePassword() {

}

func (u *UserUseCase) SendRestartEmail() {

}
