package controllers

import (
	"fmt"
	"net"
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
		fmt.Println("Cuerpo:", request.Body)
		for i, b := range request.Body {
			fmt.Printf("Body[%d]: %s\n", i, b)
		}

		/*
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



		*/
	}()
	// Si todo está correcto, responde con un mensaje de bienvenida
	// Responde con el cuerpo de la solicitud
	data := map[string]interface{}{
		"message":     "Bienvenido don cangrejo",
		"requestBody": request.Body,
	}
	util.JsonResponse(conn, 200, data)
	conn.Close()
}

func RicardoEsGay(request *router.Request, conn net.Conn) {
	fmt.Println(request.Host, request.Route, request.Method)
	body := strings.Join(request.Body, "\n")
	fmt.Println(body)
	data := map[string]string{"ok": "true", "msg": "Ricardo se come la come"}
	//util.WriteLine(conn, "Nomame, quien es el pendejo que entró aqui?")
	util.JsonResponse(conn, 200, data)
	conn.Close()
}
