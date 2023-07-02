package model

import "time"

type TransactionHeaderRepo struct {
	No        int64
	StartDate time.Time
	EndDate   time.Time
	CustName  string
	Phone     string

	ArrDetail []TransactionDetailRepo
}

type TransactionDetailRepo struct {
	Id          int64
	No          int64
	Service_Id  int
	ServiceName string
	Qty         float64
	Price       float64
	Uom         string
}
