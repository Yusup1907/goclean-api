package usecase

import (
	"fmt"
	"goclean/apperror"
	"goclean/model"
	"goclean/repo"
)

type TransactionUsecase interface {
	CreateTransaction(trx *model.TransactionHeaderRepo) error
	GetAllTransaction() ([]*model.TransactionHeaderRepo, error)
}

type transaciontUsecaseImpl struct {
	trxRepo repo.TransactionRepo
}

func (trxUsecase *transaciontUsecaseImpl) CreateTransaction(trx *model.TransactionHeaderRepo) error {
	trxDB, err := trxUsecase.trxRepo.GetTransactionHeaderByName(trx.CustName)
	if err != nil {
		return fmt.Errorf("transaciontUsecaseImpl.CreateTransaction() : %w", err)
	}

	if trxDB != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data service dengan nama %v sudah ada", trx.CustName),
		}
	}
	return trxUsecase.trxRepo.CreateTransaction(trx)
}

func (trxUsecase *transaciontUsecaseImpl) GetAllTransaction() ([]*model.TransactionHeaderRepo, error) {
	return trxUsecase.trxRepo.GetAllTransaction()
}

func NewTransactionUseCase(trxRepo repo.TransactionRepo) TransactionUsecase {
	return &transaciontUsecaseImpl{
		trxRepo: trxRepo,
	}
}
