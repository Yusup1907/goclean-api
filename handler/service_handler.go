package handler

import (
	"fmt"
	"goclean/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceHandler interface {
}

type serviceHandlerImpl struct {
	svcUsecase usecase.ServiceUsecase
}

func (svcHandler serviceHandlerImpl) GetServiceById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id harus angka",
		})
		return
	}

	svc, err := svcHandler.svcUsecase.GetServiceById(id)
	if err != nil {
		fmt.Printf("serviceHandlerImpl.GetServiceById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data service",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    svc,
	})
}

func NewServiceHandler(srv *gin.Engine, svcUsecase usecase.ServiceUsecase) ServiceHandler {
	svcHandler := &serviceHandlerImpl{
		svcUsecase: svcUsecase,
	}
	srv.GET("/service/:id", svcHandler.GetServiceById)

	return svcHandler
}
