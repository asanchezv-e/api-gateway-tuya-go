# API de Dispositivos IoT - Arquitectura Limpia

Una API REST moderna para la gestión de dispositivos IoT implementada con **Go**, **Gin** y **Clean Architecture**.

## 🚀 Características

- ✅ **Arquitectura Limpia** con principios SOLID
- ✅ **Documentación automática** con Swagger/OpenAPI
- ✅ **API REST** bien estructurada
- ✅ **Inyección de dependencias**
- ✅ **Manejo de errores** estandarizado
- ✅ **Respuestas JSON** consistentes

## 📖 Documentación Interactiva

La API incluye documentación completa con Swagger UI:

### 🌐 Acceso Directo

```
http://localhost:8080/swagger/index.html
```

### 📋 Formatos Disponibles

- **Interfaz Web**: `http://localhost:8080/swagger/index.html`
- **JSON**: `http://localhost:8080/swagger/doc.json`
- **YAML**: `http://localhost:8080/swagger/swagger.yaml`

## 🛠️ Instalación y Ejecución

### Prerrequisitos

- Go 1.19+
- Git

### Pasos

1. **Clonar el repositorio**

```bash
git clone <repository-url>
cd go-gin-projet
```

2. **Instalar dependencias**

```bash
go mod tidy
```

3. **Generar documentación Swagger**

```bash
go install github.com/swaggo/swag/cmd/swag@latest
$(go env GOPATH)/bin/swag init
```

4. **Compilar y ejecutar**

```bash
go build -o api-dispositivos main.go
./api-dispositivos
```

### Script Automatizado

```bash
./scripts/generate-docs.sh
```

## 🔗 Endpoints Disponibles

| Método | Endpoint                       | Descripción                          |
| ------ | ------------------------------ | ------------------------------------ |
| GET    | `/api/devices`                 | Obtener todos los dispositivos       |
| GET    | `/api/devices/{id}`            | Obtener dispositivo por ID           |
| GET    | `/api/devices/owner/{ownerId}` | Obtener dispositivos por propietario |
| GET    | `/api/devices/online`          | Obtener dispositivos en línea        |

## 🏗️ Arquitectura

```
├── main.go                    # Punto de entrada con configuración Swagger
├── models/                    # Modelos de datos
├── domain/
│   ├── interfaces/           # Contratos/interfaces
│   └── services/            # Lógica de negocio
├── infrastructure/
│   └── repositories/        # Acceso a datos
├── handlers/                # Controladores HTTP con anotaciones Swagger
├── routes/                  # Configuración de rutas
├── config/                  # Inyección de dependencias
├── docs/                    # Documentación generada por Swagger
└── scripts/                 # Scripts de automatización
```

## 📱 Ejemplos de Uso

### Usando cURL

```bash
# Obtener todos los dispositivos
curl http://localhost:8080/api/devices

# Obtener dispositivo específico
curl http://localhost:8080/api/devices/device-001

# Obtener dispositivos en línea
curl http://localhost:8080/api/devices/online
```

### Usando la Interfaz Swagger

1. Navega a `http://localhost:8080/swagger/index.html`
2. Expande cualquier endpoint
3. Haz clic en "Try it out"
4. Completa los parámetros requeridos
5. Haz clic en "Execute"

## 🔄 Desarrollo

### Regenerar Documentación

Después de modificar anotaciones de Swagger:

```bash
$(go env GOPATH)/bin/swag init
```

### Estructura de Respuesta

Todas las respuestas siguen el formato estándar:

```json
{
  "success": true,
  "message": "Descripción del resultado",
  "data": { ... },
  "error": "Mensaje de error (si aplica)"
}
```

## 📚 Documentación Adicional

- [Arquitectura Detallada](ARCHITECTURE.md)
- [Documentación Swagger](docs/swagger.md)

## 🤝 Contribuir

1. Fork del proyecto
2. Crear rama de feature (`git checkout -b feature/AmazingFeature`)
3. Commit de cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abrir Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para detalles.
