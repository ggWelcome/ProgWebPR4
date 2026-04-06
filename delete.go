package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/generators?id=1", nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status) // Очікувано: 204 No Content
}
