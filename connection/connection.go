package connection

import (
	"database/sql"
	"fmt"

	"myproject/models"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/meli_db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")

	return db, nil
}

func InsertCustomer(db *sql.DB, customer *models.Customer) error {
	query := "INSERT INTO Customer (email, nombre, apellido, sexo, direccion, fecha_nac, telefono) VALUES (?, ?, ?, ?, ?, ?, ?)"

	// Preparar la sentencia SQL
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Ejecutar la sentencia SQL con los valores del cliente
	_, err = stmt.Exec(
		customer.Email,
		customer.Nombre,
		customer.Apellido,
		customer.Sexo,
		customer.Direccion,
		customer.Fecha_nac,
		customer.Telefono,
	)

	// _, err := db.Exec(query, customer.Email, customer.Nombre, customer.Apellido, customer.Sexo, customer.Direccion, customer.Fecha_nac, customer.Telefono)
	if err != nil {
		return err
	}

	fmt.Println("Successfully inserted customer into the database")

	return nil
}
