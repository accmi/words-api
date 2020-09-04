package repositories

import (
	"context"
	Config "github.com/accmi/words-api/config"
	"github.com/accmi/words-api/models"
	"log"
)

// CreateUser to table users
func CreateUser(u *models.User) error {
	var err error = nil
	var id string

	err = Config.DB.QueryRow(context.Background(),
		`INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`,
		u.Email,
		u.PasswordsHash).Scan(&id)

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