package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Req struct {
	Name string `json:"name"`
}

type Res struct {
	Msg string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Req

	de := json.NewDecoder(r.Body)
	de.DisallowUnknownFields()

	err := de.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "\"name\" is required", http.StatusBadRequest)
		return
	}

	res := Res{
		Msg: "Hello " + req.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func main() {
	http.HandleFunc("/hello", HelloHandler)

	fmt.Println("OK! Server starting nowwWW")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
