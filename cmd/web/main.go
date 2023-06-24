package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	models "oct.loc/models"
	product "oct.loc/services"
)

func home(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "status Sreated"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

type Product struct {
	Name        string
	Description string
	Price       float32
}

func productIndex(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	products := models.GetProducts()
	jsonResp, err := json.Marshal(products)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
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

func main() {

	mux := mux.NewRouter()
	mux.HandleFunc("/", home).Methods("GET")
	mux.HandleFunc("/product", productIndex).Methods("GET")
	mux.HandleFunc("/product/{id}", productShow).Methods("GET")
	mux.HandleFunc("/product", productStore).Methods("POST")

	log.Println("Server Start on http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Panicln(err)
}
