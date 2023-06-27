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

	db, err := sql.Open("postgres", "user=postgres host=localhost password=P@ssw0rd dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal("Cannot start app, error when connect to DB", err.Error())
	}
	defer db.Close()

	svcRepo := repo.NewServiceRepo(db)
	svcUsecase := usecase.NewServiceUseCase(svcRepo)
	handler.NewServiceHandler(srv, svcUsecase)

	srv.Run()
}
