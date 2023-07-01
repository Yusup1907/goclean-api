package main

import (
	"database/sql"
	"goclean/handler"
	"goclean/repo"
	"goclean/usecase"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	srv := gin.Default()

	db, err := sql.Open("postgres", "user=postgres host=localhost password=12345 dbname=laundry sslmode=disable")
	if err != nil {
		log.Fatal("Cannot start app, error when connect to DB", err.Error())
	}
	defer db.Close()

	// Service
	svcRepo := repo.NewServiceRepo(db)
	svcUsecase := usecase.NewServiceUseCase(svcRepo)
	handler.NewServiceHandler(srv, svcUsecase)

	// Transaction
	trxRepo := repo.NewTransactionRepo(db)
	trxUsecase := usecase.NewTransactionUseCase(trxRepo)
	handler.NewTransactionHandler(srv, trxUsecase)

	srv.Run()
}
