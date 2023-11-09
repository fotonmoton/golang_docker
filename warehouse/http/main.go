package main

import (
	"docker/warehouse"
	"docker/warehouse/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// inmemory := db.NewWarehouseInMemoryState()
	mysql := db.NewMysqlWarehouseState()

	// wh := warehouse.NewWarehouse(inmemory)
	wh := warehouse.NewWarehouse(mysql)
	whc := &WarehouseController{warehouse: wh}

	r := mux.NewRouter()
	r.HandleFunc("/warehouse/products/list", whc.ListProducts).Methods("GET")
	r.HandleFunc("/warehouse/products/new", whc.NewProduct).Methods("GET")
	r.HandleFunc("/warehouse/products", whc.SubmitProduct).Methods("POST")

	log.Println("Listening on :8080...")
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatal(err)
	}
}
