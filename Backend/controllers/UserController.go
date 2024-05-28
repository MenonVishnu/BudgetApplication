package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vishnumenon/budgetapplication/database"
	"github.com/vishnumenon/budgetapplication/models"
	"go.mongodb.org/mongo-driver/bson"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	database.AddUser(user)
	fmt.Println(user.ID)
	json.NewEncoder(w).Encode(user)
}
