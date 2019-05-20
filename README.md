
Simple blockchain implementation written in Go (Go version 1.10).

Dependency: bolt DB - https://github.com/boltdb/bolt

To build, run `go build src/*`

This should create an executable named `blockchain` in the base directory of this repo.

Example commands:

`./blockchain printChain` - print all blocks on the blockchain

`./blockchain addBlock -data {block_data}` - add a block to the blockchain


Reference: https://jeiwan.cc/posts/building-blockchain-in-go-part-1/

