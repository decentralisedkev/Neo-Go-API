package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/decentralisedkev/Neo-Go-API/node"
)

func main() {

	plan, _ := ioutil.ReadFile("main.json")
	var data NodeList
	err := json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println("There was an error ", err)
	}
	for _, element := range data.Sites {
		block, err := element.GetBlock(200)
		if err != nil {
			continue
		}
		fmt.Printf("%+v\n", element)
		fmt.Println(block.Hash)

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
