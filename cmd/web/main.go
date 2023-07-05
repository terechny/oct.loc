package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	userController "oct.loc/controllers/user"
	models "oct.loc/models"
	product "oct.loc/services/product"
)

type Product struct {
	Name        string
	Description string
	Price       float32
}

func productIndex(w http.ResponseWriter, r *http.Request) {

	products := models.GetProducts()
	jsonResp, err := json.Marshal(products)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://oct.front")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func productStore(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(32 << 20) // 32 MB максимальный размер
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f64, _ := strconv.ParseFloat(r.FormValue("Price"), 32)
	price := float32(f64)

	Product := product.Product{}

	Product.SetName(r.FormValue("Name"))
	Product.SetDescription(r.FormValue("Description"))
	Product.SetPrice(price)

	id, _ := models.ProductStore(Product)

	resp := make(map[string]int64)
	resp["product"] = id
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func productShow(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	u64, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		panic(err)
	}

	u32 := uint32(u64)

	product := models.GetProduct(u32)

	jsonResp, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func productUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	u64, _ := strconv.ParseUint(params["id"], 10, 32)
	id := uint32(u64)

	f64, _ := strconv.ParseFloat(r.FormValue("Price"), 32)
	price := float32(f64)

	Product := product.Product{}

	Product.SetId(id)
	Product.SetName(r.FormValue("Name"))
	Product.SetDescription(r.FormValue("Description"))
	Product.SetPrice(price)

	models.ProductUpdate(Product)
}

func productDelete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	u64, _ := strconv.ParseUint(params["id"], 10, 32)
	id := uint32(u64)

	models.ProductDelete(id)
}

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/product", productIndex).Methods("GET")
	mux.HandleFunc("/product", productStore).Methods("POST")
	mux.HandleFunc("/product/{id}", productShow).Methods("GET")
	mux.HandleFunc("/product/{id}", productUpdate).Methods("PUT")
	mux.HandleFunc("/product/{id}", productDelete).Methods("DELETE")
	mux.HandleFunc("/user", userController.Store).Methods("POST")
	mux.HandleFunc("/user/{id}", userController.Show).Methods("GET")
	mux.HandleFunc("/user", userController.Index).Methods("GET")

	log.Println("Server Start on http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Panicln(err)
}
