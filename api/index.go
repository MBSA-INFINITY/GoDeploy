package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("This is MBSA's Expense Tracker")

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("I am mbsaiaditya"))
	})
	http.ListenAndServe(":8080", r)

}
