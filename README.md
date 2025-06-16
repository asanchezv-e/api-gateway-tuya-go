# API Gateway Tuya Go - SDK Oficial

Un API Gateway completo que implementa exactamente el patrón del [SDK oficial de Tuya](https://github.com/tuya/tuya-cloud-sdk-go) para Go, proporcionando acceso REST a las funcionalidades principales de Tuya Cloud.

## 🎯 Objetivo

Proporcionar una implementación limpia y directa del SDK oficial de Tuya, eliminando complejidades innecesarias y siguiendo exactamente el patrón de uso documentado, con funcionalidades completas para gestión de tokens y dispositivos.

## 📋 Características

- ✅ Implementación fiel al SDK oficial de Tuya
- ✅ Método `GetToken()` como endpoint REST
- ✅ Método `GetDevice()` como endpoint REST
- ✅ Algoritmo de firma HMAC-SHA256 exacto del SDK oficial
- ✅ Gestión automática de tokens entre requests
- ✅ Documentación Swagger completa y automática
- ✅ Arquitectura limpia siguiendo patrones del SDK oficial
- ✅ Soporte para múltiples regiones de Tuya Cloud

## 🚀 Inicio Rápido

### 1. Configurar variables de entorno

```bash
cp config.example.env .env
```

Edita `.env` con tus credenciales de Tuya Cloud:

```env
TUYA_ACCESS_ID=tu_access_id_aqui
TUYA_ACCESS_KEY=tu_access_key_aqui
TUYA_API_URL=https://openapi.tuyaus.com
```

### 2. Ejecutar la aplicación

```bash
go run main.go
```

La aplicación estará disponible en: http://localhost:8080

## 📚 Endpoints Disponibles

### API Principal

- `GET /` - Información de la API
- `GET /ping` - Health check
- `GET /swagger/index.html` - Documentación Swagger interactiva

### Tuya Cloud API (Siguiendo SDK Oficial)

- `GET /api/tuya/token` - Obtener token de acceso (equivalente a `GetToken()`)
- `GET /api/tuya/device/{deviceId}` - Obtener información de dispositivo (equivalente a `GetDevice()`)

## 🔧 Equivalencia con SDK Oficial

### SDK Oficial:

```go
import (
    "github.com/tuya/tuya-cloud-sdk-go"
)

const (
    Host     = "https://openapi.tuyaus.com"
    ClientID = "your_client_id"
    Secret   = "your_secret"
    DeviceID = "device_id"
)

func main() {
    // Obtener token
    GetToken()

    // Obtener información de dispositivo
    GetDevice(DeviceID)
}
```

### Este API Gateway:

```bash
# Configuración vía variables de entorno
export TUYA_ACCESS_ID="your_client_id"
export TUYA_ACCESS_KEY="your_secret"
export TUYA_API_URL="https://openapi.tuyaus.com"

# Obtener token
curl -X GET "http://localhost:8080/api/tuya/token"

# Obtener información de dispositivo
curl -X GET "http://localhost:8080/api/tuya/device/your_device_id"
```

## 🌍 Regiones Soportadas

Configura `TUYA_API_URL` según tu región:

- **Estados Unidos**: `https://openapi.tuyaus.com`
- **Europa**: `https://openapi.tuyaeu.com`
- **China**: `https://openapi.tuyacn.com`
- **India**: `https://openapi.tuyain.com`

## 📄 Ejemplos de Respuestas

### Token Response

```json
{
  "success": true,
  "message": "Token de Tuya obtenido exitosamente",
  "data": {
    "result": {
      "access_token": "d656fbc73863d74361ae7633c474ecab",
      "expire_time": 7200,
      "refresh_token": "a5b80c2fa143f35dcc241f7f49435d6f",
      "uid": "bay17494665128652uT8"
    },
    "success": true,
    "t": 1749845475156,
    "tid": "8f149c57489211f0ad5f1a96caf0cb23"
  }
}
```

### Device Response

```json
{
  "success": true,
  "message": "Información del dispositivo obtenida exitosamente",
  "data": {
    "result": {
      "id": "device_id",
      "name": "Device Name",
      "online": true,
      "status": [
        {
          "code": "switch_1",
          "value": true
        }
      ]
    },
    "success": true,
    "t": 1749845483966,
    "tid": "9452f158489211f0aa735aed1aeae571"
  }
}
```

## 🛠️ Desarrollo

### Instalar dependencias

```bash
go mod tidy
```

### Generar documentación Swagger

```bash
# Instalar swag si no está instalado
go install github.com/swaggo/swag/cmd/swag@latest

# Generar documentación
~/go/bin/swag init
```

### Compilar

```bash
go build -o api-tuya-gateway .
```

### Ejecutar tests

```bash
# Probar endpoint de token
curl -X GET "http://localhost:8080/api/tuya/token"

# Probar endpoint de dispositivo
curl -X GET "http://localhost:8080/api/tuya/device/your_device_id"

# Verificar documentación Swagger
curl -X GET "http://localhost:8080/swagger/index.html"
```

## 🏗️ Arquitectura

```
api-gateway-tuya-go/
├── handlers/
│   ├── tuya_handler.go    # Lógica principal de endpoints
│   └── tuya_utils.go      # Utilidades de firma HMAC-SHA256
├── routes/
│   └── tuya_routes.go     # Definición de rutas
├── docs/                  # Documentación Swagger generada
├── main.go               # Punto de entrada de la aplicación
├── .env                  # Configuración de credenciales
└── README.md            # Esta documentación
```

## 🔐 Algoritmo de Firma

Implementa exactamente el mismo algoritmo de firma HMAC-SHA256 que el SDK oficial:

### Para Token Requests:

```
stringToSign = Method + "\n" + ContentSHA256 + "\n" + Headers + "\n" + URL
signStr = clientID + timestamp + stringToSign
signature = HMAC-SHA256(signStr, secret).toUpperCase()
```

### Para Device Requests:

```
stringToSign = Method + "\n" + ContentSHA256 + "\n" + Headers + "\n" + URL
signStr = clientID + accessToken + timestamp + stringToSign
signature = HMAC-SHA256(signStr, secret).toUpperCase()
```

## 📝 Notas Técnicas

- **Gestión de tokens**: Los tokens se mantienen automáticamente entre requests
- **Firma múltiple**: Maneja correctamente tanto token requests como device requests
- **Headers automáticos**: Configura automáticamente todos los headers requeridos
- **Validación**: Verifica respuestas y maneja errores apropiadamente
- **Documentación**: Swagger UI completamente funcional e interactiva

## 🚨 Solución de Problemas

### Error "sign invalid"

- Verifica que las credenciales `TUYA_ACCESS_ID` y `TUYA_ACCESS_KEY` sean correctas
- Asegúrate de usar la región correcta en `TUYA_API_URL`

### Error "No permission"

- Habilita el data center en tu cuenta de Tuya Cloud Platform
- Verifica que tu proyecto tenga permisos para acceder a la API de dispositivos

### Token requests múltiples

- ✅ **Resuelto**: El sistema maneja correctamente múltiples requests de token consecutivos

## 🔗 Referencias

- [SDK Oficial de Tuya para Go](https://github.com/tuya/tuya-cloud-sdk-go)
- [Documentación de Firma de Tuya](https://developer.tuya.com/en/docs/iot/new-singnature?id=Kbw0q34cs2e5g)
- [Documentación de Tuya Cloud API](https://developer.tuya.com/en/docs/cloud/)
- [Portal de Desarrolladores de Tuya](https://iot.tuya.com/)

---

Este proyecto mantiene la misma funcionalidad del SDK oficial pero expuesta como API REST para mayor flexibilidad de integración, con soporte completo para gestión de tokens y consulta de dispositivos.
