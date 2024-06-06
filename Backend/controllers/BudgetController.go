package controllers

import (
	"encoding/json"
	"net/http"

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

	//TODO: Add date and User inside the budget 

	//decoding the JSON and assigning those values into the reference model variable
	_ = json.NewDecoder(r.Body).Decode(&budget)

	//performing validation
	validate := validator.New()
	err := validate.Struct(budget)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		models.ErrorResponse(w, 403, "Vallidation Error!!", errors.Error())
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

	//Validation of updated user
	validate := validator.New()
	err := validate.Struct(budget)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		models.ErrorResponse(w, 403, "Vallidation Error!!", errors.Error())
		return
	}

	database.UpdateBudget(budget, params["id"])
	budget.ID, _ = primitive.ObjectIDFromHex(params["id"])
	message := "User updated Successfully with ObjectId: " + budget.ID.Hex()
	models.SuccessResponse(w, 201, message, budget)

}

//admin 
func DeleteAllBudget(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	database.DeleteAllBudget()
	message := "All Budget Deleted Successfully"
	models.SuccessResponse(w,201,message,nil)
}

//user
func DeleteOneBudget(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params := mux.Vars(r)
	database.DeleteBudget(params["id"])
	message := "Budget deleted Successfully with ObjectId: " + params["id"]
	models.SuccessResponse(w, 201, message, nil)
}

//admin / user
func DeleteAllUsersBudget(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params := mux.Vars(r)
	database.DeleteAllUsersBudget(params["id"])
	message := "All Budget deleted Successfully related to User ObjectId: " + params["id"]
	models.SuccessResponse(w, 201, message, nil)
}

func GetOneBudget(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

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
