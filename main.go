package main

import (
	"net/http"
	"trial-class/config"
	"trial-class/controller"

	_ "trial-class/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @version         1.0
// @description     Dokomentasi REST API project Mini Ecommerce II

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:8000

func main() {
	s := gin.Default()
	config.DBConnect()

	// route
	s.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world")
	})

	// list semua produk
	s.GET("/products", controller.HandlerGetProducts)

	// create order
	s.POST("/orders", controller.HandlerPostOrder)

	// list order
	s.GET("/orders", controller.HandlerGetOrders)

	// swagger docs
	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.Run(":8000")
}
