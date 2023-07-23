package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "It works!")
	})

	go func() {
		for {
			faultyAppPort := readFile("port.txt")
			time.Sleep(time.Second * 5)
			_, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s", faultyAppPort))
			if err != nil {
				fmt.Println("It's not working :(")
			} else {
				fmt.Println("It works! :)")
			}
		}
	}()

	port := "8080"
	fmt.Println("Starting HTTP ...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting HTTP server: %s\n", err.Error())
	}
}

func readFile(filename string) string {
	for {
		content, err := os.ReadFile(filename)
		if err == nil {
			return string(content)
		}

		fmt.Printf("Something went wrong...\n")
		time.Sleep(time.Second * 5)
	}
}
