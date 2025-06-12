package repositories

import (
	"context"
	"go-gin-project/domain/interfaces"
	"go-gin-project/models"
)

// MemoryDeviceRepository implementa DeviceRepository usando memoria
// Principio SOLID: Dependency Inversion Principle (DIP) - implementación concreta
type MemoryDeviceRepository struct {
	devices []models.DeviceModel
}

// NewMemoryDeviceRepository crea un nuevo repositorio en memoria con datos de ejemplo
func NewMemoryDeviceRepository() interfaces.DeviceRepository {
	// Datos de ejemplo para demostración
	devices := []models.DeviceModel{
		{
			UUID:        "uuid-001",
			UID:         "uid-001",
			Name:        "Sensor de Temperatura Sala 1",
			IP:          "192.168.1.101",
			Sub:         false,
			Model:       "TempSensor-v2",
			Status: []struct {
				Code  string      `json:"code"`
				Value interface{} `json:"value"`
			}{
				{Code: "temperature", Value: 23.5},
				{Code: "humidity", Value: 65},
			},
			Category:    "sensor",
			Online:      true,
			ID:          "device-001",
			TimeZone:    "UTC-5",
			LocalKey:    "localkey123",
			UpdateTime:  1640995200,
			ActiveTime:  1640995200,
			OwnerID:     "owner-001",
			ProductID:   "prod-temp-001",
			ProductName: "Sensor de Temperatura Inteligente",
		},
		{
			UUID:        "uuid-002",
			UID:         "uid-002",
			Name:        "Cámara de Seguridad",
			IP:          "192.168.1.102",
			Sub:         true,
			Model:       "SecurityCam-HD",
			Status: []struct {
				Code  string      `json:"code"`
				Value interface{} `json:"value"`
			}{
				{Code: "recording", Value: true},
				{Code: "motion_detected", Value: false},
			},
			Category:    "security",
			Online:      true,
			ID:          "device-002",
			TimeZone:    "UTC-5",
			LocalKey:    "localkey456",
			UpdateTime:  1640995300,
			ActiveTime:  1640995300,
			OwnerID:     "owner-001",
			ProductID:   "prod-cam-001",
			ProductName: "Cámara de Seguridad HD",
		},
		{
			UUID:        "uuid-003",
			UID:         "uid-003",
			Name:        "Interruptor Inteligente",
			IP:          "192.168.1.103",
			Sub:         false,
			Model:       "SmartSwitch-v1",
			Status: []struct {
				Code  string      `json:"code"`
				Value interface{} `json:"value"`
			}{
				{Code: "switch", Value: false},
				{Code: "power_consumption", Value: 0},
			},
			Category:    "switch",
			Online:      false,
			ID:          "device-003",
			TimeZone:    "UTC-5",
			LocalKey:    "localkey789",
			UpdateTime:  1640994800,
			ActiveTime:  1640994800,
			OwnerID:     "owner-002",
			ProductID:   "prod-switch-001",
			ProductName: "Interruptor Inteligente WiFi",
		},
	}

	return &MemoryDeviceRepository{
		devices: devices,
	}
}

// GetAll obtiene todos los dispositivos
func (r *MemoryDeviceRepository) GetAll(ctx context.Context) ([]models.DeviceModel, error) {
	return r.devices, nil
}

// GetByID obtiene un dispositivo por su ID
func (r *MemoryDeviceRepository) GetByID(ctx context.Context, id string) (*models.DeviceModel, error) {
	for _, device := range r.devices {
		if device.ID == id {
			return &device, nil
		}
	}
	return nil, nil
}

// GetByOwnerID obtiene dispositivos por ID del propietario
func (r *MemoryDeviceRepository) GetByOwnerID(ctx context.Context, ownerID string) ([]models.DeviceModel, error) {
	var devices []models.DeviceModel
	for _, device := range r.devices {
		if device.OwnerID == ownerID {
			devices = append(devices, device)
		}
	}
	return devices, nil
}

// GetOnlineDevices obtiene solo los dispositivos en línea
func (r *MemoryDeviceRepository) GetOnlineDevices(ctx context.Context) ([]models.DeviceModel, error) {
	var onlineDevices []models.DeviceModel
	for _, device := range r.devices {
		if device.Online {
			onlineDevices = append(onlineDevices, device)
		}
	}
	return onlineDevices, nil
} 