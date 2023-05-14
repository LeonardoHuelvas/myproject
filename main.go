package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"myproject/connection"
	"myproject/controllers"

	"github.com/gorilla/mux"
)

func main() {
	// Crear una nueva conexión a la base de datos
	db, err := connection.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	// Crear una instancia del controlador
	customerController := &controllers.CustomerController{
		DB: db,
	}

	// Crear un enrutador utilizando "gorilla/mux"
	router := mux.NewRouter()

	// Configurar las rutas y utilizar los métodos de controlador correspondientes
	router.HandleFunc("/customers", customerController.GetCustomers).Methods("GET")

	// Iniciar el servidor en un goroutine
	go func() {
		log.Fatal(http.ListenAndServe(":8000", router))
	}()

	log.Println("Servidor iniciado en http://localhost:8000")

	// Esperar señales para terminar el servidor
	waitForShutdown(db)
}

func waitForShutdown(db *sql.DB) {
	// Hacer un canal para recibir la señal de terminación
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Esperar la señal
	<-shutdown

	// Cerrar la conexión a la base de datos
	db.Close()

	log.Println("Servidor detenido")
}
