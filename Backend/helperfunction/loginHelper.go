package helperfunctions

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vishnumenon/budgetapplication/database"
	"github.com/vishnumenon/budgetapplication/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	ID    primitive.ObjectID `json:"_id"`
	Email string             `json:"email"`
	Role  string             `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(user models.User) string {
	claims := Claims{
		user.ID,
		user.Email,
		string(user.Role),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenString.SignedString([]byte(database.GetEnvValue("SECRET_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	return token
}
