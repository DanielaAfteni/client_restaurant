package main

type Order struct {
	OrderId  int64              `json:"order_id"`
	ClientId int64              `json:"client_id"`
	Orders   []InformationOrder `json:"orders"`
}

type InformationOrder struct {
	RestaurantId int     `json:"restaurant_id"`
	Items        []int   `json:"items"`
	Priority     int     `json:"priority"`
	MaxWait      float64 `json:"max_wait"`
	CreatedTime  int64   `json:"created_time"`
}
