package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Human struct {
	Fname string
	Lname string
	Age   int
}

func (hum Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := Human{
		Fname: "havva nur",
		Lname: "durudeniz",
		Age:   22,
	}

	bytes, err := json.Marshal(message)

	if err != nil {
		panic(err)
	}

	//w.Write([]byte(string(bytes)))
	fmt.Fprintf(w, string(bytes))
	//formu parse etmek için
	r.ParseForm()

	//fmt.Fprintf(w, "name :%s\n lastname:%s\n age:%d\n", hum.Fname, hum.Lname, hum.Age)

	//path bilgisi alınır.
	fmt.Println("deneme")
	fmt.Fprintf(w, "path:%s", r.URL.Path[1:])

}

func main() {
	var h Human

	http.ListenAndServe(":9000", h)
}
