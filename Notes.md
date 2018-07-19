## Notes


- Currently it is not possible to get the P2P status of all nodes in the network.

Example of two handshake protocols/ race: 

telnet 18.218.255.178 20333 <- protocol 1

telnet 206.189.152.158 10332 <- protocol2


This problem will make it seem as if the majority of nodes have p2p status offline, I have been informed that this is due to the race condition in the c cde. python does not accept incoming.