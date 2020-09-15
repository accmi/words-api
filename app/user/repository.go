package user

import (
	"context"
	"fmt"
	Config "github.com/accmi/words-api/config"
	"log"
)

// CreateUser to table users
func SaveUser(u *User) error {
	var err error = nil

	commandTag, err := Config.DB.Exec(context.Background(),
		"INSERT INTO users (email, password_hash) VALUES ($1, $2)",
		u.Email,
		u.PasswordsHash)


	if err != nil || commandTag.RowsAffected() != 1 {
		log.Println("Error query:", err)
		return err
	}

	return nil
}

// GetUserPasswordByEmail to find user
func GetUserPasswordByEmail(email string, ph *string) error {
	var err error = nil

	fmt.Println(email)

	err = Config.DB.QueryRow(context.Background(),
		"SELECT password_hash FROM users WHERE email=$1",
		email,
	).Scan(ph)

	if err != nil {
		log.Println("Error query:", err)
		return err
	}

	return err
}
