package main

import (
	"fmt"
	"golang-session/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	// fs := http.FileServer(http.Dir("assets"))

	// using http
	http.HandleFunc("/students", controller.HandlingHttpReq)

	// using gorilla mux router
	router.HandleFunc("/todo", controller.HandlingTodosGETReq).Methods("GET")
	router.HandleFunc("/todo", controller.HandlingTodosPostReq).Methods("POST")
	router.HandleFunc("/todo", controller.HandlingTodosDeleteReq).Methods("DELETE")
	router.HandleFunc("/todo", controller.HandlingTodosPutReq).Methods("PUT")
	// mux.HandleFunc("/", Index)

	// router.Handle("/assets/", http.StripPrefix("/assets", fs))
	fmt.Println("started our server in 8080")
	http.ListenAndServe("localhost:8080", router)
}

// var tpl *template.Template

// func init() {
// 	tpl = template.Must(template.ParseFiles("index.html"))
// }

// func Index(w http.ResponseWriter, r *http.Request) {
// 	c := getCookie(w, r)
// 	tpl.ExecuteTemplate(w, "index.html", c)
// }

// func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
// 	c, err := r.Cookie("CurrentSession")
// 	if err != nil {
// 		csId := uuid.New()
// 		c := &http.Cookie{
// 			Name:  "CurrentSession",
// 			Value: csId.String(),
// 		}
// 		http.SetCookie(w, c)
// 	}
// 	return c
// }
