package main

import (
	"docker/warehouse"
	"net/http"
	"strconv"
)

type WarehouseController struct {
	warehouse *warehouse.Warehouse
}

func (w *WarehouseController) ListProducts(res http.ResponseWriter, req *http.Request) {
	HtmlProducts(res, w.warehouse.ListProducts())
}

func (w *WarehouseController) NewProduct(res http.ResponseWriter, req *http.Request) {
	NewProduct(res)
}

func (w *WarehouseController) SubmitProduct(res http.ResponseWriter, req *http.Request) {
	Qty, _ := strconv.ParseInt(req.PostFormValue("Qty"), 10, 64)
	Name := req.PostFormValue("Name")
	w.warehouse.AddProduct(warehouse.Product{1, Name, int(Qty)})
	HtmlProducts(res, w.warehouse.ListProducts())
}
