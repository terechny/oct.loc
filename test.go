package main

import (
	product "oct.loc/services"
)

func main() {

	Product := product.Product{}

	Product.SetName("Notebook Samsung A20")
	Product.SetDescription("Product Description")
	Product.SetPrice(48060.50)
}
