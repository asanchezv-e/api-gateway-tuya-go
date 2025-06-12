package models

// DeviceModel representa la estructura de un dispositivo IoT
// @Description Modelo completo de un dispositivo IoT
type DeviceModel struct {
	UUID   string `json:"uuid"`
	UID    string `json:"uid"`
	Name   string `json:"name"`
	IP     string `json:"ip"`
	Sub    bool   `json:"sub"`
	Model  string `json:"model"`
	Status []struct {
		Code  string      `json:"code"`
		Value interface{} `json:"value"`
	} `json:"status"`
	Category    string `json:"category"`
	Online      bool   `json:"online"`
	ID          string `json:"id"`
	TimeZone    string `json:"time_zone"`
	LocalKey    string `json:"local_key"`
	UpdateTime  int    `json:"update_time"`
	ActiveTime  int    `json:"active_time"`
	OwnerID     string `json:"owner_id"`
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
}

// DeviceStatus representa el estado de un dispositivo
// @Description Estado específico de un dispositivo con código y valor
type DeviceStatus struct {
	Code  string      `json:"code"`
	Value interface{} `json:"value"`
} 