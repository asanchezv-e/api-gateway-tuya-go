package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"
)

// Variable global para mantener el token como en el SDK oficial
var Token string

// buildHeader construye los headers siguiendo exactamente el SDK oficial
func buildHeader(req *http.Request, body []byte, clientID, secret string) {
	req.Header.Set("client_id", clientID)
	req.Header.Set("sign_method", "HMAC-SHA256")

	ts := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	req.Header.Set("t", ts)

	// Detectar si es un token request
	isTokenRequest := strings.Contains(req.URL.Path, "/token")
	
	if !isTokenRequest && Token != "" {
		req.Header.Set("access_token", Token)
	}

	sign := buildSign(req, body, ts, clientID, secret, isTokenRequest)
	req.Header.Set("sign", sign)
}

// buildSign construye la firma siguiendo exactamente el algoritmo del SDK oficial
func buildSign(req *http.Request, body []byte, t, clientID, secret string, isTokenRequest bool) string {
	headers := getHeaderStr(req)
	urlStr := getUrlStr(req)
	contentSha256 := sha256Hash(body)
	stringToSign := req.Method + "\n" + contentSha256 + "\n" + headers + "\n" + urlStr
	
	// Para token requests: SIEMPRE clientID + t + stringToSign (sin token)
	// Para otros requests: clientID + Token + t + stringToSign
	var signStr string
	if isTokenRequest {
		// Token request - NUNCA incluir token en la firma
		signStr = clientID + t + stringToSign
	} else {
		// Request con token
		signStr = clientID + Token + t + stringToSign
	}
	
	sign := strings.ToUpper(hmacSha256(signStr, secret))
	return sign
}

// sha256Hash calcula el hash SHA256 de los datos
func sha256Hash(data []byte) string {
	sha256Contain := sha256.New()
	sha256Contain.Write(data)
	return hex.EncodeToString(sha256Contain.Sum(nil))
}

// getUrlStr construye la URL string siguiendo el SDK oficial
func getUrlStr(req *http.Request) string {
	url := req.URL.Path
	keys := make([]string, 0, 10)

	query := req.URL.Query()
	for key := range query {
		keys = append(keys, key)
	}
	if len(keys) > 0 {
		url += "?"
		sort.Strings(keys)
		for _, keyName := range keys {
			value := query.Get(keyName)
			url += keyName + "=" + value + "&"
		}
	}

	if len(url) > 0 && url[len(url)-1] == '&' {
		url = url[:len(url)-1]
	}
	return url
}

// getHeaderStr construye el string de headers para la firma
func getHeaderStr(req *http.Request) string {
	signHeaderKeys := req.Header.Get("Signature-Headers")
	if signHeaderKeys == "" {
		return ""
	}
	keys := strings.Split(signHeaderKeys, ":")
	headers := ""
	for _, key := range keys {
		headers += key + ":" + req.Header.Get(key) + "\n"
	}
	// Remover el último \n si existe
	if len(headers) > 0 && headers[len(headers)-1] == '\n' {
		headers = headers[:len(headers)-1]
	}
	return headers
}

// hmacSha256 calcula HMAC-SHA256
func hmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
} 