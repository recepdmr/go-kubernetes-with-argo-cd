package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi from new version"))
	})

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	fmt.Printf("application running %s \n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
