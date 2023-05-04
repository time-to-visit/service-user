package config

import (
	"service-user/internal/domain/repository"
	"service-user/internal/domain/usecase"
	"service-user/internal/infra/cache"
	"service-user/internal/infra/jwt"
)

func genIoc() usecase.UserUseCase {
	db := GetDB()
	repoUser := repository.NewRepositoryUser(db)
	clientJwt := jwt.NewJwtClient()
	cacheProvider := cache.GetCacheProvider()
	cacheProvider.Set("ping", "pong")
	return usecase.NewUserUserCase(repoUser, clientJwt, cacheProvider)
}
