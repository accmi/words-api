package user

import "golang.org/x/crypto/bcrypt"

// User model
type User struct {
	ID    string   `json:"-"`
	Email string `json:"-"`
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
