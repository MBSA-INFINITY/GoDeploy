package handler

import (
	"fmt"
	"net/http"
)

func mbsa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to API by LearnCodeOnline</h1>")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", mbsa)
	http.ListenAndServe(":8080", nil)

}
