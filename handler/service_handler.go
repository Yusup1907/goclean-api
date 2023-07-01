package handler

import (
	"errors"
	"fmt"
	"goclean/apperror"
	"goclean/model"
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
			"errorMessage": "An error occurred when retrieving service data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    svc,
	})
}

func (svcHandler serviceHandlerImpl) GetAllService(ctx *gin.Context) {

	svc, err := svcHandler.svcUsecase.GetAllService()
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
		"data":    svc,
	})
}

func (svcHandler serviceHandlerImpl) CreateService(ctx *gin.Context) {
	var svc model.ServiceModel
	err := ctx.ShouldBindJSON(&svc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if svc.Name == "" || svc.Price == 0 || svc.Uom == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name, Price, and Uom cannot be empty",
		})
		return
	}

	if len(svc.Name) < 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name must have at least 3 characters",
		})
		return
	}

	if svc.Price < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Price is not negative",
		})
		return
	}

	err = svcHandler.svcUsecase.CreateService(&svc)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("ServiceHandler.CreateService() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("ServiceHandler.CreateService() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Cannot Insert service because error",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (svcHandler serviceHandlerImpl) UpdateService(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id cannot be empty",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id must be a number",
		})
		return
	}

	var svc model.ServiceModel
	err = ctx.ShouldBindJSON(&svc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if svc.Name == "" || svc.Price == 0 || svc.Uom == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name, Price, and Uom cannot be empty",
		})
		return
	}

	if len(svc.Name) < 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name must have at least 3 characters",
		})
		return
	}

	if svc.Price < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Price is not negative",
		})
		return
	}

	err = svcHandler.svcUsecase.UpdateService(id, &svc)
	if err != nil {
		fmt.Printf("serviceHandlerImpl.GetAllService() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Cannot Update service because error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (svcHandler serviceHandlerImpl) DeleteService(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id cannot be empty",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id must be a number",
		})
		return
	}

	err = svcHandler.svcUsecase.DeleteService(id)
	if err != nil {
		fmt.Printf("serviceHandlerImpl.GetAllService() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Cannot Delete service because error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewServiceHandler(srv *gin.Engine, svcUsecase usecase.ServiceUsecase) ServiceHandler {
	svcHandler := &serviceHandlerImpl{
		svcUsecase: svcUsecase,
	}
	srv.GET("/service/:id", svcHandler.GetServiceById)
	srv.GET("/service", svcHandler.GetAllService)
	srv.POST("/service", svcHandler.CreateService)
	srv.PUT("/service/:id", svcHandler.UpdateService)
	srv.DELETE("/service/:id", svcHandler.DeleteService)

	return svcHandler
}
