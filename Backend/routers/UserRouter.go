package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vishnumenon/budgetapplication/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	//testing route
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello From Port 4000")
	}).Methods("GET")

	//TODO: Create a sub router for user and admin

	//userprotected routes:
	//user routes
	router.HandleFunc("/api/user", controllers.AddUser).Methods("POST")           //add user route
	router.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT")    //update user route
	router.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods("DELETE") //delete user route
	router.HandleFunc("/api/users", controllers.DeleteAllUser).Methods("DELETE")  //delete all user route
	router.HandleFunc("/api/user/{id}", controllers.GetOneUser).Methods("GET")    //get all user route
	router.HandleFunc("/api/users", controllers.GetAllUser).Methods("GET")        //get all user route

	//Signin & Signout routes
	router.HandleFunc("/api/login", controllers.LogIn).Methods("POST") //Login route

	//adminprotected routes
	//budget routes
	router.HandleFunc("/api/budget", controllers.AddBudget).Methods("POST")                         //add budget route
	router.HandleFunc("/api/budget/{id}", controllers.UpdateBudget).Methods("PUT")                  //update budget route
	router.HandleFunc("/api/budget/{id}", controllers.DeleteOneBudget).Methods("DELETE")            //delete one budget route
	router.HandleFunc("/api/budgets/user/{id}", controllers.DeleteAllUsersBudget).Methods("DELETE") //delete all budgets of a specific user
	router.HandleFunc("/api/budgets", controllers.DeleteAllBudget).Methods("DELETE")                //delete all budgets
	router.HandleFunc("/api/budget/{id}", controllers.GetOneBudget).Methods("GET")                  //Get one budget
	router.HandleFunc("/api/budgets/user/{id}", controllers.GetAllUserBudget).Methods("GET")        //Get all budgets of a specific user
	router.HandleFunc("/api/budgets", controllers.GetAllBudgets).Methods("GET")                     //Get all budgets

	return router
}
