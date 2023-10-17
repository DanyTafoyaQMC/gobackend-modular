package router

import (
	"bufio"
	"fmt"
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
	var requestBodyLines []string
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
			requestBodyLines = append(requestBodyLines, ln)

			// Imprime cada línea del cuerpo para depuración
			fmt.Println("Línea del cuerpo:", ln)
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

	// Mueve la asignación de req.Body aquí
	req.Body = requestBodyLines

	if handlers, exists := r.routes[req.Route]; exists {
		if handler, exists := handlers[req.Method]; exists {
			handler(req, conn)
		} else {
			util.HttpMethodNotAllowed(conn)
		}
	} else {
		util.HttpNotFound(conn)
	}
}
