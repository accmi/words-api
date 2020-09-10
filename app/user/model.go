package user

import "golang.org/x/crypto/bcrypt"

// User model
type User struct {
	ID    string   `json:"id"`
	Email string `json:"email"`
	PasswordsHash string `json:"-"`
	Token string `json:"token"`
}
// HashPassword generate hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
// CheckPasswordHash check hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

////CreateUser Fetch all user data
//func (u *User) CreateUser() error {
//	return repositories.CreateUser(u)
//}
//
//// DeleteUser remove user
//func (u *User) DeleteUser(id int) error {
//	return repositories.DeleteUser(u, id)
//}
//// GetUsers includes all available users
//func GetUsers(us *[]User) error {
//	return repositories.GetUsers(us)
//}
