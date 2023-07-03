package manager

import (
	"goclean/repo"
	"sync"
)

type RepoManager interface {
	GetServiceRepo() repo.ServiceRepo
	GetTransactionRepo() repo.TransactionRepo
}

type repoManager struct {
	infraManager InfraManager

	svcRepo repo.ServiceRepo

	trxRepo repo.TransactionRepo
}

var onceLoadServiceRepo sync.Once
var onceLoadTransactionRepo sync.Once

func (rm *repoManager) GetServiceRepo() repo.ServiceRepo {
	onceLoadServiceRepo.Do(func() {
		rm.svcRepo = repo.NewServiceRepo(rm.infraManager.GetDB())
	})
	return rm.svcRepo
}

func (rm *repoManager) GetTransactionRepo() repo.TransactionRepo {
	onceLoadTransactionRepo.Do(func() {
		rm.trxRepo = repo.NewTransactionRepo(rm.infraManager.GetDB())
	})
	return rm.trxRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
