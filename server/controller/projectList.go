package controller

import (
	"dailyReport/models"
	"dailyReport/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func CreateProjectHandleFun(c *gin.Context) {
	var params models.ProjectList
	err := c.ShouldBindJSON(&params)
	if err != nil {
		log.Fatalln(err)
		return
	}
	createErr := service.CreateProject(&params)
	if createErr != nil {
		c.JSON(http.StatusOK, &models.Response{
			Code: 500,
			Msg:  createErr.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, &models.Response{
		Code: 200,
		Msg:  "success",
		Data: nil,
	})
}
func GetProjectListHandleFun(ctx *gin.Context) {
	projectList, err := service.GetProjectList()
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: projectList,
	})
}
func DeleteProjectItemHandleFun(c *gin.Context) {
	deleteId, err := strconv.Atoi(c.Query("deleteId"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	delError := service.DeleteProjectItem(deleteId)
	if delError != nil {
		c.JSON(http.StatusOK, models.Response{
			Code: 500,
			Msg:  delError.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: nil,
	})
}
func GetReportDetailsHandleFun(c *gin.Context) {
	reportDetails, err := service.GetReportDetails()
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Code: 200,
		Msg:  "success",
		Data: reportDetails,
	})
}
