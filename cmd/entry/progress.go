package entry

import (
	"service-user/internal/domain/entity"
	"service-user/internal/domain/usecase"

	"github.com/labstack/echo/v4"
)

type progressEntry struct {
	userCaseuse usecase.ProgressUseCase
}

func NewProgressEntry(progressUseCase usecase.ProgressUseCase) *progressEntry {
	return &progressEntry{
		progressUseCase,
	}
}

func (u *progressEntry) Register(context echo.Context) error {
	progress := context.Get("progress").(*entity.Progress)
	data, status := u.userCaseuse.GenerateProgress(*progress)
	return context.JSON(status, data)
}
