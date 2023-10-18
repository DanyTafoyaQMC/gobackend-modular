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
	r.GET("/get", controllers.GetHandler)
	r.POST("/post", controllers.PostHandler)
	r.PUT("/put", controllers.PutHandler)
	r.PATCH("/patch", controllers.PatchHandler)
	r.DELETE("/delete", controllers.DeleteHandler)

	// Inicializar el servidor
	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		util.LogError("Error al iniciar el servidor:", err)
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
