/*
* @Author: Administrator
* @Date:   2017-07-28 13:13:55
* @Last Modified by:   Administrator
* @Last Modified time: 2017-07-29 12:12:28
*/

/*

http://localhost:12345/people/2

go run main.go


 */
package main
 
import (

    "encoding/json"

    "log"
    "net/http"

    /*
    go get github.com/gorilla/mux

    https://github.com/gorilla/mux

     */
    "github.com/gorilla/mux"

    /*
    go get github.com/go-sql-driver/mysql
    https://github.com/go-sql-driver/mysql
     */
    
)
 
type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
 
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}
 
var people []Person
 
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
     SetHeader(w)
    json.NewEncoder(w).Encode(&Person{})
}
 
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	 SetHeader(w)
    json.NewEncoder(w).Encode(people)
}
 
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
     SetHeader(w)
    json.NewEncoder(w).Encode(people)
}
 
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
    }
     SetHeader(w)
    json.NewEncoder(w).Encode(people)
}
func SetHeader(w http.ResponseWriter){

	 w.Header().Set("Content-Type", "application/json")
}

/*
may retrive data from MySQL
 */
func getPeople(){

    people = append(people, Person{ID: "1", Firstname: "王", Lastname: "二狗", Address: &Address{City: "Dublin", State: "CA"}})
    people = append(people, Person{ID: "2", Firstname: "吕", Lastname: "桂花"})
  
}
 
func main() {
    router := mux.NewRouter()

    getPeople()

    router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

    log.Println("http://localhost:12345")
    log.Fatal(http.ListenAndServe(":12345", router))

}