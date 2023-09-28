package usecase

import (
	"fmt"
	"net/http"
	"service-user/internal/domain/entity"
	objectValues "service-user/internal/domain/object_values"
	"service-user/internal/domain/repository"
)

type ProgressUseCase struct {
	repoProgress repository.IRepositoryProgress
}

func NewProgressUseCase(repoProgress repository.IRepositoryProgress) ProgressUseCase {
	return ProgressUseCase{
		repoProgress,
	}
}

func (u *ProgressUseCase) GenerateProgress(progress entity.Progress) (interface{}, int) {
	newProgress := new(entity.Progress)
	var err error
	actProgress, _ := u.repoProgress.FindProgressByUserAndCategory(uint64(progress.UserID), uint64(progress.PrimaryID), uint64(progress.SecondaryID), progress.Type)
	if actProgress.ID == 0 {
		fmt.Println("insertar")
		newProgress, err = u.repoProgress.RegisterProgress(progress)
	} else {
		progress.ID = actProgress.ID
		progress.CreatedAt = actProgress.CreatedAt
		newProgress, err = u.repoProgress.UpdateProgress(progress)
	}

	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al generar progreso", nil), http.StatusBadRequest
	}

	AllProgress, err := u.repoProgress.FindProgressByUser(uint64(progress.UserID))
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al generar progreso", nil), http.StatusBadRequest
	}
	fmt.Println(newProgress)
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "progreso generado exitosamente", AllProgress), http.StatusOK
}

func (u *ProgressUseCase) SeeProgress(userId int) (interface{}, int) {
	progress, err := u.repoProgress.FindProgressByUser(uint64(userId))
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest,
			"error",
			"hubo un problema buscando progreso",
			nil,
		), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", progress), http.StatusOK
}
