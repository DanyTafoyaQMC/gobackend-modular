package controllers

import (
	"fmt"
	"net"

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

func GetHandler(request *router.Request, conn net.Conn) {
	fmt.Println(request.Host, request.Route, request.Method)
	data := map[string]interface{}{
		"message": "Bienvenido aqui puedes ver la informacion de la ruta a la que añadiste.",
		"host":    request.Host,
		"route":   request.Route,
		"params":  request.Params,
	}
	util.JsonResponse(conn, 200, data)
	conn.Close()

}

// alguien debería acabar conmigo
func PostHandler(request *router.Request, conn net.Conn) {
	fmt.Println(request.Host, request.Route, request.Method)
	data := map[string]interface{}{
		"message": "Bienvenido aqui puedes ver la informacion de la ruta a la que añadiste.",
		"host":    request.Host,
		"route":   request.Route,
		"params":  request.Params,
		"body":    request.Body,
	}
	util.JsonResponse(conn, 200, data)
	conn.Close()

}

func PutHandler(request *router.Request, conn net.Conn) {
	fmt.Println(request.Host, request.Route, request.Method)
	data := map[string]interface{}{
		"message": "Bienvenido aqui puedes ver la informacion de la ruta a la que entraste.",
		"host":    request.Host,
		"route":   request.Route,
		"params":  request.Params,
	}
	util.JsonResponse(conn, 200, data)
	conn.Close()

}
func PatchHandler(request *router.Request, conn net.Conn) {
	fmt.Println(request.Host, request.Route, request.Method)
	data := map[string]interface{}{
		"message": "Bienvenido aqui puedes ver la informacion de la ruta a la que entraste.",
		"host":    request.Host,
		"route":   request.Route,
		"params":  request.Params,
	}
	util.JsonResponse(conn, 200, data)
	conn.Close()

}

func DeleteHandler(request *router.Request, conn net.Conn) {
	fmt.Println(request.Host, request.Route, request.Method)
	data := map[string]interface{}{
		"message": "Bienvenido aqui puedes ver la informacion de la ruta a la que añadiste.",
		"host":    request.Host,
		"route":   request.Route,
		"params":  request.Params,
	}
	util.JsonResponse(conn, 200, data)
	conn.Close()

}
