package controllers

import (
	"fmt"
	"net/http"

	"github.com/xiaosasori/go-gallery/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

type Users struct {
	NewView *views.View
}

// This is used to render the form where a user can create
// a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// This is used to process the signup form when a user tries to
// create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	sf := &SignupForm{}
	parseForm(r, sf)
	fmt.Fprintln(w, "Email is", sf.Email)
	fmt.Fprintln(w, "Password is", sf.Password)
}
