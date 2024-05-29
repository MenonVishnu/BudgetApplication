package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vishnumenon/budgetapplication/database"
	"github.com/vishnumenon/budgetapplication/models"
	// "go.mongodb.org/mongo-driver/bson"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	//creating a reference model variable
	var user models.User

	//decoding the JSON and assigning those values into the reference model variable
	_ = json.NewDecoder(r.Body).Decode(&user)
	database.AddUser(user)
	fmt.Println(user.ID)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	database.UpdateUser(user, params["id"])
	json.NewEncoder(w).Encode(user)
}
