package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	httpClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	if err != nil {
		log.Printf("error creating request: %v", err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Second)
	req = req.WithContext(ctx)

	defer cancel()

	res, err := httpClient.Do(req)
	if err != nil {
		log.Printf("error making a request: %v", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("error reading the response body: %v", err)
		return
	}

	fmt.Println(string(body))
}
