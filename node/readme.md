# Package - Node


This documentation serves to explain any outliers in the code or unusual behaviour.

### GetPeers Method

There is an rpc call to return the number of peers connected to a node. However this also includes the duplicates. The solution, was to get all connected peers and remove any duplicates before counting.