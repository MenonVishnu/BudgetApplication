package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vishnumenon/budgetapplication/controllers"
	helperfunctions "github.com/vishnumenon/budgetapplication/helperfunction"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	//testing route
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello From Port 4000")
	}).Methods("GET")

	//TODO: Create a sub router for user and admin
	userRouter := router.NewRoute().Subrouter()
	adminRouter := router.NewRoute().Subrouter()

	//TODO: Protected Routes are to be tested some routes can be accessed by both admin and user

	//middleware
	userRouter.Use(helperfunctions.AuthMiddlewareForUser)
	adminRouter.Use(helperfunctions.AuthMiddlewareForAdmin)
	
	//Signin & Signout routes
	router.HandleFunc("/api/login", controllers.LogIn).Methods("POST")   //Login route
	router.HandleFunc("/api/logout", controllers.Logout).Methods("POST") //Login route

	//user routes
	router.HandleFunc("/api/user", controllers.AddUser).Methods("POST")           //add user route
	userRouter.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT")    //update user route
	userRouter.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods("DELETE") //delete user route
	userRouter.HandleFunc("/api/user/{id}", controllers.GetOneUser).Methods("GET")    //get all user route
	
	//admin only
	adminRouter.HandleFunc("/api/users", controllers.GetAllUser).Methods("GET")        //get all user route
	adminRouter.HandleFunc("/api/users", controllers.DeleteAllUser).Methods("DELETE")  //delete all user route

	//budget routes
	userRouter.HandleFunc("/api/budget", controllers.AddBudget).Methods("POST")                         //add budget route
	userRouter.HandleFunc("/api/budget/{id}", controllers.UpdateBudget).Methods("PUT")                  //update budget route
	userRouter.HandleFunc("/api/budget/{id}", controllers.DeleteOneBudget).Methods("DELETE")            //delete one budget route
	userRouter.HandleFunc("/api/budgets/user/{id}", controllers.DeleteAllUsersBudget).Methods("DELETE") //delete all budgets of a specific user
	userRouter.HandleFunc("/api/budget/{id}", controllers.GetOneBudget).Methods("GET")                  //Get one budget
	userRouter.HandleFunc("/api/budgets/user/{id}", controllers.GetAllUserBudget).Methods("GET")        //Get all budgets of a specific user
	
	//both admin & user
	adminRouter.HandleFunc("/api/budgets/user/{id}", controllers.GetAllUserBudget).Methods("GET")        //Get all budgets of a specific user
	adminRouter.HandleFunc("/api/budgets/user/{id}", controllers.DeleteAllUsersBudget).Methods("DELETE") //delete all budgets of a specific user
	
	//admin only
	adminRouter.HandleFunc("/api/budgets", controllers.DeleteAllBudget).Methods("DELETE")                //delete all budgets
	adminRouter.HandleFunc("/api/budgets", controllers.GetAllBudgets).Methods("GET")                     //Get all budgets
	
	return router
}
