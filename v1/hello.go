package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	. "github.com/kemal576/go-rest-api-demo/models"
	. "github.com/kemal576/go-rest-api-demo/repositories"
)

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
	keys, ok := r.URL.Query()["id"]
	if ok {
		id, err := strconv.Atoi(keys[0])
		checkError(err)
		user, err := userRepository.GetById(id)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		json, err := json.Marshal(user)
		checkError(err)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(json))
	} else {
		fmt.Fprintf(w, "Lütfen id değeri yollayınız!")
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	var newUser User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Verileri alırken hata oluştu! :(")
	}

	json.Unmarshal(body, &newUser)
	userRepository.Add(&newUser)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Kullanıcı başarıyla eklendi")
}

func Update(w http.ResponseWriter, r *http.Request) {
	var updatedUser User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Verileri alırken hata oluştu! :(")
	}

	json.Unmarshal(body, &updatedUser)
	if updatedUser.ID == 0 {
		fmt.Fprintf(w, "Geçersiz kullanıcı bilgisi gönderildi!")
		os.Exit(1)
	}
	userRepository.Update(&updatedUser)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Kullanıcı başarıyla güncellendi")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if ok {
		id, err := strconv.Atoi(keys[0])
		checkError(err)
		var check = userRepository.Delete(id)
		if check {
			fmt.Fprintf(w, "Kullanıcı başarıyla silindi")
		} else {
			fmt.Fprintf(w, "Kullanıcı silinirken bir hata oluştu")
		}

		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Fprintf(w, "ID bilgisi alınamadı")
	}
}

func GetActiveUsers(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(userRepository.GetActiveUsers())
	checkError(err)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(json))
}

func handleRequests() {
	http.HandleFunc("/GetAll", GetAll)
	http.HandleFunc("/GetById", GetById)
	http.HandleFunc("/Add", Add)
	http.HandleFunc("/Update", Update)
	http.HandleFunc("/Delete", Delete)
	http.HandleFunc("/GetActiveUsers", GetActiveUsers)

	log.Fatal(http.ListenAndServe(":5764", nil))
}

var userRepository = new(UserRepository)

func main() {
	userRepository.AppendUsers()
	handleRequests()

}
