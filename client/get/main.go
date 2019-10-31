package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fanadol/golang-distribute-tracing-example/models"
)

func main() {
	var data []models.Post
	client := &http.Client{}
	url := "http://localhost:8080/post"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("Error when trying to create new request : " + err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		panic("Error when trying to do request: " + err.Error())
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	fmt.Println(data)
}