package Models

import (
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
	if err := Config.DB.Create(&u).Error; err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

// DeleteUser remove user
func (u *User) DeleteUser(id string) error {
	if err := Config.DB.Where("id = ?", id).Delete(u).Error; err != nil {
		log.Panic(err)
		return err
	}

	return nil
}

// Users just slice of user
type Users []User

//GetAllUsers Fetch all user data
func (us *Users) GetAllUsers() error {
	if err := Config.DB.Find(&us).Error; err != nil {
		log.Panic(err)
		return err
	}
	return nil
}
