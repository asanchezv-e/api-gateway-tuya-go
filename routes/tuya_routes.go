package routes

import (
	"api-gateway-tuya-go/handlers"

	"github.com/gin-gonic/gin"
)

// SetupTuyaRoutes configura las rutas para Tuya Cloud API
// Siguiendo el patrón del SDK oficial de Tuya
func SetupTuyaRoutes(r *gin.Engine, tuyaHandler *handlers.TuyaHandler) {
	// Grupo de rutas para Tuya Cloud API
	tuyaGroup := r.Group("/api/tuya")
	{
		// Endpoint principal para obtener token - equivalente a token.GetTokenAPI()
		tuyaGroup.GET("/token", tuyaHandler.GetTuyaToken)
		
		// Endpoint para obtener información de dispositivo - equivalente a GetDevice()
		tuyaGroup.GET("/device/:deviceId", tuyaHandler.GetTuyaDevice)
	}
} 