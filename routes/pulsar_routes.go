package routes

import (
	"api-gateway-tuya-go/handlers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Variable global para el handler de Pulsar
var pulsarHandler *handlers.PulsarHandler

// SetupPulsarRoutes configura las rutas del listener de Pulsar
func SetupPulsarRoutes(r *gin.Engine, handler *handlers.PulsarHandler) {
	pulsarHandler = handler
	
	// Grupo de rutas para Pulsar
	pulsarGroup := r.Group("/api/pulsar")
	{
		pulsarGroup.GET("/status", GetPulsarStatus)
	}
}

// GetPulsarStatus obtiene el estado del listener de Pulsar
// @Summary Obtener estado del listener de Pulsar
// @Description Obtiene información detallada del estado actual del listener de Pulsar incluyendo configuración y conectividad. El listener se inicia automáticamente con el servidor.
// @Tags Pulsar
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Estado detallado del listener con configuración y métricas"
// @Router /pulsar/status [get]
func GetPulsarStatus(c *gin.Context) {
	if pulsarHandler == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Handler de Pulsar no inicializado",
			"status":  "not_initialized",
		})
		return
	}

	status := pulsarHandler.GetStatus()
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Estado del listener de Pulsar (se inicia automáticamente con el servidor)",
		"status":  status,
	})
}

func maskCredential(credential string) string {
	if len(credential) <= 8 {
		return strings.Repeat("*", len(credential))
	}
	return credential[:4] + strings.Repeat("*", len(credential)-8) + credential[len(credential)-4:]
} 