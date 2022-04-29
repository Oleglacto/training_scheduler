package services

import (
	"context"
	"github.com/oleglacto/traning_scheduler/internal/pkg/models"
	"github.com/oleglacto/traning_scheduler/internal/pkg/repositories"
)

type CityServiceInterface interface {
	AddNewCity(ctx context.Context, name string) (models.City, error)
}

type CityService struct {
	repository repositories.CityRepositoryInterface
}

func (s CityService) AddNewCity(ctx context.Context, name string) (models.City, error) {

	return models.City{}, nil
}
