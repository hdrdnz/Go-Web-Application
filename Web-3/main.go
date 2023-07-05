package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID    int    "json:id"
	Fname string "json:firstname"
	Lname string "json:lastname"
	Age   int    "json:age"
}

func main() {
	apiRoot := "/api"

	//.../api
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home"}
		//struct-json tipine byte türünde dönüştürülür.
		byte, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(byte))
	})

	//.../api/users
	http.HandleFunc(apiRoot+"/users", Users)

	//../api/me
	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		me := User{
			ID:    1,
			Fname: "Havva Nur",
			Lname: "Durudeniz",
			Age:   22,
		}
		//api üzerinden ulaştığın için message kısmına atmalısın.
		message := me

		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":8080", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error:", err.Error())
		//çıkış yapılır.
		os.Exit(1)
	}
}

func Users(w http.ResponseWriter, r *http.Request) {

	users := []User{
		User{ID: 5, Fname: "ahmet", Lname: "yılmaz", Age: 54},
		User{ID: 5, Fname: "Nihat", Lname: "Erdem", Age: 78},
	}

	byte, err := json.Marshal(users)
	checkError(err)
	fmt.Fprintf(w, string(byte))
}
