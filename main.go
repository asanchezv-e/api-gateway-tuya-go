// Package main contiene la configuración principal del API Gateway Tuya Go
// @title API Gateway Tuya Go - SDK Oficial
// @version 1.0
// @description API Gateway para Tuya Cloud siguiendo el patrón del SDK oficial
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
	_ "api-gateway-tuya-go/docs" // docs generado por swag
	"api-gateway-tuya-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run()
}