package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"myproject/connection"
	"myproject/models"
)

type CustomerController struct {
	DB *sql.DB
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

func (c *CustomerController) GetCustomers(w http.ResponseWriter, r *http.Request) {
	// Consultar los clientes en la base de datos
	rows, err := c.DB.Query("SELECT id, nombre, apellido, email FROM Customer")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Crear una lista para almacenar los clientes
	customers := []models.Customer{}

	// Iterar sobre los resultados de la consulta y crear objetos Customer
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.ID, &customer.Nombre, &customer.Apellido, &customer.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		customers = append(customers, customer)
	}

	// Verificar si ocurrió un error al iterar sobre los resultados
	if err = rows.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Devolver los clientes como respuesta en formato JSON
	json.NewEncoder(w).Encode(customers)
}
