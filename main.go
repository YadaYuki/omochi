package main

import (
	"encoding/json"
	"net/http"
)

type ResponseBody struct {
	Message string `json:"message"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	r := ResponseBody{Message: "Hello World"}
	resBody, _ := json.Marshal(r)
	w.Write(resBody)
}

func main() {

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8090", nil)
}
