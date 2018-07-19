package blockingester

//Prefix = block
type BlockMetric struct {
	BlockSize          int64
	HeaderSize         int64
	NumOfTrans         int64
	BlockTime          int64
	Sysfee             int64
	Netfee             float64
	NextConsensusSig   [34]byte
	Transactions       [9]int64 // Transactions[0]=number of miner transactions,Transactions[1]= # Of IssueTransactions
	AvgTransactionSize int64
	Atrributes         [10]int64 // Attributes[0]=# of Contract Hash in block
}

//Prefix = "Transactions" key = block_transID
type TransactionMetric struct {
	TransSize int64
}

const (
	MinerTransaction      = "MinerTransaction"
	IssueTransaction      = "IssueTransaction"
	ClaimTransaction      = "ClaimTransaction"
	EnrollmentTransaction = "EnrollmentTransaction"
	RegisterTransaction   = "RegisterTransaction"
	ContractTransaction   = "ContractTransaction"
	PublishTransaction    = "PublishTransaction"
	InvocationTransaction = "InvocationTransaction"
	UnknownTransaction    = "UnknownTransaction"
)

// 0x00	MinerTransaction	0	Assign byte fees
// 0x01	IssueTransaction	500|0	Inssuance of asset
// 0x02	ClaimTransaction	0	Assign GAS
// 0x20	EnrollmentTransaction	1000	Enrollment for validator
// 0x40	RegisterTransaction	10000	Assets register
// 0x80	ContractTransaction	0	Contract transaction
// 0xd0	PublishTransaction	500 * n	(Not usable) Special Transactions for Smart Contracts
// 0xd1	InvocationTransaction	0	Special transactions for calling Smart Contracts

const (
	ContractHash   = "ContractHash"
	ECDH           = "ECDH"
	Script         = "Script"
	Vote           = "Vote"
	CertUrl        = "CertUrl"
	DescriptionUrl = "DescriptionUrl"
	Description    = "Description"
	Hash           = "Hash"
	Remark         = "Remark"
	Unknown        = "Unknown"
)

// 0x00	ContractHash	Hash value of contract
// 0x02-0x03	ECDH02-ECDH03	Public key for ECDH key exchange
// 0x20	Script	Additional validation of transactions
// 0x30	Vote	For voting
// 0x80	CertUrl	Url address of certificate
// 0x81	DescriptionUrl	Url address of description
// 0x90	Description	Brief description
// 0xa1-0xaf	Hash1-Hash15	Used to store custom hash values
// 0xf0-0xff	Remark-Remark15	Remarks
