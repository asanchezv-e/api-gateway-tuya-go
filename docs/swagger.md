# Documentación Swagger - API de Dispositivos IoT

## Acceso a la Documentación

La documentación interactiva de Swagger está disponible en las siguientes URLs:

### 🌐 Interfaz Web Interactiva

```
http://localhost:8080/swagger/index.html
```

### 📄 Documentación JSON

```
http://localhost:8080/swagger/doc.json
```

### 📋 Documentación YAML

```
http://localhost:8080/swagger/swagger.yaml
```

## Características de la Documentación

### ✅ Endpoints Documentados

- **GET /api/devices** - Obtener todos los dispositivos
- **GET /api/devices/{id}** - Obtener dispositivo por ID
- **GET /api/devices/owner/{ownerId}** - Obtener dispositivos por propietario
- **GET /api/devices/online** - Obtener dispositivos en línea

### 📝 Información Incluida

- Descripción detallada de cada endpoint
- Parámetros requeridos y opcionales
- Modelos de datos con ejemplos
- Códigos de respuesta HTTP
- Estructura de respuestas de error

### 🎯 Funcionalidades Interactivas

- Probar endpoints directamente desde la interfaz
- Ver ejemplos de respuestas
- Validación automática de parámetros
- Descarga de documentación en formato OpenAPI

## Cómo Usar la Interfaz Swagger

1. **Abrir la interfaz**: Navega a `http://localhost:8080/swagger/index.html`
2. **Explorar endpoints**: Expande cualquier endpoint para ver detalles
3. **Probar endpoint**: Haz clic en "Try it out"
4. **Ingresar parámetros**: Completa los parámetros requeridos
5. **Ejecutar**: Haz clic en "Execute" para probar la API
6. **Ver respuesta**: Revisa la respuesta y códigos de estado

## Estructura de Respuesta Estándar

Todos los endpoints siguen el mismo formato de respuesta:

```json
{
  "success": true,
  "message": "Descripción del resultado",
  "data": {
    // Datos específicos del endpoint
  },
  "error": "Mensaje de error (si aplica)"
}
```

## Regenerar Documentación

Si modificas las anotaciones de Swagger en el código, regenera la documentación con:

```bash
$(go env GOPATH)/bin/swag init
```

## Ejemplos de Uso

### Obtener todos los dispositivos

```bash
curl http://localhost:8080/api/devices
```

### Obtener dispositivo específico

```bash
curl http://localhost:8080/api/devices/device-001
```

### Obtener dispositivos en línea

```bash
curl http://localhost:8080/api/devices/online
```

### Obtener dispositivos por propietario

```bash
curl http://localhost:8080/api/devices/owner/owner-001
```
