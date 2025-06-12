package config

import (
	"go-gin-project/domain/interfaces"
	"go-gin-project/domain/services"
	"go-gin-project/handlers"
	"go-gin-project/infrastructure/repositories"
)

// Container maneja la inyección de dependencias
// Principio SOLID: Dependency Inversion Principle (DIP)
type Container struct {
	DeviceRepository interfaces.DeviceRepository
	DeviceService    *services.DeviceService
	DeviceHandler    *handlers.DeviceHandler
}

// NewContainer crea un nuevo contenedor con todas las dependencias configuradas
func NewContainer() *Container {
	// Crear repositorio (se puede cambiar fácilmente por una implementación de base de datos)
	deviceRepo := repositories.NewMemoryDeviceRepository()
	
	// Crear servicio inyectando el repositorio
	deviceService := services.NewDeviceService(deviceRepo)
	
	// Crear handler inyectando el servicio
	deviceHandler := handlers.NewDeviceHandler(deviceService)
	
	return &Container{
		DeviceRepository: deviceRepo,
		DeviceService:    deviceService,
		DeviceHandler:    deviceHandler,
	}
} 