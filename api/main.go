package handler

import (
	"expense-tracker/auth"
	"expense-tracker/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("This is MBSA's Expense Tracker")

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Start)
	r.HandleFunc("/login", auth.Login).Methods("GET", "POST")
	r.HandleFunc("/logout", auth.Logout).Methods("GET", "POST")
	r.HandleFunc("/expense", controllers.AddExpense).Methods("POST")
	r.HandleFunc("/delete/{expense_id}", controllers.DeleteExpense).Methods("POST")
	r.HandleFunc("/edit/{expense_id}", controllers.EditExpense).Methods("POST")
	http.ListenAndServe(":8080", r)

}
