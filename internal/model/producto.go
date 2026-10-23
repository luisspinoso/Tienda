package model

type Producto struct {
	ID        int    `json:"id"`
	Nombre    string `json:"name"`
	Categoria string `json:"category"`
}
