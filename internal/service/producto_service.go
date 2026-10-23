package service

import (
	"Tienda/internal/model"
	"Tienda/internal/store"
	"errors"
)

type Service struct {
	store store.Store
}

func New(s store.Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) ObtenTodosLosProductos() ([]*model.Producto, error) {

	productos, err := s.store.GetAll()
	if err != nil {

		return nil, err
	}
	return productos, nil
}

func (s *Service) ObtenProductoPorID(id int) (*model.Producto, error) {
	return s.store.GetByID(id)
}

func (s *Service) CrearProducto(producto model.Producto) (*model.Producto, error) {
	if producto.Nombre == "" {
		return nil, errors.New("El nombre del producto no puede estar vacío")
	}
	return s.store.Create(&producto)
}

func (s *Service) UpdateAllProducto(id int, producto model.Producto) (*model.Producto, error) {
	if producto.Nombre == "" {
		return nil, errors.New("El nombre del producto no puede estar vacío")
	}
	return s.store.Update(id, &producto)
}

func (s *Service) RemoverProducto(id int) error {
	return s.store.Delete(id)
}
