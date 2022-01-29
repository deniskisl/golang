package main

import ("fmt"
	"net/http")
	
type User struct{
	name string
	age uint16
	}
	
func home_page(w http.ResponseWriter, r *http.Request){
	
}

func bob_page(w http.ResponseWriter, r *http.Request){
	bob := User{"Bob", 18}
	fmt.Fprint(w, bob.getAllInfo())
}

func rob_page(w http.ResponseWriter, r *http.Request){
	rob := User{"Rob", 40}
	fmt.Fprint(w, rob.getAllInfo())
}

func (u *User) getAllInfo() string{
	return fmt.Sprintf("Username is %s, age is %d", u.name, u.age)
}

func handleRequest(){
	http.HandleFunc("/", home_page)
	http.HandleFunc("/Bob/", bob_page)
	http.HandleFunc("/Rob/", rob_page)
	http.ListenAndServe(":8080", nil)
}

func main(){

	handleRequest()
}
