package database

const (
	Nodes        uint8 = 0x01 //key=prefix+nodeaddrPort
	Blocks       uint8 = 0x02 //key = prefix+blockHash + H || another key also would be prefix + blockNum + suffix
	Address      uint8 = 0x03 // key = prefix + address
	Token        uint8 = 0x04 // key = prefix+TokenHash
	Transactions uint8 = 0x05 // key
	// This will have data on a specific block, such as time it took to reach consensus.
	// The amount of transactions in it, including the different types of transactions inside of it
	//
)

// average block size 24 hours period (Approx 2880 blocks) - memory
// tx's per day 24 hour period (Approx 2880 blocks) - memory

// time it took for a transaction --

// When a block comes in, we look at how many blocks in memory, if >2.88k we delete last known and recalculate
