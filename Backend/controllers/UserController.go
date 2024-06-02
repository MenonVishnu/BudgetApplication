package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/vishnumenon/budgetapplication/database"
	helperfunctions "github.com/vishnumenon/budgetapplication/helperfunction"
	"github.com/vishnumenon/budgetapplication/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	//creating a reference model variable
	var user models.User

	//decoding the JSON and assigning those values into the reference model variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	//performing validation
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		// log.Fatal(errors)
		return
	}
	//performing validation for Role
	if !helperfunctions.ValidateRole(user) {
		http.Error(w, fmt.Sprintf("Validation error: %s", "Undefined Role"), http.StatusBadRequest)
		return
	}

	//check whether a user with the same email-ID exists or not - done
	if database.CheckUser(user.Email) {
		http.Error(w, fmt.Sprintf("User Aldready Exists: Please Login"), http.StatusBadRequest)
		return
	}

	//before saving it into the database you need to encrypt the password.
	user.Password = database.EncryptPassword(user.Password)

	user.ID = database.AddUser(user)
	fmt.Println(user.ID)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)

	//Validation of updated user
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		// log.Fatal(errors)
		return
	}
	//performing validation for Role
	if !helperfunctions.ValidateRole(user) {
		http.Error(w, fmt.Sprintf("Validation error: %s", "Undefined Role"), http.StatusBadRequest)
		return
	}

	database.UpdateUser(user, params["id"])
	user.ID, _ = primitive.ObjectIDFromHex(params["id"])
	json.NewEncoder(w).Encode(user)
}

// TODO: Delete User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	database.DeleteUser(params["id"])
	json.NewEncoder(w).Encode("User Deleted Successfully")
	
}

//TODO: Delete All User
func DeleteAllUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	database.DeleteAllUser()
	json.NewEncoder(w).Encode("Database Truncated Successfully!!")
}

//TODO: Get One User

//TODO: Get All User
