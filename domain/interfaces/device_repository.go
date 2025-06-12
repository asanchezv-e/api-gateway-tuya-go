package interfaces

import (
	"context"
	"go-gin-project/models"
)

// DeviceRepository define la interfaz para el repositorio de dispositivos
// Principio SOLID: Dependency Inversion Principle (DIP)
type DeviceRepository interface {
	GetAll(ctx context.Context) ([]models.DeviceModel, error)
	GetByID(ctx context.Context, id string) (*models.DeviceModel, error)
	GetByOwnerID(ctx context.Context, ownerID string) ([]models.DeviceModel, error)
	GetOnlineDevices(ctx context.Context) ([]models.DeviceModel, error)
} 