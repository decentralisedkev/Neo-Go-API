package blockingester

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"

	"github.com/decentralisedkev/Neo-Go-API/database"
	"github.com/decentralisedkev/Neo-Go-API/models"
	"github.com/decentralisedkev/Neo-Go-API/node"
)

func getBlock(node node.Node, blockNumber int) *models.BlockRes {
	block, err := node.GetBlock(blockNumber)
	if err != nil {
		return nil
	}
	return &block
}

// saveBlockAndTime

func SaveBlockMetrics(block models.BlockRes) error {

	db, _ := database.NewLDBDatabase("dirname", 0, 0)
	table := database.NewTable(db, "block")

	blockmetric := BlockMetric{}
	blockmetric.BlockTime = int64(block.Time)

	blockmetric.BlockSize = int64(block.Size)
	blockHeaderSize := 4 + 32 + 32 + 4 + 4 + 8 + 20 + 1 + len(block.Script.Invocation) + len(block.Script.Verification) + 1
	blockmetric.HeaderSize = int64(blockHeaderSize)
	copy(blockmetric.NextConsensusSig[:], block.Nextconsensus)
	blockmetric.NumOfTrans = int64(len(block.Tx))

	blockmetric.Sysfee = 0
	blockmetric.Netfee = 0
	blockmetric.Transactions = ProcessTransactionType(block)
	blockmetric.Atrributes = processAttributes(block)
	blockmetric.AvgTransactionSize = processAverageTransactionSize(block)
	for _, tx := range block.Tx {
		SysFeeAsInt, _ := strconv.Atoi(tx.SysFee)
		NetFeeAsInt, _ := strconv.ParseFloat(tx.NetFee, 64)
		blockmetric.Sysfee = blockmetric.Sysfee + int64(SysFeeAsInt)
		blockmetric.Netfee = blockmetric.Netfee + NetFeeAsInt
	}

	blockMetricBuffer := new(bytes.Buffer)
	err := binary.Write(blockMetricBuffer, binary.LittleEndian, &blockmetric)

	blockIndexBuffer := new(bytes.Buffer)
	err = binary.Write(blockIndexBuffer, binary.LittleEndian, int64(block.Index))

	err = table.Put(blockIndexBuffer.Bytes(), blockMetricBuffer.Bytes())

	return err
}

func ProcessTransactionType(block models.BlockRes) [9]int64 {

	totalForSpecificTransaction := [9]int64{}

	for _, tx := range block.Tx {
		switch tx.Type {
		case MinerTransaction:
			totalForSpecificTransaction[0]++
		case IssueTransaction:
			totalForSpecificTransaction[1]++
		case ClaimTransaction:
			totalForSpecificTransaction[2]++
		case EnrollmentTransaction:
			totalForSpecificTransaction[3]++
		case RegisterTransaction:
			totalForSpecificTransaction[4]++
		case ContractTransaction:
			totalForSpecificTransaction[5]++
		case PublishTransaction:
			totalForSpecificTransaction[6]++
		case InvocationTransaction:
			totalForSpecificTransaction[7]++
		default:
			totalForSpecificTransaction[8]++
		}
	}
	return totalForSpecificTransaction
}

func processAttributes(block models.BlockRes) [10]int64 {

	totalForSpecificAttribute := [10]int64{}

	for _, tx := range block.Tx {

		// var usage = tx.Attributes[0]["usage"]
		if len(tx.Attributes) > 0 {
			for _, attr := range tx.Attributes {
				var usage = attr["usage"]
				if strings.HasPrefix(usage, ContractHash) {
					totalForSpecificAttribute[0]++
				} else if strings.HasPrefix(usage, ECDH) {
					totalForSpecificAttribute[1]++

				} else if strings.HasPrefix(usage, Script) {

					totalForSpecificAttribute[2]++
				} else if strings.HasPrefix(usage, Vote) {
					totalForSpecificAttribute[3]++

				} else if strings.HasPrefix(usage, CertUrl) {
					totalForSpecificAttribute[4]++

				} else if strings.HasPrefix(usage, DescriptionUrl) {
					totalForSpecificAttribute[5]++

				} else if strings.HasPrefix(usage, Description) {
					totalForSpecificAttribute[6]++

				} else if strings.HasPrefix(usage, Hash) {
					totalForSpecificAttribute[7]++

				} else if strings.HasPrefix(usage, Remark) {
					totalForSpecificAttribute[8]++

				} else {

					totalForSpecificAttribute[9]++

				}

			}
		}
	}
	return totalForSpecificAttribute
}

func processAverageTransactionSize(block models.BlockRes) int64 {
	totalSize := 0

	for _, tx := range block.Tx {
		totalSize += tx.Size
	}
	return int64(totalSize) / int64(len(block.Tx)) // always a Miner transaction, denominator never zero
}
