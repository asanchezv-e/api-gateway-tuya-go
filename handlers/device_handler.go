package handlers

import (
	"go-gin-project/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeviceHandler maneja las peticiones HTTP para dispositivos
// Principio SOLID: Single Responsibility Principle (SRP)
type DeviceHandler struct {
	deviceService *services.DeviceService
}

// NewDeviceHandler crea una nueva instancia del handler de dispositivos
func NewDeviceHandler(deviceService *services.DeviceService) *DeviceHandler {
	return &DeviceHandler{
		deviceService: deviceService,
	}
}

// APIResponse estructura estándar para respuestas de la API
// @Description Estructura estándar de respuesta de la API
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// GetAllDevices maneja GET /api/devices
// @Summary Obtener todos los dispositivos
// @Description Obtiene la lista completa de dispositivos registrados en el sistema
// @Tags dispositivos
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse{data=[]models.DeviceModel} "Lista de dispositivos obtenida exitosamente"
// @Failure 500 {object} APIResponse "Error interno del servidor"
// @Router /devices [get]
func (h *DeviceHandler) GetAllDevices(c *gin.Context) {
	ctx := c.Request.Context()
	
	devices, err := h.deviceService.GetAllDevices(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "Error al obtener dispositivos",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Dispositivos obtenidos exitosamente",
		Data:    devices,
	})
}

// GetDeviceByID maneja GET /api/devices/:id
// @Summary Obtener dispositivo por ID
// @Description Obtiene un dispositivo específico usando su identificador único
// @Tags dispositivos
// @Accept json
// @Produce json
// @Param id path string true "ID del dispositivo"
// @Success 200 {object} APIResponse{data=models.DeviceModel} "Dispositivo obtenido exitosamente"
// @Failure 400 {object} APIResponse "ID de dispositivo inválido"
// @Failure 404 {object} APIResponse "Dispositivo no encontrado"
// @Failure 500 {object} APIResponse "Error interno del servidor"
// @Router /devices/{id} [get]
func (h *DeviceHandler) GetDeviceByID(c *gin.Context) {
	ctx := c.Request.Context()
	deviceID := c.Param("id")

	device, err := h.deviceService.GetDeviceByID(ctx, deviceID)
	if err != nil {
		if err == services.ErrDeviceNotFound {
			c.JSON(http.StatusNotFound, APIResponse{
				Success: false,
				Message: "Dispositivo no encontrado",
				Error:   err.Error(),
			})
			return
		}
		if err == services.ErrInvalidID {
			c.JSON(http.StatusBadRequest, APIResponse{
				Success: false,
				Message: "ID de dispositivo inválido",
				Error:   err.Error(),
			})
			return
		}
		
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "Error al obtener dispositivo",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Dispositivo obtenido exitosamente",
		Data:    device,
	})
}

// GetDevicesByOwner maneja GET /api/devices/owner/:ownerId
// @Summary Obtener dispositivos por propietario
// @Description Obtiene todos los dispositivos que pertenecen a un propietario específico
// @Tags dispositivos
// @Accept json
// @Produce json
// @Param ownerId path string true "ID del propietario"
// @Success 200 {object} APIResponse{data=[]models.DeviceModel} "Dispositivos del propietario obtenidos exitosamente"
// @Failure 400 {object} APIResponse "ID de propietario inválido"
// @Failure 500 {object} APIResponse "Error interno del servidor"
// @Router /devices/owner/{ownerId} [get]
func (h *DeviceHandler) GetDevicesByOwner(c *gin.Context) {
	ctx := c.Request.Context()
	ownerID := c.Param("ownerId")

	devices, err := h.deviceService.GetDevicesByOwner(ctx, ownerID)
	if err != nil {
		if err == services.ErrInvalidID {
			c.JSON(http.StatusBadRequest, APIResponse{
				Success: false,
				Message: "ID de propietario inválido",
				Error:   err.Error(),
			})
			return
		}
		
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "Error al obtener dispositivos del propietario",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Dispositivos del propietario obtenidos exitosamente",
		Data:    devices,
	})
}

// GetOnlineDevices maneja GET /api/devices/online
// @Summary Obtener dispositivos en línea
// @Description Obtiene únicamente los dispositivos que están actualmente en línea
// @Tags dispositivos
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse{data=[]models.DeviceModel} "Dispositivos en línea obtenidos exitosamente"
// @Failure 500 {object} APIResponse "Error interno del servidor"
// @Router /devices/online [get]
func (h *DeviceHandler) GetOnlineDevices(c *gin.Context) {
	ctx := c.Request.Context()
	
	devices, err := h.deviceService.GetOnlineDevices(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "Error al obtener dispositivos en línea",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Dispositivos en línea obtenidos exitosamente",
		Data:    devices,
	})
} 