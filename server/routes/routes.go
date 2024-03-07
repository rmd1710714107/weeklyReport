package routes

import (
	"dailyReport/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() error {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	router.POST("/newProject", controller.CreateProjectHandleFun)
	router.GET("/getProjectList", controller.GetProjectListHandleFun)
	router.DELETE("/delProjectList", controller.DeleteProjectItemHandleFun)
	router.GET("/getReportDetails", controller.GetReportDetailsHandleFun)
	err := router.Run(":3213")
	if err != nil {
		return err
	}
	return nil
}
