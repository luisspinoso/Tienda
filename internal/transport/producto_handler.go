package transport

import (
	"Tienda/internal/model"
	"Tienda/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ProductoHandler struct {
	service *service.Service
}

func New(s *service.Service) *ProductoHandler {
	return &ProductoHandler{
		service: s,
	}
}

func (h *ProductoHandler) HandleProductos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productos, err := h.service.ObtenTodosLosProductos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productos)

	case http.MethodPost:
		var producto model.Producto
		if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		created, err := h.service.CrearProducto(producto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)

	default:
		http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)
	}
}

func (h *ProductoHandler) HandleProductoByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/productos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "No lo encontre", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		producto, err := h.service.ObtenProductoPorID(id)
		if err != nil {
			http.Error(w, "No lo encontre", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(producto)

	case http.MethodPut:
		var producto model.Producto
		if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
			http.Error(w, "Input invalido", http.StatusBadRequest)
			return
		}

		update, err := h.service.UpdateAllProducto(id, producto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(update)

	case http.MethodDelete:
		if err := h.service.RemoverProducto(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)
	}
}
