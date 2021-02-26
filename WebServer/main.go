package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

type User struct {
	ID    string
	Name  string
	Phone string
}

var users = []User{
	User{
		ID:    "0",
		Name:  "test1",
		Phone: "+201143145911",
	},
	User{
		ID:    "1",
		Name:  "test2",
		Phone: "+201143145912",
	},
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling req", r)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Hello World"}`)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		get(w, r)
	} else if r.Method == "POST" {
		post(w, r)
	} else {
		errorHandler(w, r, http.StatusInternalServerError, fmt.Errorf("invalid method"))
	}
	fmt.Println("handling req", r)
	fmt.Println("handling req", w)
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling get request")
	query := r.URL.Query()
	id := query.Get("id")
	var result []byte
	var err error
	if id == "" {
		result, err = json.Marshal(users)
	} else {
		intId, err := strconv.Atoi(id)
		if err == nil {
			result, err = json.Marshal(users[intId])
		}
	}

	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(result))
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling post request")
	var user User
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("ddddd2")
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}

	users = append(users, user)
	fmt.Println("ddddd")

	w.WriteHeader(http.StatusCreated)
}

func errorHandler(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	w.WriteHeader(statusCode)

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `{error:%v`, err.Error())
}
