package main

import (
	"os"

	"github.com/giankas/AZIENDALE/routes"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Carica le route API
	routes.RegisterRoutes(router)

	// Servi file statici frontend dalla cartella "frontend"
	router.Use(static.Serve("/", static.LocalFile("./frontend", true)))

	// Avvia server sulla porta specificata da Railway o default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
