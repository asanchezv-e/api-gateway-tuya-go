// Package main contiene la configuración principal de la API de Dispositivos IoT
// @title API de Dispositivos IoT
// @version 1.0
// @description API REST para gestión de dispositivos IoT con arquitectura limpia
// @termsOfService http://swagger.io/terms/
// @contact.name Soporte API
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api
package main

import (
	_ "go-gin-project/docs" // docs generado por swag
	"go-gin-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
	routes.SetupRoutes(r)
  r.Run()
}