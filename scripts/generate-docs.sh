#!/bin/bash

# Script para regenerar la documentación de Swagger
# Uso: ./scripts/generate-docs.sh

echo "🔄 Regenerando documentación de Swagger..."

# Verificar que swag esté instalado
SWAG_PATH=$(go env GOPATH)/bin/swag
if [ ! -f "$SWAG_PATH" ]; then
    echo "❌ swag no está instalado. Instalando..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generar documentación
echo "📝 Generando documentación..."
$SWAG_PATH init

if [ $? -eq 0 ]; then
    echo "✅ Documentación generada exitosamente!"
    echo "🌐 Interfaz disponible en: http://localhost:8080/swagger/index.html"
    echo "📄 JSON disponible en: http://localhost:8080/swagger/doc.json"
else
    echo "❌ Error al generar la documentación"
    exit 1
fi

# Compilar aplicación
echo "🔨 Compilando aplicación..."
go build -o api-dispositivos main.go

if [ $? -eq 0 ]; then
    echo "✅ Aplicación compilada exitosamente!"
    echo "🚀 Para ejecutar: ./api-dispositivos"
else
    echo "❌ Error al compilar la aplicación"
    exit 1
fi

echo "🎉 Proceso completado!" 