package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/vishnumenon/budgetapplication/database"
	"github.com/vishnumenon/budgetapplication/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	//creating a reference model variable
	var budget models.Budget

	//decoding the JSON and assigning those values into the reference model variable
	_ = json.NewDecoder(r.Body).Decode(&budget)

	//Add date and User inside the budget
	budget.Date = primitive.NewDateTimeFromTime(time.Now())
	user := database.GetUser(budget.User.ID.Hex())
	budget.User = &user
	// dummy password inside this so that validation is successfull
	budget.User.Password = "dummypasswordhacker"

	//performing validation
	validate := validator.New()
	err := validate.Struct(budget)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		models.ErrorResponse(w, 403, "Vallidation Error!!", errors.Error())
		return
	}

	//if user is admin then do not allow him to add budget
	if budget.User.Role == "Admin" {
		err := map[string]string{"Key": "Role", "Field": "Admin Not allowed for adding budget"}
		models.ErrorResponse(w, 403, "Database Error", err)
		return
	}

	budget.ID = database.AddBudget(budget)
	message := "Budget created Successfully with ObjectId: " + budget.ID.Hex()
	models.SuccessResponse(w, 201, message, budget)
}


func UpdateBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	var budget models.Budget

	_ = json.NewDecoder(r.Body).Decode(&budget)

	//add updated time as well
	budget.Date = primitive.NewDateTimeFromTime(time.Now())
	//Validation of updated budget
	validate := validator.New()
	err := validate.Struct(budget)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		models.ErrorResponse(w, 403, "Vallidation Error!!", errors.Error())
		return
	}

	database.UpdateBudget(budget, params["id"])
	budget.ID, _ = primitive.ObjectIDFromHex(params["id"])
	message := "Budget updated Successfully with ObjectId: " + budget.ID.Hex()
	models.SuccessResponse(w, 201, message, budget)

}

// admin
func DeleteAllBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deletedCount := database.DeleteAllBudget()
	if deletedCount == 0 {
		err := map[string]string{"Key": "Database", "Field": "No Budgets present in Database"}
		models.ErrorResponse(w, 404, "Database Error", err)
		return
	}
	message := "All Budget Deleted Successfully"
	models.SuccessResponse(w, 201, message, nil)
}

// user
func DeleteOneBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deletedCount := database.DeleteBudget(params["id"])
	if deletedCount == 0 {
		err := map[string]string{"Key": "Database", "Field": "No Budgets present in Database with ObjectId: " + params["id"]}
		models.ErrorResponse(w, 404, "Database Error", err)
		return
	}
	message := "Budget deleted Successfully with ObjectId: " + params["id"]
	models.SuccessResponse(w, 201, message, nil)
}

// admin / user
func DeleteAllUsersBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deletedCount := database.DeleteAllUsersBudget(params["id"])
	if deletedCount == 0 {
		err := map[string]string{"Key": "Database", "Field": "No Budgets present associated with UserId:  " + params["id"]}
		models.ErrorResponse(w, 404, "Database Error", err)
		return
	}
	message := "All Budget deleted Successfully related to User ObjectId: " + params["id"]
	models.SuccessResponse(w, 201, message, nil)
}

// admin
func GetAllBudgets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	budgets := database.GetAllBudgets()
	if len(budgets) == 0 {
		err := map[string]string{"Key": "Database", "Field": "No Budgets present in Database"}
		models.ErrorResponse(w, 404, "Database Error", err)
		return
	}

	message := "All Budgets successfully retrieved!!"
	models.SuccessResponse(w, 201, message, budgets)
}

// user
func GetOneBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	budget := database.GetOneBudget(params["id"])

	if budget.ID == primitive.NilObjectID {
		err := map[string]string{"Key": "Budget.Id", "Field": "No Budget with ID found"}
		message := "No Budget with ID found, ObjectID: " + params["id"]
		models.ErrorResponse(w, 404, message, err)
		return
	}
	message := "Budget successfully retrieved with ObjectId: " + params["id"]
	models.SuccessResponse(w, 201, message, budget)
}

// admin / user
func GetAllUserBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	budgets := database.GetAllUserBudget(params["id"])

	if len(budgets) == 0 {
		err := map[string]string{"Key": "Database", "Field": "No Budgets present for the Current User in Database"}
		models.ErrorResponse(w, 404, "Database Error", err)
		return
	}
	message := "Budget successfully retrieved for User with ObjectId: " + params["id"]
	models.SuccessResponse(w, 201, message, budgets)
}
