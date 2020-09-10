package user

import (
	"context"
	"github.com/accmi/words-api/app/utils"
	Config "github.com/accmi/words-api/config"
	"log"
)

// CreateUser to table users
func SaveUser(u *User) utils.ErrorsResponseInterface {
	var err error = nil

	commandTag, err := Config.DB.Exec(context.Background(),
		"INSERT INTO users (email, password_hash) VALUES ($1, $2)",
		u.Email,
		u.PasswordsHash)


	if err != nil || commandTag.RowsAffected() != 1 {
		log.Panicln("Error query:", err)
		return utils.ErrorsResponse{}.GetError(err.Error())
	}

	return nil
}

// CheckUser to check if user exists
func CheckUser(u *User) error {
	var err error = nil
	var id string
	var email string

	err = Config.DB.QueryRow(context.Background(),
		"SELECT id, email FROM users WHERE email=$1 and password_hash=$2",
		u.Email,
		u.PasswordsHash).Scan(&id, &email)

	u.ID = id

	if err != nil {
		log.Panicln("Error query:", err)
		return err
	}
	return  err
}

//// DeleteUser by id
//func DeleteUser(u *models.User, id int) error {
//	var err error = nil
//	var commandTag pgconn.CommandTag
//
//	commandTag, err = Config.DB.Exec(context.Background(),`DELETE FROM users WHERE id=$1`, id)
//
//	if err != nil {
//		log.Panicln("Error query:", err)
//		return err
//	}
//
//	if commandTag.RowsAffected() != 1 {
//		err = errors.New("not found")
//		return err
//	}
//
//	return  err
//}
//// GetUsers gets all available users from table users
//func GetUsers(us *[]models.User) error {
//	var err error = nil
//	var rows pgx.Rows
//	rows, err = Config.DB.Query(context.Background(), "SELECT id, name, email FROM users")
//
//	if err != nil {
//		log.Panicln("Error query:", err)
//		return err
//	}
//
//	defer rows.Close()
//
//	for rows.Next() {
//		var id int
//		var name string
//		var email string
//
//		err = rows.Scan(&id, &name, &email)
//		if err != nil {
//			log.Panicln(err)
//			break
//		}
//
//		user := models.User{
//			Name: name,
//			ID: id,
//			Email: email,
//		}
//
//		*us = append(*us, user)
//	}
//
//	return  err
//}