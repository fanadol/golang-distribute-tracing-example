package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fanadol/golang-distribute-tracing-example/models"
	"github.com/nats-io/stan.go"
)

func main() {
	// Create new post to server
	client := &http.Client{}
	post := models.Post{
		Title: "Cat and Dogs",
		Body:  "Once upon a time in hohaw land",
	}
	payload, _ := json.Marshal(post)
	url := "http://localhost:8080/post"

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		panic("Error when trying to create new request : " + err.Error())
	}

	_, err = client.Do(request)
	if err != nil {
		panic("Error when trying to do request: " + err.Error())
	}

	fmt.Println("Success HTTP POST!")

	// Publish message
	sc, err := stan.Connect("test-cluster", "client-clientID")
	if err != nil {
		panic("Error when connecting to stan: " + err.Error())
	}

	err = sc.Publish("foo", []byte("Hellow Sir"))
	if err != nil {
		panic("Error when trying to publish: " + err.Error())
	}

	fmt.Println("Success Publish to Broker")
}
