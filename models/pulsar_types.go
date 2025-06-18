package models

import "time"

// PulsarMessage representa la estructura base de un mensaje de Pulsar
type PulsarMessage struct {
	Data     string `json:"data"`
	Protocol int    `json:"protocol"`
	PV       string `json:"pv"`
	Sign     string `json:"sign"`
	T        int64  `json:"t"`
}

// DeviceStatusChange representa el cambio de estado decodificado de un dispositivo
type DeviceStatusChange struct {
	DataID     string        `json:"dataId"`
	DevID      string        `json:"devId"`
	ProductKey string        `json:"productKey"`
	Status     []StatusItem  `json:"status"`
}

// StatusItem representa un elemento de estado individual
type StatusItem struct {
	Code  string      `json:"code"`
	T     int64       `json:"t"`
	Value interface{} `json:"value"`
}

// DeviceEventLog estructura para logging de eventos de dispositivos
type DeviceEventLog struct {
	Timestamp   time.Time   `json:"timestamp"`
	DeviceID    string      `json:"device_id"`
	ProductKey  string      `json:"product_key"`
	EventType   string      `json:"event_type"`
	StatusCode  string      `json:"status_code,omitempty"`
	OldValue    interface{} `json:"old_value,omitempty"`
	NewValue    interface{} `json:"new_value"`
	Message     string      `json:"message"`
}

// PulsarConfig configuración para el cliente de Pulsar
type PulsarConfig struct {
	AccessID    string `json:"access_id"`
	AccessKey   string `json:"access_key"`
	PulsarAddr  string `json:"pulsar_addr"`
	Topic       string `json:"topic"`
	DebugMode   bool   `json:"debug_mode"`
} 