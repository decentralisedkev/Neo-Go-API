package node

import (
	"fmt"
	"strconv"

	"github.com/decentralisedkev/Neo-Go-API/models"
	"github.com/decentralisedkev/Neo-Go-API/rpc"
	"github.com/decentralisedkev/Neo-Go-API/utils/slice"
)

type Node struct {
	Protocol string
	Url      string
	Location string
	Type     string
	Locale   string
	Address  string
	Port     uint16
}

func (n *Node) String() string {
	if n.Type == "RPC" {
		return n.Address + ":" + strconv.Itoa(int(n.Port))
	} else {
		return "API Methods not implemented"
	}
}

func (n *Node) GetVersion() (*models.VersionResults, error) {
	if n.Type == "RPC" {

		return rpc.GetVersion(n.String())
	} else {
		// rest methods not implemented
		return nil, fmt.Errorf("API methods not implemented")

	}
}
func (n *Node) GetBlockCount() (int64, error) {
	if n.Type == "RPC" {
		return rpc.GetBlockCount(n.String())
	} else {
		// rest methods not implemented
		return 0, fmt.Errorf("API methods not implemented")

	}
}
func (n *Node) GetRawMempool() ([]string, error) {
	if n.Type == "RPC" {
		return rpc.GetRawMempool(n.String())
	} else {
		// rest methods not implemented
		return []string{}, fmt.Errorf("API methods not implemented")

	}
}
func (n *Node) GetRawMempoolCount() (int, error) {
	mempool, err := n.GetRawMempool()
	if err != nil {
		return 0, err
	}
	return len(mempool), nil
}
func (n *Node) GetPeers() ([]models.Host, error) {
	if n.Type == "RPC" {

		peers, err := rpc.GetPeers(n.String())

		if err != nil {
			return nil, err
		}

		duplicatesRemoved := sliceutils.RemoveDuplicates(peers["connected"])
		//The rpc call to getpeercount returns the number of peers including duplicates.
		return duplicatesRemoved, nil

	} else {
		// rest methods not implemented
		return nil, fmt.Errorf("API methods not implemented")

	}

}

func (n *Node) GetPeersCount() (int, error) {
	peers, err := n.GetPeers()
	if err != nil {
		return 0, err
	}
	return len(peers), nil
}

func (n *Node) GetLatency() (int64, error) {
	if n.Type == "RPC" {
		return rpc.GetLatency(n.String())
	} else {
		return 0, fmt.Errorf("API methods not implemented")
	}

}
