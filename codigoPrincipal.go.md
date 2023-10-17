package main

import (
	"net"

	controllers ".git/ErnestoDanielTafoyaMolina/controllers"
	router ".git/ErnestoDanielTafoyaMolina/router"
	util ".git/ErnestoDanielTafoyaMolina/utils"
)

func main() {

	// Configurar el enrutador
	r := router.NewRouter()

	//rutas
	r.GET("/", controllers.RootHandler)
	r.GET("/pija", controllers.PijaJaHandler)
	r.GET("/gay", controllers.ElProfeEsGay)
	r.POST("/login", controllers.LoginHandler)

	// Inicializar el servidor
	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		util.LogError("Errowr al iniciar el servidor:", err)
		return
	}
	defer listener.Close()

	util.LogInfo("Servidor escuchando en :3001")

	for {
		conn, err := listener.Accept()
		if err != nil {
			util.LogError("Error al aceptar la conexión:", err)
			continue
		}

		go r.ServeHTTP(conn)
	}
}
package router

import (
	"bufio"
	"net"
	"strings"

	util ".git/ErnestoDanielTafoyaMolina/utils"
)

// Estructura de Request
type Request struct {
	Method  string
	Route   string
	Params  map[string]string
	Headers map[string]string
	Host    string
	Body    []string
}

// Definición del tipo RouteHandler
type RouteHandler func(request *Request, conn net.Conn)

// Definición del tipo Router
type Router struct {
	routes map[string]map[string]RouteHandler
}

// Función NewRouter para crear una nueva instancia de Router
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]RouteHandler),
	}
}

// Función para añadir rutas de manera dinámica al enrutador
func (r *Router) AddRoute(method string, route string, handler RouteHandler) {
	if r.routes[route] == nil {
		r.routes[route] = make(map[string]RouteHandler)
	}
	r.routes[route][method] = handler
}

// Métodos para definir rutas HTTP (GET, POST, PUT, PATCH, DELETE)
func (r *Router) GET(route string, handler RouteHandler) {
	r.AddRoute("GET", route, handler)
}

func (r *Router) POST(route string, handler RouteHandler) {
	r.AddRoute("POST", route, handler)
}

func (r *Router) PUT(route string, handler RouteHandler) {
	r.AddRoute("PUT", route, handler)
}

func (r *Router) PATCH(route string, handler RouteHandler) {
	r.AddRoute("PATCH", route, handler)
}

func (r *Router) DELETE(route string, handler RouteHandler) {
	r.AddRoute("DELETE", route, handler)
}

// Función ServeHTTP para manejar las solicitudes HTTP
func (r *Router) ServeHTTP(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	var header []string
	var body []string
	var f = true

	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			f = false
		}
		if f == true {
			header = append(header, ln)
		}
		if f == false {
			body = append(body, ln)
		}
		if f == false && ln == "" {
			break
		}
	}

	req := &Request{
		Params:  make(map[string]string),
		Headers: make(map[string]string),
	}
	for i, h := range header {
		if i == 0 {
			spl := strings.Split(h, " ")
			req.Method = spl[0]
			req.Route = spl[1]
		} else {
			spl := strings.Split(h, ": ")
			if spl[0] == "Host" {
				req.Host = spl[1]
			} else {
				req.Headers[spl[0]] = spl[1]
			}
		}
	}

	if handlers, exists := r.routes[req.Route]; exists {
		if handler, exists := handlers[req.Method]; exists {
			req.Body = body // Asigna el cuerpo de la solicitud a req.Body
			handler(req, conn)
		} else {
			util.HttpMethodNotAllowed(conn)
		}
	} else {
		util.HttpNotFound(conn)
	}
}
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
package controllers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	router ".git/ErnestoDanielTafoyaMolina/router"
	util ".git/ErnestoDanielTafoyaMolina/utils"
)

// Función controladora para la ruta raíz "/"
func RootHandler(request *router.Request, conn net.Conn) {
	fmt.Println(request.Host, request.Route, request.Method)
	data := map[string]string{"message": "Hola mundo, quiero queso."}
	util.JsonResponse(conn, 200, data)
	conn.Close()
}

func PijaJaHandler(request *router.Request, conn net.Conn) {
	fmt.Println(request.Host, request.Route, request.Method)
	data := map[string]string{
		"ok":       "true",
		"message":  "has accedido a la pija correctamente",
		"Provecho": "true",
	}
	util.JsonResponse(conn, 200, data)
	conn.Close()
}

type LoginRequest struct {
	Username string `json:"username"`
	Pass     string `json:"pass"`
}

// Función controladora para la ruta "/login"
func LoginHandler(request *router.Request, conn net.Conn) {

	//Imprime el encabezado y el cuerpo para depuración
	fmt.Println("Encabezados:")
	for i, h := range request.Headers {
		fmt.Printf("Header[%d]: %s\n", i, h)
	}
	fmt.Println("Host:", request.Host)
	fmt.Println("Ruta:", request.Route)
	fmt.Println("Método:", request.Method)
	go func() {
		fmt.Println("Cuerpo:")
		for i, b := range request.Body {
			fmt.Printf("Body[%d]: %s\n", i, b)
		}

		// aqui obtengo los datos del cuerpo de la solicitud
		body := strings.Join(request.Body, "\n")

		// Decodificar el JSON en una estructura LoginRequest
		var loginRequest LoginRequest
		if err := json.Unmarshal([]byte(body), &loginRequest); err != nil {
			data := map[string]string{"error": "Solicitud JSON no válida"}
			util.JsonResponse(conn, 400, data)
			conn.Close()
			return
		}

		// Verifica si el usuario y la contraseña son correctos
		user := "root"
		psw := "pass"

		if user != loginRequest.Username || psw != loginRequest.Pass {
			data := map[string]string{"error": "Nombre de usuario o contraseña incorrectos"}
			util.JsonResponse(conn, 404, data)
			conn.Close()
			return
		}

		// Si todo está correcto, responde con un mensaje de bienvenida
		data := map[string]string{"message": "Bienvenido don cangrejo"}
		util.JsonResponse(conn, http.StatusOK, data)
		conn.Close()
	}()
}

func ElProfeEsGay(request *router.Request, conn net.Conn) {
	fmt.Println(request.Host, request.Route, request.Method)
	body := strings.Join(request.Body, "\n")
	fmt.Println(body)
	data := map[string]string{"ok": "true", "msg": "El profe es gay y le da al ricardo"}
	util.JsonResponse(conn, 200, data)
	conn.Close()
}
