package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBlock(t *testing.T) {
	node := Node{
		Protocol: "https",
		URL:      "seed1.switcheo.network",
		Location: "Singapore",
		Locale:   "sg",
		Port:     "10331",
		Type:     "RPC",
	}

	block, err := node.GetBlock(200)
	if err != nil {
		assert.Fail(t, "There was an error getting the block from the node "+err.Error())
	}

	expected := "0xe5313431ecd1b59d1cb7848f35dadc5e51ebd700fd232fe3502003e14aeb9bf7"
	actual := block.Hash
	assert.Equal(t, expected, actual)
}
