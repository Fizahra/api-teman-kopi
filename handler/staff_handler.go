package handler

import (
	"api_temankopi/dto"
	"api_temankopi/helper"
	"api_temankopi/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StaffHandler interface {
	AddStaff(ctx *gin.Context)
	EditStaff(ctx *gin.Context)
	DeleteStaff(ctx *gin.Context)
	ViewStaff(ctx *gin.Context)
}

type staffHandler struct {
	staffService service.StaffService
}

func NewStaffHandler(staffService service.StaffService) *staffHandler {
	return &staffHandler{
		staffService: staffService,
	}
}

func (s *staffHandler) AddStaff(ctx *gin.Context) {
	var staff dto.StaffCreate
	err := ctx.ShouldBindJSON(&staff)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	staffData, err := s.staffService.AddStaff(staff)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request!", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", staffData)
	ctx.JSON(http.StatusOK, res)

}

func (s *staffHandler) EditStaff(ctx *gin.Context) {
	var editStaff dto.StaffUpdate
	err := ctx.ShouldBindJSON(&editStaff)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to find Staff", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	staffData, err := s.staffService.EditStaff(id, editStaff)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to update staff", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	response := helper.BuildResponse(true, "OK", staffData)
	ctx.JSON(http.StatusCreated, response)
	return

}

func (s *staffHandler) DeleteStaff(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	err := s.staffService.DeleteStaff(id)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get id", "Please enter a valid id", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response := helper.BuildResponse(true, "Staff hass been deleted!", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (s *staffHandler) ViewStaff(ctx *gin.Context) {
	staff, err := s.staffService.ViewStaff()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}
	response := helper.BuildResponse(true, "OK", staff)
	ctx.JSON(http.StatusOK, response)
}
