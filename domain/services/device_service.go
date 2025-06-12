package services

import (
	"context"
	"errors"
	"go-gin-project/domain/interfaces"
	"go-gin-project/models"
)

var (
	ErrDeviceNotFound = errors.New("dispositivo no encontrado")
	ErrInvalidID      = errors.New("ID de dispositivo inválido")
)

// DeviceService implementa la lógica de negocio para dispositivos
// Principio SOLID: Single Responsibility Principle (SRP)
type DeviceService struct {
	deviceRepo interfaces.DeviceRepository
}

// NewDeviceService crea una nueva instancia del servicio de dispositivos
func NewDeviceService(deviceRepo interfaces.DeviceRepository) *DeviceService {
	return &DeviceService{
		deviceRepo: deviceRepo,
	}
}

// GetAllDevices obtiene todos los dispositivos
func (s *DeviceService) GetAllDevices(ctx context.Context) ([]models.DeviceModel, error) {
	devices, err := s.deviceRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	
	// Aquí podríamos aplicar reglas de negocio adicionales
	// como filtrado, ordenamiento, etc.
	
	return devices, nil
}

// GetDeviceByID obtiene un dispositivo por su ID
func (s *DeviceService) GetDeviceByID(ctx context.Context, id string) (*models.DeviceModel, error) {
	if id == "" {
		return nil, ErrInvalidID
	}
	
	device, err := s.deviceRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	if device == nil {
		return nil, ErrDeviceNotFound
	}
	
	return device, nil
}

// GetDevicesByOwner obtiene dispositivos por ID del propietario
func (s *DeviceService) GetDevicesByOwner(ctx context.Context, ownerID string) ([]models.DeviceModel, error) {
	if ownerID == "" {
		return nil, ErrInvalidID
	}
	
	return s.deviceRepo.GetByOwnerID(ctx, ownerID)
}

// GetOnlineDevices obtiene solo los dispositivos en línea
func (s *DeviceService) GetOnlineDevices(ctx context.Context) ([]models.DeviceModel, error) {
	return s.deviceRepo.GetOnlineDevices(ctx)
} 