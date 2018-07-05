# Neo-Go-API


## Structure

Node - When the main.json file is loaded, each node will be unmarshalled into a Node struct. The node folder contains the node struct and all the necessary rpc and api methods that can be called on the node struct.

rpc - This folder contains all of the rpc functions that the node struct will call.

models - This contains all of structs that will be outputted by the rpc functions. It is put into a seperate package because, when the API package is implemented, it will also use this package.