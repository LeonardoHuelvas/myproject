package models

type Item struct {
	ID           int    `json:"id"`
	Nombre       string `json:"nombre"`
	Estado       string `json:"estado"`
	Categoria_ID int    `json:"category_id"`
}
