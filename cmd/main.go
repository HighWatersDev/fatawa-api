package main

import (
	"context"
	"fatawa-api/pkg/controllers"
	"fatawa-api/pkg/db"
	"fatawa-api/pkg/routes"
	"fatawa-api/pkg/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	server               *gin.Engine
	fatwaService         services.FatwaService
	FatwaController      controllers.FatwaController
	fatwaCollection      *mongo.Collection
	FatwaRouteController routes.FatwaRouteController
)

func init() {

	ctx := context.TODO()
	db.ConnectDB()

	fatwaCollection = db.MI.Collection
	fatwaService = services.NewFatwaService(fatwaCollection, ctx)
	FatwaController = controllers.NewFatwaController(fatwaService)
	FatwaRouteController = routes.NewFatwaControllerRoute(FatwaController)

	server = gin.Default()
}

func startGinServer() {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowAllOrigins = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api/v1")

	FatwaRouteController.FatwaRoute(router)
	log.Fatal(server.Run(":3000"))
}

func main() {
	startGinServer()
}
