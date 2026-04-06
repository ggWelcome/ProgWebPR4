package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	data := []byte(`{"name":"Solar Plant","power":500,"status":"active"}`)
	resp, err := http.Post("http://localhost:8080/generators", "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))
}
