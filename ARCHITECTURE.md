# Arquitectura Limpia - API de Dispositivos IoT

## Descripción General

Esta API REST implementa una arquitectura limpia (Clean Architecture) siguiendo los principios SOLID para la gestión de dispositivos IoT. La arquitectura está diseñada para ser escalable, mantenible y testeable.

## Documentación con Swagger

La API incluye documentación automática con Swagger/OpenAPI que proporciona:

- **Interfaz interactiva**: `http://localhost:8080/swagger/index.html`
- **Documentación JSON**: `http://localhost:8080/swagger/doc.json`
- **Archivo YAML**: `http://localhost:8080/swagger/swagger.yaml`

### Características de Swagger

- Documentación automática basada en anotaciones en el código
- Interfaz web para probar endpoints directamente
- Validación automática de parámetros
- Ejemplos de respuestas y modelos de datos
- Descarga de especificación OpenAPI

## Principios SOLID Implementados

### 1. Single Responsibility Principle (SRP)

- **DeviceService**: Se encarga únicamente de la lógica de negocio de dispositivos
- **DeviceHandler**: Solo maneja las peticiones HTTP y respuestas
- **DeviceRepository**: Únicamente responsable del acceso a datos

### 2. Open/Closed Principle (OCP)

- Las rutas están separadas en módulos (`device_routes.go`) permitiendo extensibilidad
- Nuevos endpoints se pueden agregar sin modificar código existente

### 3. Liskov Substitution Principle (LSP)

- `MemoryDeviceRepository` implementa `DeviceRepository` interface
- Se puede sustituir por cualquier otra implementación (base de datos, cache, etc.)

### 4. Interface Segregation Principle (ISP)

- `DeviceRepository` tiene métodos específicos y cohesivos
- No fuerza a implementar métodos innecesarios

### 5. Dependency Inversion Principle (DIP)

- `DeviceService` depende de la abstracción `DeviceRepository`, no de implementaciones concretas
- Inyección de dependencias a través del `Container`

## Estructura de Carpetas

```
├── main.go                              # Punto de entrada
├── models/                              # Modelos de datos
│   └── device.go
├── domain/                              # Lógica de dominio
│   ├── interfaces/                      # Interfaces/contratos
│   │   └── device_repository.go
│   └── services/                        # Servicios de negocio
│       └── device_service.go
├── infrastructure/                      # Capa de infraestructura
│   └── repositories/                    # Implementaciones de repositorios
│       └── memory_device_repository.go
├── handlers/                            # Controladores HTTP
│   └── device_handler.go
├── routes/                              # Configuración de rutas
│   ├── routes.go
│   └── device_routes.go
└── config/                              # Configuración y DI
    └── container.go
```

## Flujo de Datos

1. **HTTP Request** → `DeviceHandler` (Capa de Presentación)
2. **Handler** → `DeviceService` (Capa de Aplicación)
3. **Service** → `DeviceRepository` (Interface)
4. **Repository** → `MemoryDeviceRepository` (Capa de Infraestructura)

## Endpoints Disponibles

### GET /api/devices

Obtiene todos los dispositivos registrados

### GET /api/devices/:id

Obtiene un dispositivo específico por su ID

### GET /api/devices/owner/:ownerId

Obtiene todos los dispositivos de un propietario específico

### GET /api/devices/online

Obtiene solo los dispositivos que están en línea

## Formato de Respuesta Estándar

```json
{
  "success": true,
  "message": "Descripción del resultado",
  "data": {...},
  "error": "Mensaje de error (si aplica)"
}
```

## Ventajas de esta Arquitectura

1. **Testabilidad**: Cada capa se puede testear independientemente
2. **Mantenibilidad**: Separación clara de responsabilidades
3. **Escalabilidad**: Fácil agregar nuevas funcionalidades
4. **Flexibilidad**: Cambiar implementaciones sin afectar otras capas
5. **Reutilización**: Servicios pueden ser reutilizados en diferentes contextos

## Extensibilidad

Para agregar nuevas funcionalidades:

1. **Nuevos endpoints**: Agregar en `device_routes.go` y métodos en `DeviceHandler`
2. **Nueva lógica de negocio**: Extender `DeviceService`
3. **Nuevas fuentes de datos**: Implementar `DeviceRepository` interface
4. **Nuevos modelos**: Agregar en el directorio `models/`

## Ejemplo de Uso

```bash
# Obtener todos los dispositivos
curl http://localhost:8080/api/devices

# Obtener dispositivo específico
curl http://localhost:8080/api/devices/device-001

# Obtener dispositivos de un propietario
curl http://localhost:8080/api/devices/owner/owner-001

# Obtener dispositivos en línea
curl http://localhost:8080/api/devices/online
```

## Regeneración de Documentación Swagger

Para regenerar la documentación cuando se modifiquen las anotaciones:

### Método Manual

```bash
$(go env GOPATH)/bin/swag init
```

### Usando el Script Automatizado

```bash
./scripts/generate-docs.sh
```

El script automatizado:

- Verifica e instala swag si no está disponible
- Regenera la documentación
- Compila la aplicación
- Proporciona URLs de acceso

## Notas de Desarrollo

- Las anotaciones de Swagger están en los handlers (`handlers/device_handler.go`)
- La configuración principal está en `main.go`
- Los archivos generados están en el directorio `docs/`
- La interfaz web está disponible en `/swagger/index.html`
