package models

type Category struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Path   string `json:"path"`
}
