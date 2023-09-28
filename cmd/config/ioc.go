package config

import (
	"service-user/internal/domain/repository"
	"service-user/internal/domain/usecase"
	"service-user/internal/infra/cache"
	"service-user/internal/infra/jwt"
)

func genIoc() (usecase.UserUseCase, usecase.ProgressUseCase) {
	db := GetDB()
	repoUser := repository.NewRepositoryUser(db)
	repoProgress := repository.NewRepositoryProgress(db)
	clientJwt := jwt.NewJwtClient()
	cacheProvider := cache.GetCacheProvider()
	cacheProvider.Set("ping", "pong")
	return usecase.NewUserUserCase(
		repoUser,
		clientJwt,
		cacheProvider,
		repoProgress,
	), usecase.NewProgressUseCase(repoProgress)
}
