package repositories

import (
	"context"
	"github.com/oleglacto/traning_scheduler/internal/pkg/models"
)

type CityRepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.City, error)
}
