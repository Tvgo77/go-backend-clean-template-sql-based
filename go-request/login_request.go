package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Login() {
	url := "http://localhost:8080/login"
	jsonData := []byte("{\"email\":\"test@gmail.com\", \"password\":\"password\"}")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

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