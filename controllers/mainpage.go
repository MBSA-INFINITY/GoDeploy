package controllers

import (
	"expense-tracker/auth"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

type Expense struct {
	Id          string
	Description string
	Amount      int
}

var expense1 Expense = Expense{randomString(10), "Car", 15000}
var all_expenses []Expense = []Expense{expense1}

func Start(w http.ResponseWriter, r *http.Request) {
	if auth.IsAuthenticated(w, r) {
		w.Header().Add("Content-Type", "text/html; charset=UTF-8")
		fp := path.Join("templates", "index.html")

		templ, err := template.ParseFiles(fp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := templ.Execute(w, all_expenses); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func AddExpense(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	// Access form values
	expense_id := randomString(10)
	expense_description := r.FormValue("expenseDesc")
	expense_amount_str := r.FormValue("expenseAmount")
	expense_amount, _ := strconv.Atoi(expense_amount_str)
	newExpense := Expense{expense_id, expense_description, expense_amount}
	all_expenses = append(all_expenses, newExpense)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for idx, val := range all_expenses {
		if val.Id == params["expense_id"] {
			all_expenses = append(all_expenses[:idx], all_expenses[idx+1:]...)
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func EditExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	expense_description := r.FormValue("expenseDesc")
	expense_amount_str := r.FormValue("expenseAmount")
	expense_amount, _ := strconv.Atoi(expense_amount_str)
	for idx, val := range all_expenses {
		if val.Id == params["expense_id"] {
			all_expenses[idx].Description = expense_description
			all_expenses[idx].Amount = expense_amount
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
