package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("⇨ internal-api: GET /process")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Hello from internal service"}`))
	})

	http.ListenAndServe(":8080", nil)
}
