package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/bryancr89/groceries-api/models"
	"encoding/json"
	"fmt"
	"github.com/bryancr89/groceries-api/logic"
	"github.com/bryancr89/groceries-api/dao"
)

func main() {
	logic.Initialize(dao.ProductsDAO{})
	router := mux.NewRouter()
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, _ := logic.GetProducts()
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product, _ := logic.GetProduct(params["id"])
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	createdProduct, _ := logic.CreateProduct(product)
	json.NewEncoder(w).Encode(createdProduct)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deletedProduct, _ := logic.DeleteProduct(params["id"])
	json.NewEncoder(w).Encode(deletedProduct)
}