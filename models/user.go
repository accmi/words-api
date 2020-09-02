package Models

import (
	"log"

	Config "github.com/accmi/words-api/config"
)

//GetAllUsers Fetch all user data
func GetAllUsers(users *[]User) (err error) {
	if err = Config.DB.Find(&users).Error; err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

//CreateUser Fetch all user data
func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(&user).Error; err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

// DeleteUser remove user
func DeleteUser(u *User, id string) error {
	if err := Config.DB.Where("id = ?", id).Delete(u).Error; err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
