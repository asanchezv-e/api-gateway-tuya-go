package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// TuyaHandler maneja las peticiones relacionadas con Tuya Cloud API
type TuyaHandler struct{}

// NewTuyaHandler crea una nueva instancia del handler de Tuya
func NewTuyaHandler() *TuyaHandler {
	return &TuyaHandler{}
}

// TuyaTokenResponse estructura de respuesta del token según el SDK oficial
type TuyaTokenResponse struct {
	Result struct {
		AccessToken  string `json:"access_token"`
		ExpireTime   int    `json:"expire_time"`
		RefreshToken string `json:"refresh_token"`
		UID          string `json:"uid"`
	} `json:"result"`
	Success bool   `json:"success"`
	T       int64  `json:"t"`
	TID     string `json:"tid"`
}

// TuyaDeviceResponse estructura de respuesta para información de dispositivo
type TuyaDeviceResponse struct {
	Result  interface{} `json:"result"`
	Success bool        `json:"success"`
	T       int64       `json:"t"`
	TID     string      `json:"tid"`
}

// TuyaAPIResponse estructura estándar para respuestas de la API de Tuya
type TuyaAPIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Debug   interface{} `json:"debug,omitempty"`
}

// GetTuyaToken maneja GET /api/tuya/token siguiendo el patrón del SDK oficial
// @Summary Obtener token de acceso de Tuya
// @Description Obtiene un token de acceso de Tuya Cloud API usando las credenciales configuradas
// @Tags tuya
// @Accept json
// @Produce json
// @Success 200 {object} TuyaAPIResponse{data=TuyaTokenResponse} "Token obtenido exitosamente"
// @Failure 500 {object} TuyaAPIResponse "Error al obtener token"
// @Router /tuya/token [get]
func (h *TuyaHandler) GetTuyaToken(c *gin.Context) {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(".env"); err != nil {
		// Si no existe .env, intenta cargar config.example.env
		_ = godotenv.Load("config.example.env")
	}
	
	accessID := os.Getenv("TUYA_ACCESS_ID")
	accessKey := os.Getenv("TUYA_ACCESS_KEY")
	apiURL := os.Getenv("TUYA_API_URL")

	if accessID == "" || accessKey == "" || apiURL == "" {
		c.JSON(http.StatusInternalServerError, TuyaAPIResponse{
			Success: false,
			Message: "Configuración de Tuya incompleta",
			Error:   "TUYA_ACCESS_ID, TUYA_ACCESS_KEY y TUYA_API_URL son requeridos",
		})
		return
	}

	// Usar el método del SDK oficial
	tokenResponse, err := getTuyaToken(accessID, accessKey, apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TuyaAPIResponse{
			Success: false,
			Message: "Error al obtener token de Tuya",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, TuyaAPIResponse{
		Success: true,
		Message: "Token de Tuya obtenido exitosamente",
		Data:    tokenResponse,
	})
}

// GetTuyaDevice maneja GET /api/tuya/device/:deviceId siguiendo el patrón del SDK oficial
// @Summary Obtener información de dispositivo
// @Description Obtiene información de un dispositivo específico por su ID
// @Tags tuya
// @Accept json
// @Produce json
// @Param deviceId path string true "ID del dispositivo"
// @Success 200 {object} TuyaAPIResponse{data=TuyaDeviceResponse} "Información del dispositivo obtenida exitosamente"
// @Failure 400 {object} TuyaAPIResponse "ID de dispositivo requerido"
// @Failure 500 {object} TuyaAPIResponse "Error al obtener información del dispositivo"
// @Router /tuya/device/{deviceId} [get]
func (h *TuyaHandler) GetTuyaDevice(c *gin.Context) {
	deviceID := c.Param("deviceId")
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, TuyaAPIResponse{
			Success: false,
			Message: "ID de dispositivo requerido",
			Error:   "El parámetro deviceId es obligatorio",
		})
		return
	}

	// Cargar variables de entorno
	if err := godotenv.Load(".env"); err != nil {
		_ = godotenv.Load("config.example.env")
	}
	
	accessID := os.Getenv("TUYA_ACCESS_ID")
	accessKey := os.Getenv("TUYA_ACCESS_KEY")
	apiURL := os.Getenv("TUYA_API_URL")

	if accessID == "" || accessKey == "" || apiURL == "" {
		c.JSON(http.StatusInternalServerError, TuyaAPIResponse{
			Success: false,
			Message: "Configuración de Tuya incompleta",
			Error:   "TUYA_ACCESS_ID, TUYA_ACCESS_KEY y TUYA_API_URL son requeridos",
		})
		return
	}

	// Asegurar que tenemos un token válido
	if Token == "" {
		_, err := getTuyaToken(accessID, accessKey, apiURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, TuyaAPIResponse{
				Success: false,
				Message: "Error al obtener token para consultar dispositivo",
				Error:   err.Error(),
			})
			return
		}
	}

	// Usar el método del SDK oficial para obtener dispositivo
	deviceResponse, err := getTuyaDevice(deviceID, accessID, accessKey, apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TuyaAPIResponse{
			Success: false,
			Message: "Error al obtener información del dispositivo",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, TuyaAPIResponse{
		Success: true,
		Message: "Información del dispositivo obtenida exitosamente",
		Data:    deviceResponse,
	})
}

// getTuyaToken implementa la función GetToken del SDK oficial
func getTuyaToken(clientID, secret, host string) (*TuyaTokenResponse, error) {
	method := "GET"
	body := []byte(``)
	req, err := http.NewRequest(method, host+"/v1.0/token?grant_type=1", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error creando request: %v", err)
	}

	buildHeader(req, body, clientID, secret)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando request: %v", err)
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo respuesta: %v", err)
	}

	var tokenResponse TuyaTokenResponse
	if err := json.Unmarshal(bs, &tokenResponse); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	if !tokenResponse.Success {
		return nil, fmt.Errorf("respuesta no exitosa de Tuya: %s", string(bs))
	}

	// Guardar token globalmente como en el SDK oficial
	if v := tokenResponse.Result.AccessToken; v != "" {
		Token = v
	}

	return &tokenResponse, nil
}

// getTuyaDevice implementa la función GetDevice del SDK oficial
func getTuyaDevice(deviceID, clientID, secret, host string) (*TuyaDeviceResponse, error) {
	method := "GET"
	body := []byte(``)
	req, err := http.NewRequest(method, host+"/v1.0/devices/"+deviceID, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error creando request: %v", err)
	}

	buildHeader(req, body, clientID, secret)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando request: %v", err)
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo respuesta: %v", err)
	}

	var deviceResponse TuyaDeviceResponse
	if err := json.Unmarshal(bs, &deviceResponse); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	if !deviceResponse.Success {
		return nil, fmt.Errorf("respuesta no exitosa de Tuya: %s", string(bs))
	}

	return &deviceResponse, nil
} 