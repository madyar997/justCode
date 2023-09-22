package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", helloWorld)
	log.Println("server is starting on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("new hello world request at ", time.Now().UTC())
	time.Sleep(5 * time.Second)
	w.Write([]byte("hello world"))
}
