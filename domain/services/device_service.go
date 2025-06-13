package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-gin-project/domain/interfaces"
	"go-gin-project/models"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"context"
	"errors"

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

func GetTuyaToken(clientID, clientSecret, apiURL string) (*models.TuyaTokenResponse, error) {
    method := "GET"
    body := []byte("")
    ts := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	contentSha256 := sha256.New()
    contentSha256.Write(body)
    bodyHash := hex.EncodeToString(contentSha256.Sum(nil))
    stringToSign := method + "\n" + bodyHash + "\n\n/v1.0/token?grant_type=1"
    signStr := clientID + ts + stringToSign
    mac := hmac.New(sha256.New, []byte(clientSecret))
    mac.Write([]byte(signStr))
    sign := strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))

    url := fmt.Sprintf("%s/v1.0/token?grant_type=1", apiURL)
    req, _ := http.NewRequest(method, url, nil)
    req.Header.Set("client_id", clientID)
    req.Header.Set("sign_method", "HMAC-SHA256")
    req.Header.Set("t", ts)
    req.Header.Set("sign", sign)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("[TuyaToken] Error en la petición HTTP:", err)
        return nil, err
    }
    defer resp.Body.Close()
    bodyResp, _ := io.ReadAll(resp.Body)

    fmt.Println("[TuyaToken] Código de estado HTTP:", resp.StatusCode)
    fmt.Println("[TuyaToken] Respuesta cruda:", string(bodyResp))

    var tokenResp models.TuyaTokenResponse
    if err := json.Unmarshal(bodyResp, &tokenResp); err != nil {
        fmt.Println("[TuyaToken] Error al parsear JSON:", err)
        return nil, err
    }
    if !tokenResp.Success {
        fmt.Println("[TuyaToken] Error: respuesta no exitosa de Tuya", tokenResp.Msg)
        return &tokenResp, fmt.Errorf("no se pudo obtener token Tuya: %s", tokenResp.Msg)
    }
    return &tokenResp, nil
}