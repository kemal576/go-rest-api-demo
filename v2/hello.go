package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	. "github.com/kemal576/go-rest-api-demo/v2/models"
	. "github.com/kemal576/go-rest-api-demo/v2/repositories"
)

var userRepository = new(UserRepository)

func main() {
	userRepository.AppendUsers()
	handleRequests()
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/GetAll", GetAll).Methods("GET")
	r.HandleFunc("/GetById/{id:[0-9]+}", GetById).Methods("GET")
	r.HandleFunc("/GetActiveUsers", GetActiveUsers).Methods("GET")
	r.HandleFunc("/GetByUsername/{username}", GetByUsername).Methods("GET")
	r.HandleFunc("/GetByAgeFilter", GetByAgeFilter).Methods("GET").Queries("min", "{min:[0-9,]+}", "max", "{max:[0-9,]+}")
	r.HandleFunc("/Add", Add).Methods("POST")
	r.HandleFunc("/Update", Update).Methods("PUT")
	r.HandleFunc("/Delete/{id:[0-9]+}", Delete).Methods("DELETE")

	http.Handle("/", r)
	http.ListenAndServe(":5764", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(userRepository.GetAll())
	checkError(err)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(json))
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	checkError(err)
	user, err := userRepository.GetById(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		os.Exit(1)
	}
	json, err := json.Marshal(user)
	checkError(err)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(json))
}

func GetByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	user, err := userRepository.GetByUsername(username)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		os.Exit(1)
	}
	json, err := json.Marshal(user)
	checkError(err)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(json))
}

func GetByAgeFilter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	min, _ := strconv.Atoi(vars["min"])
	max, _ := strconv.Atoi(vars["max"])
	users, err := userRepository.GetByAgeFilter(min, max)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		os.Exit(1)
	}
	json, err := json.Marshal(users)
	checkError(err)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(json))
}

func Add(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Fatalf("Unable to decode: %v", err)
	}

	userRepository.Add(&newUser)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Kullanıcı başarıyla eklendi")
}

func Update(w http.ResponseWriter, r *http.Request) {
	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		log.Fatalf("Unable to decode: %v", err)
		os.Exit(1)
	}

	if updatedUser.ID == 0 {
		fmt.Fprintf(w, "Geçersiz kullanıcı bilgisi gönderildi!")
		os.Exit(1)
	}
	userRepository.Update(&updatedUser)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Kullanıcı başarıyla güncellendi")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	checkError(err)
	var check = userRepository.Delete(id)
	if check {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Kullanıcı başarıyla silindi")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Kullanıcı silinirken bir hata oluştu")
	}
}

func GetActiveUsers(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(userRepository.GetActiveUsers())
	checkError(err)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(json))
}
