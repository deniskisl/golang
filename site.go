package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  uint16
}

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<H1>Main page</H1> <p><a href="../Rob/">Rob</a></p>
	<p><a href="../Bob/">Bob</a></p>`)
	// rob := User{"Rob", 40}
	// bob := User{"Bob", 18}
	// tmpl1, _ := template.ParseFiles("templates/home_page.html")
	// tmpl2, _ := template.ParseFiles("templates/bob_page.html")
	// tmpl1.Execute(w, rob)
	// tmpl2.Execute(w, bob)
}

func bob_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 18}
	tmpl2, _ := template.ParseFiles("templates/bob_page.html")
	tmpl2.Execute(w, bob)
	// fmt.Fprint(w, bob.getAllInfo())
}

func rob_page(w http.ResponseWriter, r *http.Request) {
	rob := User{"Rob", 40}
	tmpl3, _ := template.ParseFiles("templates/rob_page.html")
	tmpl3.Execute(w, rob)
	// fmt.Fprint(w, rob.getAllInfo())
}

// func (u *User) getAllInfo() string {
// 	return fmt.Sprintf("Username is %s, age is %d", u.name, u.age)
// }

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/Bob/", bob_page)
	http.HandleFunc("/Rob/", rob_page)
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()
}
