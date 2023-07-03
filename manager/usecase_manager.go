package manager

import (
	"goclean/usecase"
	"sync"
)

type UsecaseManager interface {
	GetServiceUsecase() usecase.ServiceUsecase
	GetTransactionUsecase() usecase.TransactionUsecase
}

type usecaseManager struct {
	repoManager RepoManager

	svcUsecase usecase.ServiceUsecase

	trxUsecase                 usecase.TransactionUsecase
	onceLoadServiceUsecase     sync.Once
	onceLoadTransactionUsecase sync.Once
}

func (um *usecaseManager) GetServiceUsecase() usecase.ServiceUsecase {
	um.onceLoadServiceUsecase.Do(func() {
		um.svcUsecase = usecase.NewServiceUseCase(um.repoManager.GetServiceRepo())
	})
	return um.svcUsecase
}

func (um *usecaseManager) GetTransactionUsecase() usecase.TransactionUsecase {
	um.onceLoadTransactionUsecase.Do(func() {
		trxRepo := um.repoManager.GetTransactionRepo()
		svcRepo := um.repoManager.GetServiceRepo()
		um.trxUsecase = usecase.NewTransactionUseCase(trxRepo, svcRepo)
	})
	return um.trxUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
