package main

import (
	"math/rand"
	"sync/atomic"
	"time"
)

type Client struct {
	Id int64 `json:"id"`
}

var clientId int64
var orderId int64

func createClient() {
	for id := 0; id < NrClients; id++ {
		CreationOfClient()
	}
}

func (cl *Client) clientGenerateOrder() Order {
	order := Order{
		ClientId: int64(cl.Id),
		OrderId:  atomic.AddInt64(&orderId, 1),
		Orders:   make([]InformationOrder, 0),
	}
	var RestaurantId1 = rand.Intn(2) + 1
	var restautanId int64
	for i := 0; i < RestaurantId1; i++ {
		// var RestaurantId = rand.Intn(2) + 1
		var RestaurantId = atomic.AddInt64(&restautanId, 1)
		var maxFoodNr = rand.Intn(5) + 1
		var InformationOrder = InformationOrder{
			RestaurantId: int(RestaurantId),
			Items:        make([]int, maxFoodNr),
		}
		InformationOrder.Priority = 6 - maxFoodNr
		for id := 0; id < maxFoodNr; id++ {
			InformationOrder.Items[id] = rand.Intn(5) + 1
		}
		InformationOrder.MaxWait = float64(50) * 1.3
		InformationOrder.CreatedTime = time.Now().UnixNano()
		order.Orders = append(order.Orders, InformationOrder)
	}
	return order
}
