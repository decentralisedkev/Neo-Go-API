package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	blockingester "github.com/decentralisedkev/Neo-Go-API/BlockIngester"
	"github.com/decentralisedkev/Neo-Go-API/database"
	"github.com/decentralisedkev/Neo-Go-API/node"
)

func main() {

	db, _ := database.NewLDBDatabase("dirname", 0, 0)
	table := database.NewTable(db, "block")

	blockIndexBuffer := new(bytes.Buffer)
	binary.Write(blockIndexBuffer, binary.LittleEndian, int64(2499476))

	block, _ := table.Get(blockIndexBuffer.Bytes())
	r := bytes.NewReader(block)
	var blockMetric blockingester.BlockMetric
	binary.Read(r, binary.LittleEndian, &blockMetric)

	fmt.Printf("%+v\n", blockMetric)

	// db.Close()
	// os.RemoveAll(dirname)

	// plan, _ := ioutil.ReadFile("main.json")
	// var data NodeList
	// err := json.Unmarshal(plan, &data)
	// if err != nil {
	// 	fmt.Println("There was an error ", err)
	// }
	// for _, element := range data.Sites {
	// 	block, err := element.GetBlock(2499476)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	err = blockingester.SaveBlockMetrics(block)

	// 	break

	// }

}

// We want to collect as much informaton as we can about a block/transaction
// receiving the transaction, we can store it's type, time, amount?
// do we need to store trans?
type NodeList struct {
	Name     string      `json:"name"`
	PollTime string      `json:"pollTime"`
	Sites    []node.Node `json:"sites"`
}
