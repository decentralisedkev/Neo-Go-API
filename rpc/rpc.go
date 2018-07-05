package rpc

import (
	"fmt"

	"github.com/decentralisedkev/Neo-Go-API/models"
	"github.com/ybbus/jsonrpc"
)

type Rpc struct {
	Url string
}

func GetVersion(url string) (*models.VersionResults, error) {
	rpcClient := jsonrpc.NewClient(url)
	response, err := rpcClient.Call("getversion", "")
	var versionRes models.VersionResults
	response.GetObject(&versionRes)
	if err != nil {
		return nil, err
	}

	return &versionRes, nil
}
func GetBlockCount(url string) (int64, error) {
	rpcClient := jsonrpc.NewClient(url)
	response, err := rpcClient.Call("getblockcount", "")

	if err != nil {
		return 0, err
	}

	return response.GetInt()
}

func GetPeers(url string) (map[string][]models.Host, error) {
	rpcClient := jsonrpc.NewClient(url)

	response, err := rpcClient.Call("getpeers", "")

	var res map[string][]models.Host
	err = response.GetObject(&res)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetRawMempool(url string) ([]string, error) {
	rpcClient := jsonrpc.NewClient(url)

	response, err := rpcClient.Call("getrawmempool", "")
	res := []string{}
	if err != nil {
		return res, err
	}

	transactions := response.Result.([]interface{})

	for _, transaction := range transactions {
		if transactionAsString, ok := transaction.(string); ok {
			res = append(res, transactionAsString)
		}

	}

	return res, nil
}

func (u *Rpc) GetBlock(index int) (models.BlockRes, error) {
	rpcClient := jsonrpc.NewClient(u.Url)
	response, err := rpcClient.Call("getblock", index, 1)
	fmt.Println(response)
	if err != nil {

		return models.BlockRes{}, err
	}

	var res *models.BlockRes

	err = response.GetObject(&res) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	if err != nil || res == nil {
		// some error on json unmarshal level or json result field was null

		return models.BlockRes{}, err
	}
	return *res, nil

}
