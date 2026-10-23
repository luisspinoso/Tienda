package store

import (
	"Tienda/internal/model"
	"database/sql"
)

type Store interface {
	GetAll() ([]*model.Producto, error)
	GetByID(id int) (*model.Producto, error)
	Create(producto *model.Producto) (*model.Producto, error)
	Update(id int, producto *model.Producto) (*model.Producto, error)
	Delete(id int) error
}
type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

func (s *store) GetAll() ([]*model.Producto, error) {
	q := `SELECT id, name, category FROM productos`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []*model.Producto
	for rows.Next() {
		var b model.Producto
		if err := rows.Scan(&b.ID, &b.Nombre, &b.Categoria); err != nil {
			return nil, err
		}

		productos = append(productos, &b)
	}

	return productos, nil
}

func (s *store) GetByID(id int) (*model.Producto, error) {
	q := `SELECT id, name, category FROM productos WHERE id = ?`

	b := model.Producto{}
	err := s.db.QueryRow(q, id).Scan(&b.ID, &b.Nombre, &b.Categoria)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (s *store) Create(producto *model.Producto) (*model.Producto, error) {
	q := `INSERT INTO productos (name, category) VALUES (?, ?)`

	resp, err := s.db.Exec(q, producto.Nombre, producto.Categoria)
	if err != nil {
		return nil, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}

	producto.ID = int(id)

	return producto, nil
}
func (s *store) Update(id int, producto *model.Producto) (*model.Producto, error) {
	q := `UPDATE productos SET name = ?, category = ? WHERE id = ?`

	_, err := s.db.Exec(q, producto.Nombre, producto.Categoria, id)
	if err != nil {
		return nil, err
	}

	producto.ID = id

	return producto, nil
}

func (s *store) Delete(id int) error {
	q := `DELETE from productos WHERE id = ?`

	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}
