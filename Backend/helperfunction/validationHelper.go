package helperfunctions

import (
	"log"

	"github.com/vishnumenon/budgetapplication/models"
	"golang.org/x/crypto/bcrypt"
)

func ValidateRole(user models.User) bool {
	if user.Role == "Admin" || user.Role == "User" {
		return true
	}
	return false
}

// Encrypting Password
func EncryptPassword(password string) string {
	passwordByte := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}

// Comparing Given password with already present password
func CheckPassword(hashedPassword string, givenPassword string) bool {
	hashedPasswordByte := []byte(hashedPassword)
	givenPasswordByte := []byte(givenPassword)

	err := bcrypt.CompareHashAndPassword(hashedPasswordByte, givenPasswordByte)
	if err != nil {
		return false
	}
	return true
}

func GenerateToken() {

}
