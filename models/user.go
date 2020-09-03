package models

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"log"

	Config "github.com/accmi/words-api/config"
)

// User model
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//CreateUser Fetch all user data
func (u *User) CreateUser() error {
	var err error = nil
	var id uint

	err = Config.DB.QueryRow(context.Background(),
		`INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`,
		u.Name,
		u.Email).Scan(&id)

	u.ID = id

	if err != nil {
		log.Panicln("Error query:", err)
		return err
	}
	return  err
}

// DeleteUser remove user
func (u *User) DeleteUser(id string) error {
	var err error = nil
	var commandTag pgconn.CommandTag

	commandTag, err = Config.DB.Exec(context.Background(),`DELETE FROM users WHERE id=$1`, id)

	if err != nil {
		log.Panicln("Error query:", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		err = errors.New("not found")
		return err
	}

	return  err
}

// Users just slice of user
type Users []User

//GetAllUsers Fetch all user data
func (us *Users) GetAllUsers() error {
	var err error = nil
	var rows pgx.Rows
	rows, err = Config.DB.Query(context.Background(), "SELECT id, name, email FROM users")

	if err != nil {
		log.Panicln("Error query:", err)
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint
		var name string
		var email string

		err = rows.Scan(&id, &name, &email)
		if err != nil {
			log.Panicln(err)
			break
		}

		user := User{
			Name: name,
			ID: id,
			Email: email,
		}

		*us = append(*us, user)
	}

	return  err
}
