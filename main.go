package main

import (
	// Package json implements encoding and decoding of JSON.
	// The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions.

	// Package ioutil implements some I/O utility functions.

	// Package os provides a platform-independent interface to operating system functionality.

	// importing the gin, because is a high-performance HTTP web framework written in Golang (Go).

	"bytes"
	"encoding/json"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const NrClients = 3
const TIME_UNIT = 250

func main() {
	router := gin.Default()
	createClient()
	router.Run(":8091")
}

func CreationOfClient() *Client {
	var client = &Client{}
	//cl.Id = rand.Intn(10000) + 1
	client.Id = atomic.AddInt64(&clientId, 1)
	order := client.clientGenerateOrder()
	jsonBody, err := json.Marshal(order)
	if err != nil {
		log.Err(err).Msg("Error!!")
	}
	contentType := "application/json"
	//_, err = http.Post("http://food_ordering_service_restaurant:8090/order", contentType, bytes.NewReader(jsonBody))
	_, err = http.Post("http://localhost:8090/order", contentType, bytes.NewReader(jsonBody))
	if err != nil {
		log.Err(err).Msg("Error!!!")
	}
	log.Printf("The new order was created by client: %+v", client.Id)
	return client
}
