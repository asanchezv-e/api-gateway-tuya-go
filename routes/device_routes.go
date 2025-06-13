package routes

import (
	"go-gin-project/handlers"

	"github.com/gin-gonic/gin"
)

// SetupDeviceRoutes configura las rutas para dispositivos
// Principio SOLID: Open/Closed Principle (OCP) - extensible sin modificar código existente
func SetupDeviceRoutes(router *gin.Engine, deviceHandler *handlers.DeviceHandler) {
	// Grupo de rutas para API de dispositivos
	api := router.Group("/api")
	{
		devices := api.Group("/devices")
		{
			devices.GET("", deviceHandler.GetAllDevices)                    // GET /api/devices
			devices.GET("/online", deviceHandler.GetOnlineDevices)           // GET /api/devices/online
			devices.GET("/:id", deviceHandler.GetDeviceByID)                // GET /api/devices/:id
			devices.GET("/owner/:ownerId", deviceHandler.GetDevicesByOwner) // GET /api/devices/owner/:ownerId
		}
		api.GET("/tuya/token", deviceHandler.GetTuyaToken) // GET /api/tuya/token
	}
}