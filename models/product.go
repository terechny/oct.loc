package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	product "oct.loc/services"
)

var database *sql.DB

func ConnectDB() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_shop")

	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	database = db

	//fmt.Println("Successfully connected to database!")
}

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float32
}

func GetProduct(id uint32) Product {

	ConnectDB()

	product := Product{}

	row := database.QueryRow(`SELECT id, name, description, price FROM products WHERE id=?;`, id)

	err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price)

	if err != nil {
		fmt.Println(err)
	}

	return product
}

func GetProducts() []Product {

	ConnectDB()
	defer database.Close()

	rows, err := database.Query("SELECT id, name, description, price FROM `products`")

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	return products
}

func ProductStore(p product.Product) (int64, error) {

	ConnectDB()
	defer database.Close()

	result, err := database.Exec("INSERT INTO `products` (`name`, `description`, `price`) VALUES (?,?,?)", p.Name(), p.Description(), p.Price())

	if err != nil {
		return 0, err
	}

	lastInserId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastInserId, nil
}
