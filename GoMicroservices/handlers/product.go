package handlers

import (
	"context"
	"fmt"
	"log"
	"microservices/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	pl := data.GetProducts()
	err := pl.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marhal json", http.StatusInternalServerError)
		return
	}
}

func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("Adding product %#v", prod)

	data.AddProduct(prod)
	fmt.Fprint(rw, "Successfully added the product")
}

func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("Updating product with id %v", id)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Unexpected error occurred", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(rw, "Successfully updated the product")
}

type KeyProduct struct{}

func (p *Product) MiddleWareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		if err = prod.Validate(); err != nil {
			http.Error(rw, fmt.Sprintf("Error validating data %v", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
