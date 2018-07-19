package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/decentralisedkev/Neo-Go-API/node"
)

func main() {

	plan, _ := ioutil.ReadFile("main.json")
	var data NodeList
	err := json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println("There was an error ", err)
	}
	var wg sync.WaitGroup
	for _, element := range data.Sites {
		wg.Add(1)
		go fmt.Println(element.GetLatency())
	}

}

// We want to collect as much informaton as we can about a block/transaction
// receiving the transaction, we can store it's type, time, amount?
// do we need to store trans?
type NodeList struct {
	Name     string      `json:"name"`
	PollTime string      `json:"pollTime"`
	Sites    []node.Node `json:"sites"`
}
