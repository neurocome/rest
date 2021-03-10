package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/neurocome/rest/controller"
	"github.com/neurocome/rest/middlewares"
	"github.com/neurocome/rest/services"
	// gindump "github.com/tpkeeper/gin-dump"
)

var (
	employeService    services.EmployeService      = services.New()
	employeController controller.EmployeController = controller.New(employeService)
)

func setLogOut() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setLogOut()
	server := gin.New()
	// server.Use(gin.Recovery(), middlewares.Logger(), middlewares.Auth(),gindump.Dump())

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.Auth())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/employe", func(ctx *gin.Context) {
			ctx.JSON(200, employeController.FindAll())
		})

		apiRoutes.POST("/employe", func(ctx *gin.Context) {
			err := employeController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input Success"})
			}
		})
	}

	// viewRoutes := server.Group("/view")
	// {
	// 	viewRoutes.GET("/employe", employeController.ShowAll)
	// }

	server.Run(":3000")
}
