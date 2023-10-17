// mierda que mi aplciacion puede necesitar
package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

// Función para leer una línea de texto desde una conexión y devolverla como string
func ReadLine(conn net.Conn) (string, error) {
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}

// Función para escribir una cadena en una conexión
func WriteLine(conn net.Conn, line string) error {
	_, err := conn.Write([]byte(line + "\n"))
	return err
}

// Funciones para manejar respuestas HTTP
func HttpNotFound(conn net.Conn) {
	response := "HTTP/1.1 404 Not Found\r\n" +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		"404 Not Found"
	conn.Write([]byte(response))
	conn.Close()
}

func HttpMethodNotAllowed(conn net.Conn) {
	response := "HTTP/1.1 405 Method Not Allowed\r\n" +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		"405 Method Not Allowed"
	conn.Write([]byte(response))
	conn.Close()
}

func HttpError(conn net.Conn, statusCode int, errorMessage string) {
	response := fmt.Sprintf("HTTP/1.1 %d Internal Server Error\r\n"+
		"Content-Type: application/json\r\n"+
		"\r\n"+
		`{"error":"%s"}`, statusCode, errorMessage)
	conn.Write([]byte(response))
	conn.Close()
}

func JsonResponse(conn net.Conn, statusCode int, data interface{}) {
	responseData, err := json.Marshal(data)
	if err != nil {
		HttpError(conn, http.StatusInternalServerError, err.Error())
		return
	}

	response := fmt.Sprintf("HTTP/1.1 %d OK\r\n"+
		"Content-Type: application/json\r\n"+
		"\r\n"+
		"%s", statusCode, responseData)

	conn.Write([]byte(response))
}

// LogError registra un mensaje de error en la consola
func LogError(message string, err error) {
	fmt.Printf("[ERROR] %s: %v\n", message, err)
}

func LogInfo(message string) {
	fmt.Printf("[INFO] %s\n", message)
}
