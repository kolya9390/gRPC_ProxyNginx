package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		// Логика обработки запроса на загрузку изображений
		fmt.Fprintf(w, "Image loaded successfully")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}