package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Profile(uid string, token string) {
	url := "http://localhost:8080/profile/" + uid

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	req.Header.Set("Authorization", "Bearer " + token)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	
	fmt.Println("Response body:", string(body))
}