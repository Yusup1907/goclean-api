package usecase

import (
	"fmt"
	"goclean/apperror"
	"goclean/model"
	"goclean/repo"
	"time"
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

	trxFinal := &model.TransactionHeaderRepo{
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 2),
		CustName:  trx.CustName,
		Phone:     trx.Phone,
	}

	for _, svc := range trx.ArrDetail {
		det, err := trxUsecase.trxRepo.GetServiceById(svc.Service_Id)
		if err != nil {
			return fmt.Errorf("transaciontUsecaseImpl.CreateTransaction() : %w", err)
		}

		if det == nil {
			return apperror.AppError{
				ErrorCode:    1,
				ErrorMassage: fmt.Sprintf("data service dengan id %v tidak ada", svc.Service_Id),
			}
		}

		data := &model.TransactionDetailRepo{
			ServiceName: det.Name,
			Price:       det.Price,
			Uom:         det.Uom,
			Qty:         svc.Qty,
		}
		trxFinal.ArrDetail = append(trxFinal.ArrDetail, *data)
	}

	return trxUsecase.trxRepo.CreateTransaction(trxFinal)
}

func (trxUsecase *transaciontUsecaseImpl) GetAllTransaction() ([]*model.TransactionHeaderRepo, error) {
	return trxUsecase.trxRepo.GetAllTransaction()
}

func NewTransactionUseCase(trxRepo repo.TransactionRepo) TransactionUsecase {
	return &transaciontUsecaseImpl{
		trxRepo: trxRepo,
	}
}
