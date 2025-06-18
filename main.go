// Package main contiene la configuración principal del API Gateway Tuya Go
// @title API Gateway Tuya Go - SDK Oficial con Listener Pulsar
// @version 1.0
// @description API Gateway para Tuya Cloud con listener de mensajes Pulsar para cambios de estado de dispositivos
// @termsOfService http://swagger.io/terms/
// @contact.name Soporte API
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "api-gateway-tuya-go/docs" // docs generado por swag
	"api-gateway-tuya-go/handlers"
	"api-gateway-tuya-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("🚀 Iniciando API Gateway Tuya Go con Listener Pulsar...")
	log.Println("📡 Endpoints disponibles:")
	log.Println("   - API Tuya: /api/tuya/*")
	log.Println("   - Pulsar Status: GET /api/pulsar/status")
	log.Println("   - Swagger: /swagger/index.html")
	
	// Crear handler de Pulsar
	pulsarHandler := handlers.NewPulsarHandler()
	
	// Iniciar listener de Pulsar automáticamente
	log.Println("🎧 Iniciando listener de Pulsar automáticamente...")
	if err := pulsarHandler.StartListener(); err != nil {
		log.Printf("⚠️ Error iniciando listener de Pulsar: %v", err)
		log.Println("📋 Verifica la configuración en .env o config.example.env")
		log.Println("🔄 El servidor continuará sin el listener de Pulsar")
	} else {
		log.Println("✅ Listener de Pulsar iniciado exitosamente")
	}
	
	// Configurar Gin
	r := gin.Default()
	
	// Configurar rutas pasando el handler de Pulsar
	routes.SetupRoutes(r, pulsarHandler)
	
	// Configurar servidor HTTP
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	
	// Iniciar servidor en goroutine
	go func() {
		log.Println("🌐 Servidor HTTP iniciado en puerto 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Error iniciando servidor: %v", err)
		}
	}()
	
	// Configurar captura de señales para cierre elegante
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	log.Println("🛑 Cerrando servidor...")
	
	// Detener listener de Pulsar primero
	log.Println("🔌 Deteniendo listener de Pulsar...")
	pulsarHandler.Stop()
	
	// Cerrar servidor HTTP con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("⚠️ Error cerrando servidor: %v", err)
	} else {
		log.Println("✅ Servidor cerrado exitosamente")
	}
	
	log.Println("�� ¡Hasta luego!")
}