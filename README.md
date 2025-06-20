# API Gateway Tuya Go - SDK Oficial con Listener Pulsar

Un API Gateway completo que implementa exactamente el patrón del [SDK oficial de Tuya](https://github.com/tuya/tuya-cloud-sdk-go) para Go, proporcionando acceso REST a las funcionalidades principales de Tuya Cloud con un **listener de mensajes Pulsar en tiempo real** para cambios de estado de dispositivos.

## 🎯 Objetivo

Proporcionar una implementación limpia y directa del SDK oficial de Tuya, eliminando complejidades innecesarias y siguiendo exactamente el patrón de uso documentado, con funcionalidades completas para gestión de tokens, dispositivos y **monitoreo en tiempo real de cambios de estado**.

## 📋 Características

### 🔧 API Gateway REST

- ✅ Implementación fiel al SDK oficial de Tuya
- ✅ Método `GetToken()` como endpoint REST
- ✅ Método `GetDevice()` como endpoint REST
- ✅ Algoritmo de firma HMAC-SHA256 exacto del SDK oficial
- ✅ Gestión automática de tokens entre requests
- ✅ Soporte para múltiples regiones de Tuya Cloud

### 🎧 Listener Pulsar en Tiempo Real

- ✅ **Listener automático de mensajes Pulsar**
- ✅ **Monitoreo en tiempo real de cambios de estado de dispositivos**
- ✅ **Configuración de ambientes (testing/production)**
- ✅ **Visualización intuitiva con emojis y formato estructurado**
- ✅ **Decodificación automática de mensajes AES**
- ✅ **Reconocimiento inteligente de tipos de estado**
- ✅ **Logs detallados con timestamps**

### 📚 Documentación y Herramientas

- ✅ Documentación Swagger completa y automática
- ✅ Arquitectura limpia siguiendo patrones del SDK oficial
- ✅ Scripts de generación de documentación

## 🚀 Inicio Rápido

### 1. Configurar variables de entorno

```bash
cp config.example.env .env
```

Edita `.env` con tus credenciales de Tuya Cloud:

```env
# Credenciales de Tuya Cloud
TUYA_ACCESS_ID=tu_access_id_aqui
TUYA_ACCESS_KEY=tu_access_key_aqui

# URL de la API de Tuya según tu región
TUYA_API_URL=https://openapi.tuyaeu.com

# Configuración de Pulsar para listener de mensajes
TUYA_PULSAR_ADDR=pulsar+ssl://mqe.tuyaeu.com:7285

# Ambiente de ejecución (testing/production)
TUYA_ENVIRONMENT=testing
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

### Pulsar Listener API

- `GET /api/pulsar/status` - Estado detallado del listener de Pulsar (automático)

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

Configura `TUYA_API_URL` y `TUYA_PULSAR_ADDR` según tu región:

### API URLs:

- **Estados Unidos**: `https://openapi.tuyaus.com`
- **Europa**: `https://openapi.tuyaeu.com`
- **China**: `https://openapi.tuyacn.com`
- **India**: `https://openapi.tuyain.com`

### Pulsar URLs:

- **Estados Unidos**: `pulsar+ssl://mqe.tuyaus.com:7285`
- **Europa**: `pulsar+ssl://mqe.tuyaeu.com:7285`
- **China**: `pulsar+ssl://mqe.tuyacn.com:7285`

## 🎧 Listener de Pulsar en Tiempo Real

### ¿Qué es el Listener de Pulsar?

El listener de Pulsar es un componente que **se conecta automáticamente al servicio de mensajería de Tuya Cloud** para recibir notificaciones en tiempo real cuando los estados de tus dispositivos cambian. Esto significa que tu aplicación será notificada instantáneamente cuando:

- Un interruptor se enciende o apaga
- Un sensor detecta movimiento
- La temperatura de un termostato cambia
- La batería de un dispositivo se agota
- Cualquier otro cambio de estado ocurre

### 🚀 Inicio Automático

El listener **se inicia automáticamente** cuando ejecutas el servidor:

```bash
go run main.go
```

Salida esperada:

```
🚀 Iniciando API Gateway Tuya Go con Listener Pulsar...
🎧 Iniciando listener de Pulsar automáticamente...
🌍 Ambiente: TESTING
📡 Conectando a: pulsar+ssl://mqe.tuyaeu.com:7285
📝 Topic: persistent://your_access_id/out/event-test
✅ Consumer creado exitosamente
📡 Esperando cambios de estado...
🌐 Servidor HTTP iniciado en puerto 8080
```

### 🔧 Configuración de Ambientes

El sistema soporta dos ambientes:

#### 🧪 Testing (Recomendado para desarrollo)

```env
TUYA_ENVIRONMENT=testing
```

- Usa topic con sufijo `-test`
- Ideal para desarrollo y pruebas
- Valor por defecto si no se especifica

#### 🏭 Production (Para producción)

```env
TUYA_ENVIRONMENT=production
```

- Usa topic de producción oficial
- Solo para entornos de producción

### 📊 Visualización de Cambios de Estado

Cuando un dispositivo cambia de estado, verás una salida estructurada como esta:

```
────────────────────────────────────────────────────────────────────────────────
🏠 CAMBIO DE ESTADO - 14:30:25 15/12/2023
🔧 Dispositivo: bf1234567890abcdef
🏷️ Producto: switch_product_key
📄 ID de Datos: data_12345

📊 Estados Actualizados (2):
   🔘 Switch 1: 🟢 ENCENDIDO ⏰ 14:30:25
   ┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
   💡 Switch Led: 🔴 APAGADO ⏰ 14:30:25
────────────────────────────────────────────────────────────────────────────────
```

### 🎨 Emojis Contextuales

El sistema usa emojis inteligentes según el tipo de estado:

- 🔘 Switches generales
- 💡 Luces LED
- 🌡️ Temperatura
- 💧 Humedad
- 🔋 Batería
- 📶 Señal
- 🚪 Puertas
- 🏃 Movimiento
- 🚨 Alarmas
- ⚡ Energía/Poder

### 📈 Monitoreo del Listener

Consulta el estado del listener en cualquier momento:

```bash
curl -X GET "http://localhost:8080/api/pulsar/status"
```

Respuesta ejemplo:

```json
{
  "success": true,
  "message": "Estado del listener de Pulsar (se inicia automáticamente con el servidor)",
  "status": {
    "status": "running",
    "is_running": true,
    "is_connected": true,
    "environment": "testing",
    "topic": "persistent://your_access_id/out/event-test",
    "pulsar_addr": "pulsar+ssl://mqe.tuyaeu.com:7285",
    "start_time": "2023-12-15T14:30:25Z",
    "uptime_seconds": 3600,
    "last_message": "2023-12-15T15:25:10Z",
    "seconds_since_last_message": 300
  }
}
```

### ⚠️ Solución de Problemas del Listener

#### Error "sign invalid" en Pulsar

- Verifica que `TUYA_ACCESS_ID` y `TUYA_ACCESS_KEY` sean correctos
- Asegúrate de usar la región correcta en `TUYA_PULSAR_ADDR`

#### No recibe mensajes

- Verifica que tus dispositivos estén enviando datos
- Confirma que el ambiente (`testing` vs `production`) sea el correcto
- Revisa que el topic tenga el formato correcto en los logs

#### Listener no se inicia

- Verifica que todas las variables de entorno estén configuradas
- Confirma que la URL de Pulsar use `pulsar+ssl://`
- El servidor continuará funcionando aunque el listener falle

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
│   ├── tuya_handler.go      # Lógica principal de endpoints REST
│   ├── tuya_utils.go        # Utilidades de firma HMAC-SHA256
│   └── pulsar_handler.go    # Listener de Pulsar para tiempo real
├── routes/
│   ├── tuya_routes.go       # Rutas de API Tuya
│   ├── pulsar_routes.go     # Rutas de monitoreo Pulsar
│   └── routes.go            # Configuración de rutas principal
├── models/
│   ├── token.go             # Modelos de autenticación
│   └── pulsar_types.go      # Modelos de mensajes Pulsar
├── docs/                    # Documentación Swagger generada
├── scripts/
│   └── generate-docs.sh     # Script para generar documentación
├── main.go                 # Punto de entrada con listener automático
├── .env                    # Configuración de credenciales
├── config.example.env      # Plantilla de configuración
└── README.md              # Esta documentación
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

### 🔐 Problemas de API REST

#### Error "sign invalid"

- Verifica que las credenciales `TUYA_ACCESS_ID` y `TUYA_ACCESS_KEY` sean correctas
- Asegúrate de usar la región correcta en `TUYA_API_URL`

#### Error "No permission"

- Habilita el data center en tu cuenta de Tuya Cloud Platform
- Verifica que tu proyecto tenga permisos para acceder a la API de dispositivos

#### Token requests múltiples

- ✅ **Resuelto**: El sistema maneja correctamente múltiples requests de token consecutivos

### 🎧 Problemas del Listener de Pulsar

#### Error "sign invalid" en Pulsar

- Verifica que `TUYA_ACCESS_ID` y `TUYA_ACCESS_KEY` sean correctos
- Asegúrate de usar la región correcta en `TUYA_PULSAR_ADDR`

#### No recibe mensajes

- Verifica que tus dispositivos estén enviando datos
- Confirma que el ambiente (`testing` vs `production`) sea el correcto
- Revisa que el topic tenga el formato correcto en los logs

#### Listener no se inicia

- Verifica que todas las variables de entorno estén configuradas
- Confirma que la URL de Pulsar use `pulsar+ssl://`
- El servidor continuará funcionando aunque el listener falle

#### Desconexiones frecuentes

- Verifica la estabilidad de tu conexión a internet
- Consulta los logs para identificar patrones de error
- El listener se reconecta automáticamente

## 🔗 Referencias

### 📚 Documentación Oficial de Tuya

- [SDK Oficial de Tuya para Go](https://github.com/tuya/tuya-cloud-sdk-go)
- [SDK de Pulsar para Go](https://github.com/tuya/tuya-pulsar-sdk-go)
- [Documentación de Firma de Tuya](https://developer.tuya.com/en/docs/iot/new-singnature?id=Kbw0q34cs2e5g)
- [Documentación de Tuya Cloud API](https://developer.tuya.com/en/docs/cloud/)
- [Documentación de Pulsar Messaging](https://developer.tuya.com/en/docs/iot/message-subscription?id=Kbb1qb9gn44x1)

### 🌐 Plataformas de Tuya

- [Portal de Desarrolladores de Tuya](https://iot.tuya.com/)
- [Tuya Cloud Platform](https://platform.tuya.com/)

### 🛠️ Herramientas de Desarrollo

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Swagger/OpenAPI](https://swagger.io/)
- [Apache Pulsar](https://pulsar.apache.org/)

---

## 🎯 Conclusión

Este proyecto ofrece una **solución completa y moderna** que combina:

✅ **API REST** tradicional siguiendo exactamente el patrón del SDK oficial de Tuya  
✅ **Listener en tiempo real** para cambios de estado de dispositivos  
✅ **Visualización intuitiva** con emojis y formato estructurado  
✅ **Configuración flexible** de ambientes (testing/production)  
✅ **Documentación completa** con Swagger UI  
✅ **Arquitectura robusta** con manejo de errores y reconexión automática

**Ideal para**: Aplicaciones IoT que necesitan tanto consultas bajo demanda como notificaciones en tiempo real de cambios de estado de dispositivos Tuya.

**Ventajas**: Mayor flexibilidad de integración que el SDK nativo, manteniendo toda la funcionalidad y agregando capacidades de monitoreo en tiempo real.
