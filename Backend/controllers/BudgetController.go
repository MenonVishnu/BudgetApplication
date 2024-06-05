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
