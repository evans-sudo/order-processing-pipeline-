package main

import (
	"encoding/json"
	"log"
)



func main() {
 for _, rawOrder := range rawOrders {
	var newOrders order 
	err := json.Unmarshal([]byte(rawOrder), &newOrders)
	if err != nil {
		log.Print(err)
		continue
	}
 }
}




var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}
