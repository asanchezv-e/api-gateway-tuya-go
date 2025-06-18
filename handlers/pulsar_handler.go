package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"api-gateway-tuya-go/models"

	"github.com/joho/godotenv"
	pulsar "github.com/tuya/tuya-pulsar-sdk-go"
	"github.com/tuya/tuya-pulsar-sdk-go/pkg/tylog"
	"github.com/tuya/tuya-pulsar-sdk-go/pkg/tyutils"
)

// PulsarHandler maneja los mensajes de Pulsar
type PulsarHandler struct {
	config       models.PulsarConfig
	client       pulsar.Client
	consumer     pulsar.Consumer
	ctx          context.Context
	cancel       context.CancelFunc
	isRunning    bool
	isConnected  bool
	startTime    time.Time
	lastMessage  time.Time
	mu           sync.RWMutex
	stopComplete chan bool
}

// NewPulsarHandler crea una nueva instancia del handler de Pulsar
func NewPulsarHandler() *PulsarHandler {
	ctx, cancel := context.WithCancel(context.Background())
	return &PulsarHandler{
		ctx:          ctx,
		cancel:       cancel,
		stopComplete: make(chan bool, 1),
	}
}

// LoadConfig carga la configuración desde las variables de entorno
func (p *PulsarHandler) LoadConfig() error {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(".env"); err != nil {
		// Si no existe .env, intenta cargar config.example.env
		_ = godotenv.Load("config.example.env")
	}

	accessID := os.Getenv("TUYA_ACCESS_ID")
	accessKey := os.Getenv("TUYA_ACCESS_KEY")
	pulsarAddr := os.Getenv("TUYA_PULSAR_ADDR")

	if accessID == "" || accessKey == "" || pulsarAddr == "" {
		return fmt.Errorf("configuración de Pulsar incompleta: TUYA_ACCESS_ID, TUYA_ACCESS_KEY y TUYA_PULSAR_ADDR son requeridos")
	}

	// Verificar que las credenciales no sean valores de ejemplo
	if accessID == "your_access_id_here" || accessKey == "your_access_key_here" {
		return fmt.Errorf("credenciales de ejemplo detectadas. Por favor, configura credenciales reales de Tuya en .env")
	}

	// Validar dirección Pulsar según documentación oficial
	// Ejemplo: "pulsar+ssl://mqe.tuyaus.com:7285" para US
	// La documentación indica usar pulsar+ssl:// para conexiones seguras
	if !strings.HasPrefix(pulsarAddr, "pulsar+ssl://") {
		if strings.HasPrefix(pulsarAddr, "pulsar://") {
			pulsarAddr = strings.Replace(pulsarAddr, "pulsar://", "pulsar+ssl://", 1)
			fmt.Printf("⚠️ Cambiando a SSL: %s\n", pulsarAddr)
		} else {
			return fmt.Errorf("URL de Pulsar inválida. Debe usar formato: pulsar+ssl://mqe.region.com:7285")
		}
	}

	p.config = models.PulsarConfig{
		AccessID:   accessID,
		AccessKey:  accessKey,
		PulsarAddr: pulsarAddr,
		Topic:      pulsar.TopicForAccessID(accessID),
	}

	return nil
}

// StartListener inicia el listener de Pulsar siguiendo el patrón oficial
func (p *PulsarHandler) StartListener() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	if p.isRunning {
		return fmt.Errorf("el listener ya está en ejecución")
	}

	if err := p.LoadConfig(); err != nil {
		return fmt.Errorf("error cargando configuración: %v", err)
	}

	// Configurar logs según documentación oficial
	tylog.SetGlobalLog("sdk", false) // false = mostrar logs para debug
	
	fmt.Println("🎧 Listener de dispositivos Tuya iniciado")
	fmt.Printf("📡 Conectando a: %s\n", p.config.PulsarAddr)
	fmt.Printf("📝 Topic: %s\n", p.config.Topic)

	// Crear cliente según el ejemplo oficial
	cfg := pulsar.ClientConfig{
		PulsarAddr: p.config.PulsarAddr,
	}
	client := pulsar.NewClient(cfg)
	p.client = client

	// Crear consumer según el patrón oficial
	csmCfg := pulsar.ConsumerConfig{
		Topic: p.config.Topic,
		Auth:  pulsar.NewAuthProvider(p.config.AccessID, p.config.AccessKey),
	}
	
	consumer, err := client.NewConsumer(csmCfg)
	if err != nil {
		p.isConnected = false
		return fmt.Errorf("error creando consumer: %v", err)
	}
	p.consumer = consumer
	p.isConnected = true
	p.isRunning = true
	p.startTime = time.Now()
	
	// Solo crear nuevo canal si no existe o está cerrado
	if p.stopComplete == nil {
		p.stopComplete = make(chan bool, 1)
	}

	fmt.Println("✅ Consumer creado exitosamente")
	fmt.Println("📡 Esperando cambios de estado...")

	// Iniciar el consumidor en goroutine
	go p.startConsumer()

	return nil
}

// startConsumer inicia el consumidor de mensajes según el patrón oficial
func (p *PulsarHandler) startConsumer() {
	defer func() {
		fmt.Println("🔄 startConsumer finalizando...")
		
		p.mu.Lock()
		wasRunning := p.isRunning
		p.isRunning = false
		p.isConnected = false
		p.mu.Unlock()
		
		if r := recover(); r != nil {
			fmt.Printf("❌ Panic en listener: %v\n", r)
		}
		
		// Solo señalar parada si estaba corriendo (evitar señales duplicadas)
		if wasRunning {
			select {
			case p.stopComplete <- true:
				fmt.Println("✅ Señal de parada enviada")
			default:
				fmt.Println("⚠️ Canal de parada lleno o cerrado")
			}
		}
		
		fmt.Println("🏁 startConsumer terminado")
	}()

	fmt.Println("🎧 Iniciando recepción de mensajes...")
	
	// Crear handler según el ejemplo oficial con AesSecret = accessKey[8:24]
	handler := &DeviceMessageHandler{
		AesSecret: p.config.AccessKey[8:24], // Según documentación oficial
		AccessKey: p.config.AccessKey,
		DebugMode: p.config.DebugMode,
		parent:    p, // Referencia al handler padre para actualizar lastMessage
	}
	
	// Usar ReceiveAndHandle como en el ejemplo oficial
	// Este método debería respetar la cancelación del contexto
	fmt.Println("📡 Iniciando ReceiveAndHandle...")
	p.consumer.ReceiveAndHandle(p.ctx, handler)
	fmt.Println("📡 ReceiveAndHandle terminado")
}

// DeviceMessageHandler implementa el handler de mensajes según el patrón oficial
type DeviceMessageHandler struct {
	AesSecret string
	AccessKey string
	DebugMode bool
	parent    *PulsarHandler
}

// HandlePayload implementa la interfaz requerida por el SDK oficial
func (h *DeviceMessageHandler) HandlePayload(ctx context.Context, msg pulsar.Message, payload []byte) error {
	// Actualizar timestamp del último mensaje
	if h.parent != nil {
		h.parent.mu.Lock()
		h.parent.lastMessage = time.Now()
		h.parent.mu.Unlock()
	}

	// Log de payload recibido según documentación oficial
	tylog.Info("payload preview", tylog.String("payload", string(payload)))
	fmt.Printf("🔍 [%s] Mensaje recibido\n", time.Now().Format("15:04:05"))

	// Decodificar payload según el ejemplo oficial
	m := map[string]interface{}{}
	err := json.Unmarshal(payload, &m)
	if err != nil {
		tylog.Error("json unmarshal failed", tylog.ErrorField(err))
		return nil
	}

	// Extraer data según patrón oficial
	bs, ok := m["data"].(string)
	if !ok {
		fmt.Printf("⚠️ No se encontró campo 'data' en payload\n")
		return nil
	}

	// Decodificar base64 según ejemplo oficial
	de, err := base64.StdEncoding.DecodeString(bs)
	if err != nil {
		tylog.Error("base64 decode failed", tylog.ErrorField(err))
		return nil
	}

	// Usar AES secret según documentación: accessKey[8:24]
	decode := tyutils.EcbDecrypt(de, []byte(h.AesSecret))
	tylog.Info("aes decode", tylog.ByteString("decode payload", decode))

	// Mostrar datos decodificados
	fmt.Printf("📄 Datos decodificados: %s\n", string(decode))

	// Intentar decodificar como cambio de estado de dispositivo
	var deviceChange models.DeviceStatusChange
	if err := json.Unmarshal(decode, &deviceChange); err != nil {
		fmt.Printf("📊 Mensaje genérico (no es cambio de dispositivo)\n")
		return nil
	}

	// Procesar cambio de estado
	h.processDeviceStatusChange(deviceChange)

	return nil
}

// processDeviceStatusChange procesa un cambio de estado de dispositivo
func (h *DeviceMessageHandler) processDeviceStatusChange(change models.DeviceStatusChange) {
	timestamp := time.Now()
	
	fmt.Printf("\n🔔 [%s] Dispositivo: %s\n", 
		timestamp.Format("15:04:05"), 
		change.DevID)
	
	for _, status := range change.Status {
		fmt.Printf("   📊 %s: %v\n", status.Code, status.Value)
	}
}

// Stop detiene el listener de Pulsar
func (p *PulsarHandler) Stop() {
	p.mu.Lock()
	if !p.isRunning {
		p.mu.Unlock()
		fmt.Println("⚠️ El listener ya está detenido")
		return
	}
	
	fmt.Println("🛑 Deteniendo listener...")
	fmt.Printf("🔍 Estado actual - Running: %v, Connected: %v\n", p.isRunning, p.isConnected)
	
	// Marcar como no ejecutándose para evitar nuevas operaciones
	p.isRunning = false
	p.mu.Unlock()
	
	// Cancelar contexto para detener el consumer
	if p.cancel != nil {
		fmt.Println("🚪 Cancelando contexto...")
		p.cancel()
	}
	
	// Intentar cerrar el consumer primero para forzar la desconexión
	if p.consumer != nil {
		fmt.Println("🔌 Cerrando consumer...")
		
		// Cerrar el consumer en una goroutine separada para evitar bloqueos
		consumerClosed := make(chan error, 1)
		go func() {
			consumerClosed <- p.consumer.Close()
		}()
		
		// Esperar cierre del consumer con timeout
		select {
		case err := <-consumerClosed:
			if err != nil {
				fmt.Printf("⚠️ Error cerrando consumer: %v\n", err)
			} else {
				fmt.Println("✅ Consumer cerrado exitosamente")
			}
		case <-time.After(5 * time.Second):
			fmt.Println("⚠️ Timeout cerrando consumer, forzando desconexión")
		}
		
		p.consumer = nil
	}
	
	// Esperar a que startConsumer termine con timeout reducido
	if p.stopComplete != nil {
		fmt.Println("⏱️ Esperando confirmación de parada...")
		select {
		case <-p.stopComplete:
			fmt.Println("✅ startConsumer terminó correctamente")
		case <-time.After(3 * time.Second):
			fmt.Println("⚠️ Timeout esperando que startConsumer termine")
		}
		
		// Drenar el canal si tiene datos pendientes
		select {
		case <-p.stopComplete:
		default:
		}
	}
	
	// El cliente Pulsar no requiere cierre explícito según la documentación
	if p.client != nil {
		fmt.Println("✅ Cliente Pulsar liberado")
		p.client = nil
	}
	
	// Actualizar estado final
	p.mu.Lock()
	p.isRunning = false
	p.isConnected = false
	p.mu.Unlock()
	
	fmt.Println("🔄 Listener completamente desconectado")
}

// GetStatus devuelve el estado del listener
func (p *PulsarHandler) GetStatus() map[string]interface{} {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	status := "stopped"
	if p.isRunning && p.isConnected {
		status = "running"
	} else if p.isRunning && !p.isConnected {
		status = "connecting"
	} else if !p.isRunning && p.isConnected {
		status = "stopping"
	}
	
	response := map[string]interface{}{
		"status":       status,
		"is_running":   p.isRunning,
		"is_connected": p.isConnected,
	}
	
	// Agregar información adicional solo si está configurado
	if p.config.Topic != "" {
		response["topic"] = p.config.Topic
		response["pulsar_addr"] = p.config.PulsarAddr
	}
	
	// Agregar tiempos si está corriendo
	if p.isRunning {
		response["start_time"] = p.startTime.Format(time.RFC3339)
		response["uptime_seconds"] = int(time.Since(p.startTime).Seconds())
		
		if !p.lastMessage.IsZero() {
			response["last_message"] = p.lastMessage.Format(time.RFC3339)
			response["seconds_since_last_message"] = int(time.Since(p.lastMessage).Seconds())
		} else {
			response["last_message"] = nil
			response["seconds_since_last_message"] = nil
		}
	}
	
	return response
} 