package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/kittanutp/salesrecorder/config"
	"github.com/kittanutp/salesrecorder/database"
	"github.com/kittanutp/salesrecorder/middleware"
	"github.com/kittanutp/salesrecorder/service"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: config.Origins,
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	router.SetTrustedProxies(config.Origins)

	crtl := service.DBController{Database: database.Connect()}

	adminRoutes := router.Group(("api/admin"))
	adminRoutes.Use(middleware.AuthApp())
	{
		adminRoutes.GET("all-item", crtl.GetItems)

	}

	userRoutes := router.Group("api/user")
	userRoutes.Use(middleware.AuthApp())
	{
		userRoutes.POST("new", crtl.CreateUser)
		userRoutes.POST("login", crtl.LogIn)
	}

	userAuthRoutes := router.Group("api/user")
	userAuthRoutes.Use(middleware.AuthUser())
	{
		userAuthRoutes.GET("info", crtl.GetUserInfo)
	}

	itemRoutes := router.Group("api/item")
	itemRoutes.Use(middleware.AuthUser())
	{
		itemRoutes.GET("user", crtl.GetItemsByUser)
		itemRoutes.POST("create", crtl.CreateItem)
	}

	saleRoutes := router.Group("api/sale")
	saleRoutes.Use(middleware.AuthUser())
	{
		// 	saleRoutes.GET("get-sales", service.GetSaleItems)
		saleRoutes.POST("add-sale", crtl.AddSale)
	}

	router.Run("localhost:8000")
}
