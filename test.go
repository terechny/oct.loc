package main

import (
	"fmt"

	"oct.loc/models"
	product "oct.loc/services"
)

func main() {

	Product := product.Product{}

	Product.SetName("Notebook Samsung A20")
	Product.SetDescription("Product Description")
	Product.SetPrice(48060.50)

	//id, _ := models.ProductStore(Product)

	p := models.GetProduct(10)

	fmt.Println(p)
}
