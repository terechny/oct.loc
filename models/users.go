package models

import (
	"fmt"

	user "oct.loc/services/user"
)

type User struct {
	Id         uint32
	Firstname  string
	Secondname string
	Email      string
	Phone      string
	Password   string
}

func UserStore(user user.User) (int64, error) {

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

	err := row.Scan(&user.Id, &user.Firstname, &user.Secondname, &user.Email, &user.Phone, &user.Password)

	if err != nil {
		fmt.Println(err)
	}

	return user
}

func GetUsers() []User {

	ConnectDB()
	defer database.Close()

	rows, err := database.Query("SELECT id, firstname, secondname, email, phone FROM `users`")

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Firstname, &user.Secondname, &user.Email, &user.Phone)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, user)
	}

	return users
}
