# Neo-Go-API


## Scope

While reviewing the possible scopes and surveying the community, a list was made as to what features were needed.
This project will aim to provide all of the necessary features that have been requested. If you require a new feature be added please submit an issue with the reason as to why you believe this feature is needed in order to monitor the network.

This project will provide the user with the raw data, an api will be made so that the data can be ingested as a JSON, however reading from the datastore would be faster.

## Requirements

An API primarily written in Golang that collects information on the nodes and the state of the network. Although API servers such as neoscan will not be included in the scope of this project, this project allows the implementation of these.

## Requested Features

Below you will find a list of requested features gathered from surveying the community.

[ x ] Node Latency

[ x ] GetBlock, GetPeers, GetRawMempool, GetVersion

[ - ] Node Stability

[ - ] Average Block Production Times (Mem)

[ - ] Average Trans size (Lifetime)

[ - ] Average Attribute size (Lifetime)

[ - ]  Attribute Count (Lifetime)

[ - ] Transaction count / Transactions per block(General) 

[ - ] Consensus Sigs Per Block

[ - ] Header + Block Size

[ - ] Sys and network fee distribution

[ - ] Total and Avg Neo and Gas per block

[ ] Transactions Per Second (Mem)(2.8k)

[ - ] TX sizes over time

[ ] Value of Out's per Block

[ ] Number of Vins Per Block

[ ] Addresses -> Trans count (This would require a script parses for network transactions)

## Scoped Features

[ - ] Avg Block Size

[ ] Size of UTXO Per Block?

[ ] Mempool size


## Infrastructure decisions 

Leveldb will be used, so that the program can be ran without a cloud hosting provider. The alternative would have been dynamodb, however since they are both key-value stores, the interface will not change much, if you decide to switch to dynamodb.

## Setup

An import module will be provided which will alllow you to decode the chain.acc file, so that users do not have to sync with the other nodes from block zero.

## Referenced APIS

- https://www.blockchain.com/charts

## Structure

Node - When the main.json file is loaded, each node will be unmarshalled into a Node struct. The node folder contains the node struct and all the necessary rpc and api methods that can be called on the node struct.

rpc - This folder contains all of the rpc functions that the node struct will call.

models - This contains all of structs that will be outputted by the rpc functions. It is put into a seperate package because, when the API package is implemented, it will also use this package.


