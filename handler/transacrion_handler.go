package handler

import (
	"errors"
	"fmt"
	"goclean/apperror"
	"goclean/model"
	"goclean/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
}

type transactionHandlerImpl struct {
	trxUsecase usecase.TransactionUsecase
}

func (trxHandler transactionHandlerImpl) CreateTransaction(ctx *gin.Context) {
	trx := model.TransactionHeaderRepo{
		ArrDetail: []model.TransactionDetailRepo{},
	}
	if err := ctx.ShouldBindJSON(&trx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(trx.CustName) < 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name must have at least 3 characters",
		})
		return
	}

	if !trx.EndDate.After(trx.StartDate) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "End date must be greater than start date",
		})
		return
	}

	if len(trx.Phone) < 11 || len(trx.Phone) > 12 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Phone number must be 11 or 12 digits",
		})
		return
	}

	err := trxHandler.trxUsecase.CreateTransaction(&trx)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("ServiceHandler.CreateTransaction() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("ServiceHandler.CreateTransaction() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Cannot Insert transaction because error",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (trxHandler transactionHandlerImpl) GetAllTransaction(ctx *gin.Context) {
	trx, err := trxHandler.trxUsecase.GetAllTransaction()
	if err != nil {
		fmt.Printf("serviceHandlerImpl.GetAllService() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred when retrieving service data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    trx,
	})
}

func NewTransactionHandler(srv *gin.Engine, trxUsecase usecase.TransactionUsecase) TransactionHandler {
	trxHandler := &transactionHandlerImpl{
		trxUsecase: trxUsecase,
	}
	// srv.GET("/service/:id", svcHandler.GetServiceById)
	srv.GET("/service", trxHandler.GetAllTransaction)
	srv.POST("/transaction", trxHandler.CreateTransaction)
	// srv.PUT("/service/:id", svcHandler.UpdateService)
	// srv.DELETE("/service/:id", svcHandler.DeleteService)

	return trxHandler
}
