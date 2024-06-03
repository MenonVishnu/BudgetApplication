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
		models.ErrorResponse(w, 403, "Vallidation Error!!", errors.Error())
		return
	}
	//performing validation for Role
	if !helperfunctions.ValidateRole(user) {
		err := map[string]string{"Key": "User.Role", "Field": "Role Not Defined"}
		models.ErrorResponse(w, 403, "Vallidation Error!!", err)
		return
	}

	//check whether a user with the same email-ID exists or not - done
	if database.CheckUser(user.Email) == "" {
		err := map[string]string{"Key": "User.Email", "Field": "User with same Email ID exists, Please Login"}
		models.ErrorResponse(w, 409, "Vallidation Error!!", err)
		return
	}

	//before saving it into the database you need to encrypt the password.
	user.Password = helperfunctions.EncryptPassword(user.Password)

	user.ID = database.AddUser(user)
	message := "User created Successfully with ObjectId: " + user.ID.Hex()
	models.SuccessResponse(w, 201, message, user)
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
		models.ErrorResponse(w, 403, "Vallidation Error!!", errors.Error())
		return
	}
	//performing validation for Role
	if !helperfunctions.ValidateRole(user) {
		err := map[string]string{"Key": "User.Role", "Field": "Role Not Defined"}
		models.ErrorResponse(w, 403, "Vallidation Error!!", err)
		return
	}

	database.UpdateUser(user, params["id"])
	user.ID, _ = primitive.ObjectIDFromHex(params["id"])
	message := "User updated Successfully with ObjectId: " + user.ID.Hex()
	models.SuccessResponse(w, 201, message, user)
}

// Delete User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	database.DeleteUser(params["id"])
	message := "User deleted Successfully with ObjectId: " + params["id"]
	models.SuccessResponse(w, 201, message, nil)
}

// Delete All User
func DeleteAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	database.DeleteAllUser()
	message := "All Users Deleted Successfully"
	models.SuccessResponse(w, 201, message, nil)
}

// Get One User
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)

	user := database.GetUser(params["id"])
	if user.ID == primitive.NilObjectID {
		err := map[string]string{"Key": "User.Id", "Field": "User Not Found!"}
		message := "User not found with ObjectId: " + params["id"]
		models.ErrorResponse(w, 404, message, err)
		return
	}
	message := "User successfully retrieved with ObjectId: " + params["id"]
	models.SuccessResponse(w, 201, message, user)
}

// Get All Users
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	users := database.GetAllUser()
	if len(users) == 0 {
		err := map[string]string{"Key": "Database", "Field": "No User present in Database"}
		models.ErrorResponse(w, 404, "Database Error", err)
		return
	}

	message := "All users successfully retrieved!!"
	models.SuccessResponse(w, 201, message, users)

}

//login and logout feature

func LogIn(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)

	pass  := database.CheckUser(user.Email)

	//if no password then throw error because no user exists with that email
	if pass == ""{
		err := map[string]string{"Key": "Authentication", "Field": "Invalid Email or Password"}
		models.ErrorResponse(w, 401, "Authentication Error", err)
	}

	if helperfunctions.CheckPassword(pass, user.Password){
		fmt.Println("User Successfully Logged In")
	}

}
