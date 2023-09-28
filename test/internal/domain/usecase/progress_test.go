package usecase_test

import (
	"errors"
	"net/http"
	"service-user/internal/domain/entity"
	"service-user/internal/domain/usecase"
	mocks "service-user/test/mock"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func Test_UseCase_User_GenerateProgressOk(t *testing.T) {
	progressRe := mocks.NewIRepositoryProgress(t)
	progress := entity.Progress{
		Model: entity.Model{ID: 1},
	}

	progressRe.On("FindProgressByUserAndCategory", mock.Anything, mock.Anything).Return(&progress, nil)
	progressRe.On("UpdateProgress", mock.Anything).Return(nil, nil)
	progressUseCase := usecase.NewProgressUseCase(progressRe)
	_, status := progressUseCase.GenerateProgress(progress)
	assert.Equal(t, status, http.StatusOK)
}

func Test_UseCase_User_GenerateProgressErrinSearch(t *testing.T) {
	progressRe := mocks.NewIRepositoryProgress(t)
	progress := entity.Progress{
		Model: entity.Model{ID: 0},
	}

	progressRe.On("FindProgressByUserAndCategory", mock.Anything, mock.Anything).Return(&progress, nil)
	progressRe.On("RegisterProgress", mock.Anything).Return(nil, nil)
	progressUseCase := usecase.NewProgressUseCase(progressRe)
	_, status := progressUseCase.GenerateProgress(progress)
	assert.Equal(t, status, http.StatusOK)
}

func Test_UseCase_User_GenerateProgressErrinInsert(t *testing.T) {
	progressRe := mocks.NewIRepositoryProgress(t)
	progress := entity.Progress{
		Model: entity.Model{ID: 0},
	}

	progressRe.On("FindProgressByUserAndCategory", mock.Anything, mock.Anything).Return(&progress, nil)
	progressRe.On("RegisterProgress", mock.Anything).Return(nil, errors.New("err"))
	progressUseCase := usecase.NewProgressUseCase(progressRe)
	_, status := progressUseCase.GenerateProgress(progress)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_UseCase_User_SeeProgress(t *testing.T) {
	progressRe := mocks.NewIRepositoryProgress(t)
	progress := []entity.Progress{
		{
			Model: entity.Model{ID: 0},
		},
	}

	progressRe.On("FindProgressByUser", mock.Anything).Return(progress, nil)
	progressUseCase := usecase.NewProgressUseCase(progressRe)
	_, status := progressUseCase.SeeProgress(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_UseCase_User_SeeProgressErr(t *testing.T) {
	progressRe := mocks.NewIRepositoryProgress(t)
	progressRe.On("FindProgressByUser", mock.Anything).Return(nil, errors.New("err"))
	progressUseCase := usecase.NewProgressUseCase(progressRe)
	_, status := progressUseCase.SeeProgress(1)
	assert.Equal(t, status, http.StatusBadRequest)

}
