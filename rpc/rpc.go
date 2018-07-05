package rpc

import (
	"fmt"

	"github.com/ybbus/jsonrpc"
)

type Rpc struct {
	Url string
}

func (u *Rpc) GetVersion() (VersionResults, error) {
	rpcClient := jsonrpc.NewClient(u.Url)
	response, err := rpcClient.Call("getversion", "")
	var versionRes VersionResults
	response.GetObject(&versionRes)
	if err != nil {
		return versionRes, err
	}

	return versionRes, nil
}
func (u *Rpc) GetBlockCount() (int64, error) {
	rpcClient := jsonrpc.NewClient(u.Url)
	response, err := rpcClient.Call("getblockcount", "")

	if err != nil {
		return response.GetInt()
	}

	return response.GetInt()
}

type Host struct {
	Address string
	Port    int
}

func (u *Rpc) GetPeers() (map[string][]Host, error) {
	rpcClient := jsonrpc.NewClient(u.Url)

	response, err := rpcClient.Call("getpeers", "")

	if err != nil || response == nil {
		return nil, nil
	}
	// fmt.Println(response)
	var res map[string][]Host
	err = response.GetObject(&res)
	return res, err
}

func (u *Rpc) GetRawMempool() []string {
	rpcClient := jsonrpc.NewClient(u.Url)

	response, err := rpcClient.Call("getrawmempool", "")
	res := []string{}
	if err != nil || response == nil {
		return res
	}

	transactions := response.Result.([]interface{})

	for _, transaction := range transactions {
		if transactionAsString, ok := transaction.(string); ok {
			res = append(res, transactionAsString)
		}

	}

	return res
}

func (u *Rpc) GetBlock(index int) (BlockRes, error) {
	rpcClient := jsonrpc.NewClient(u.Url)
	response, err := rpcClient.Call("getblock", index, 1)
	fmt.Println(response)
	if err != nil {

		return BlockRes{}, err
	}

	var res *BlockRes

	err = response.GetObject(&res) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	if err != nil || res == nil {
		// some error on json unmarshal level or json result field was null

		return BlockRes{}, err
	}
	return *res, nil

}
