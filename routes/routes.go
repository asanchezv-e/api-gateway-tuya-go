package routes

import (
	"api-gateway-tuya-go/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	// Rutas básicas de salud de la API
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API Gateway Tuya Go - SDK Oficial",
			"version": "1.0.0",
			"status":  "online",
			"swagger": "http://localhost:8080/swagger/index.html",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Configurar Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configurar handler de Tuya siguiendo el patrón del SDK oficial
	tuyaHandler := handlers.NewTuyaHandler()
	
	// Configurar rutas de Tuya - equivalente a usar config.SetEnv() + token.GetTokenAPI()
	SetupTuyaRoutes(r, tuyaHandler)
}