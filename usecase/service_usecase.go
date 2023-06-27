package usecase

import (
	"goclean/model"
	"goclean/repo"
)

type ServiceUsecase interface {
	GetServiceById(int) (*model.ServiceModel, error)
}

type serviceUsecaseImpl struct {
	svcRepo repo.ServiceRepo
}

func (svcUsecase *serviceUsecaseImpl) GetServiceById(id int) (*model.ServiceModel, error) {
	return svcUsecase.svcRepo.GetServiceById(id)
}

func NewServiceUseCase(svcRepo repo.ServiceRepo) ServiceUsecase {
	return &serviceUsecaseImpl{
		svcRepo: svcRepo,
	}
}
