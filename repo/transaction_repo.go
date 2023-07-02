package repo

import (
	"database/sql"
	"fmt"
	"goclean/model"
)

type TransactionRepo interface {
	CreateTransaction(trx *model.TransactionHeaderRepo) error
	GetServiceById(int) (*model.ServiceModel, error)
	GetTransactionHeaderByName(string) (*model.TransactionHeaderRepo, error)
	GetTransactionDetailByTrxNo(int64) (*model.TransactionDetailRepo, error)
	GetAllTransaction() ([]*model.TransactionHeaderRepo, error)
}

type transactionRepoImpl struct {
	db *sql.DB
}

func (trxRepo *transactionRepoImpl) CreateTransaction(trx *model.TransactionHeaderRepo) error {
	tx, err := trxRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("CreateTransaction() Begin : %w", err)
	}

	qry := "INSERT INTO tr_header(start_date, end_date, cust_name, phone_no) VALUES($1, $2, $3, $4) RETURNING no"

	err = tx.QueryRow(qry, trx.StartDate, trx.EndDate, trx.CustName, trx.Phone).Scan(&trx.No)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("CreateTransaction() Header : %w", err)
	}

	qry = "INSERT INTO tr_detail(trx_no, service_name, qty, price, uom) VALUES($1, $2, $3, $4, $5)"
	for _, det := range trx.ArrDetail {
		_, err := tx.Exec(qry, trx.No, det.ServiceName, det.Qty, det.Price, det.Uom)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("CreateTransaction() Detail : %w", err)
		}
	}

	tx.Commit()

	return nil
}

func (trxRepo *transactionRepoImpl) GetServiceById(id int) (*model.ServiceModel, error) {
	qry := "SELECT id, name, uom, price FROM ms_service WHERE id = $1"

	svc := &model.ServiceModel{}
	err := trxRepo.db.QueryRow(qry, id).Scan(&svc.Id, &svc.Name, &svc.Uom, &svc.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on transactionRepoImpl.getServiceById() : %w", err)
	}
	return svc, nil
}

func (trxRepo *transactionRepoImpl) GetTransactionHeaderByName(custName string) (*model.TransactionHeaderRepo, error) {
	qry := "SELECT no, start_date, end_date, cust_name, phone_no FROM tr_header WHERE cust_name = $1"

	trx := &model.TransactionHeaderRepo{}
	err := trxRepo.db.QueryRow(qry, custName).Scan(&trx.No, &trx.StartDate, &trx.EndDate, &trx.CustName, &trx.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on transactionRepoImpl.GetTransactionHeaderByName() : %w", err)
	}
	return trx, nil
}

func (trxRepo *transactionRepoImpl) GetTransactionDetailByTrxNo(trxNo int64) (*model.TransactionDetailRepo, error) {
	qry := "SELECT id,trx_no, service_name, qty, uom, price FROM tr_detail WHERE trx_no = $1"

	trx := &model.TransactionDetailRepo{}
	err := trxRepo.db.QueryRow(qry, trxNo).Scan(&trx.Id, &trx.No, &trx.ServiceName, &trx.Qty, &trx.Price, &trx.Uom)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on transactionRepoImpl.GetTransactionDetailByTrxNo() : %w", err)
	}
	return trx, nil
}

func (trxRepo *transactionRepoImpl) GetAllTransaction() ([]*model.TransactionHeaderRepo, error) {
	qry := `SELECT 
				trh.no, 
				trh.start_date, 
				trh.end_date, 
				trh.cust_name, 
				trh.phone_no, 
				trd.id,
				trd.trx_no,
				trd.service_name, 
				trd.qty, 
				trd.uom, 
				trd.price
			FROM 
				tr_header trh
			JOIN 
				tr_detail trd ON trh.no = trd.trx_no`

	rows, err := trxRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("GetAllTransactions() Query: %w", err)
	}

	defer rows.Close()

	var arrTransaction []*model.TransactionHeaderRepo
	for rows.Next() {
		trxHeader := &model.TransactionHeaderRepo{}
		trxDetail := model.TransactionDetailRepo{}
		err := rows.Scan(
			&trxHeader.No, &trxHeader.StartDate, &trxHeader.EndDate,
			&trxHeader.CustName, &trxHeader.Phone, &trxDetail.Id, &trxDetail.No, &trxDetail.ServiceName,
			&trxDetail.Qty, &trxDetail.Uom, &trxDetail.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("GetAllTransactions() Scan: %w", err)
		}
		trxHeader.ArrDetail = append(trxHeader.ArrDetail, trxDetail)
		arrTransaction = append(arrTransaction, trxHeader)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllTransactions() Rows: %w", err)
	}

	return arrTransaction, nil
}

func NewTransactionRepo(db *sql.DB) TransactionRepo {
	return &transactionRepoImpl{
		db: db,
	}
}
