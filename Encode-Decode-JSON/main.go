package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Decode Request
	var rr Request

	err := json.NewDecoder(r.Body).Decode(&rr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Print
	fmt.Printf("%+v\n", rr)

	//Write Response Encode
	resp := Response{
		Code:    http.StatusOK,
		Message: "Thank You",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)

}

// func main() {
// 	http.HandleFunc("/", handler)

// 	fmt.Println("Starting Server...")

// 	err := http.ListenAndServe("127.0.0.1:8000", nil)

// 	if err != nil {
// 		panic("errors http")
// 	}
// }
