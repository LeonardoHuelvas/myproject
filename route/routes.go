package route

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"myproject/connection"
	"myproject/models"

	"github.com/gorilla/mux"
)

type Route struct {
	DB *sql.DB
}

type CustomerController struct {
	DB *sql.DB
}

func (route *Route) NewRouter() *mux.Router {
	// Crear un nuevo router
	r := mux.NewRouter()

	customerController := &CustomerController{
		DB: route.DB,
	}

	// Configurar las rutas y utilizar los métodos de controlador correspondientes
	r.HandleFunc("/customers", customerController.GetCustomers).Methods("GET")
	r.HandleFunc("/customers", customerController.CreateCustomer).Methods("POST")

	return r
}

func (c *CustomerController) GetCustomers(w http.ResponseWriter, r *http.Request) {
	// Aquí puedes hacer lo que necesites para obtener los datos del cliente
	// Por ahora, solo vamos a devolver un cliente de ejemplo
	customer := &models.Customer{
		ID:        0,
		Email:     "samm@example.com",
		Nombre:    "Javier",
		Apellido:  "Borja",
		Sexo:      "M",
		Direccion: "Address",
		Fecha_nac: "1990-01-01",
		Telefono:  "300011212",
	}
	json.NewEncoder(w).Encode(customer)
}

func (c *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insertar el cliente en la base de datos
	err = connection.InsertCustomer(c.DB, &customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}
