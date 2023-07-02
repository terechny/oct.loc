package models

import (
	"fmt"

	user "oct.loc/services/user"
)

type User struct {
	id         uint32
	Firstname  string
	Secondname string
	Email      string
	Phone      string
	Password   string
}

func UserStore(user user.User) (int64, error) {

	fmt.Println(user)

	ConnectDB()
	defer database.Close()

	result, err := database.Exec("INSERT INTO `users` (`firstname`, `secondname`, `email`, `phone`, `password`) VALUES (?, ?, ?, ?, ?);", user.Firstname(), user.Secondname(), user.Email(), user.Phone(), user.Password())

	if err != nil {
		return 0, err
	}

	lastInserId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastInserId, nil
}

func UserGet(id uint32) User {

	ConnectDB()

	user := User{}

	row := database.QueryRow(`SELECT id, firstname, secondname, email, phone, password FROM users WHERE id = ?`, id)

	err := row.Scan(&user.id, &user.Firstname, &user.Secondname, &user.Email, &user.Phone, &user.Password)

	if err != nil {
		fmt.Println(err)
	}

	return user
}
