package main

import (
	"docker/warehouse"
	"html/template"
	"io"
)

var productsTemplate = template.Must(template.ParseFiles("views/products.html"))
var newProduct = template.Must(template.ParseFiles("views/newProduct.html"))

func HtmlProducts(w io.Writer, products []warehouse.Product) {
	productsTemplate.Execute(w, products)
}

func NewProduct(w io.Writer) {
	newProduct.Execute(w, nil)
}
