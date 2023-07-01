package model

import "time"

type TransactionHeaderRepo struct {
	No        int64
	StartDate time.Time `binding:"required"`
	EndDate   time.Time `binding:"required"`
	CustName  string    `binding:"required"`
	Phone     string    `binding:"required"`

	ArrDetail []TransactionDetailRepo
}

type TransactionDetailRepo struct {
	Id          int64
	No          int64   `binding:"required"`
	ServiceName string  `binding:"required"`
	Qty         float64 `binding:"required"`
	Price       float64 `binding:"required"`
	Uom         string  `binding:"required"`
}
