package usecase

import (
	"fmt"
	"goclean/apperror"
	"goclean/model"
	"goclean/repo"
)

type ServiceUsecase interface {
	GetServiceById(int) (*model.ServiceModel, error)
	GetAllService() ([]*model.ServiceModel, error)
	CreateService(svc *model.ServiceModel) error
	UpdateService(id int, svc *model.ServiceModel) error
	DeleteService(int) error
}

type serviceUsecaseImpl struct {
	svcRepo repo.ServiceRepo
}

func (svcUsecase *serviceUsecaseImpl) GetServiceById(id int) (*model.ServiceModel, error) {
	return svcUsecase.svcRepo.GetServiceById(id)
}

func (svcUsecase *serviceUsecaseImpl) GetAllService() ([]*model.ServiceModel, error) {
	return svcUsecase.svcRepo.GetAllService()
}

func (svcUsecase *serviceUsecaseImpl) CreateService(svc *model.ServiceModel) error {
	svcDB, err := svcUsecase.svcRepo.GetServiceByName(svc.Name)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.InsertService() : %w", err)
	}

	if svcDB != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data service dengan nama %v sudah ada", svc.Name),
		}
	}

	return svcUsecase.svcRepo.CreateService(svc)
}

func (svcUsecase *serviceUsecaseImpl) UpdateService(id int, svc *model.ServiceModel) error {
	svcDB, err := svcUsecase.svcRepo.GetServiceById(id)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.UpdateService() : %w", err)
	}

	if svcDB == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data service dengan id %v tidak ada", svc.Id),
		}
	}

	svcDb, err := svcUsecase.svcRepo.GetServiceByName(svc.Name)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.InsertService() : %w", err)
	}

	if svcDb != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data service dengan nama %v sudah ada", svc.Name),
		}
	}

	return svcUsecase.svcRepo.UpdateService(id, svc)
}

func (svcUsecase *serviceUsecaseImpl) DeleteService(id int) error {
	svcDB, err := svcUsecase.svcRepo.GetServiceById(id)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.DeleteService(): %w", err)
	}

	if svcDB == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data service dengan id %v tidak ada", id),
		}
	}

	return svcUsecase.svcRepo.DeleteService(id)
}

func NewServiceUseCase(svcRepo repo.ServiceRepo) ServiceUsecase {
	return &serviceUsecaseImpl{
		svcRepo: svcRepo,
	}
}
