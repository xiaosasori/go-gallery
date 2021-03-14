package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xiaosasori/go-gallery/controllers"
	"github.com/xiaosasori/go-gallery/models"
)

const (
	host     = "arjuna.db.elephantsql.com"
	port     = 5432
	user     = "hwrvxfng"
	password = "RnzDzwy5Jg-1L9tI9hedGV3M1ykyZcpG"
	dbname   = "hwrvxfng"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	must(err)
	defer us.Close()
	// us.DestructiveReset()
	us.AutoMigrate()
	// user := models.User{
	// 	Name:  "MA",
	// 	Email: "ma@gmail.com",
	// }
	// if err := us.Create(&user); err != nil {
	// 	panic(err)
	// }
	user, err := us.ByID(1)
	must(err)
	fmt.Println(user)
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/signup", usersC.NewView).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
