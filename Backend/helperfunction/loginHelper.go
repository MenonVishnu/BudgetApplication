package helperfunctions

import (
	"log"
	"net/http"
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(database.GetEnvValue("SECRET_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

func ValidateToken(tokenString string) (Claims, *jwt.Token) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(database.GetEnvValue("SECRET_KEY")), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return *claims, token
}

// middleware for Authentication for User
func AuthMiddlewareForUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			log.Fatal(err)
		}

		claims, token := ValidateToken(cookie.Value)
		if !token.Valid {
			//error message
		}

		if claims.Role != "User" {
			//error message
		}

		next.ServeHTTP(w, r)
	})
}

// middleware for Authentication for Admin
func AuthMiddlewareForAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			log.Fatal(err)
		}

		claims, token := ValidateToken(cookie.Value)
		if !token.Valid {
			//error message
		}

		if claims.Role != "Admin" {
			//error message
		}

		next.ServeHTTP(w, r)

	})
}
