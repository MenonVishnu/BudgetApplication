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

	//user routes
	router.HandleFunc("/api/user", controllers.AddUser).Methods("POST")        //add user route
	router.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT") //update user route

	return router
}
