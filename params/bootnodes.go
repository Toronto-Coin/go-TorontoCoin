// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{

	// Tcoin
	"enode://40703c1469b0b4ede61753d66f88d1e7a45849b71a18a07207bf9fa02ee7f7677b6dae5edb2fc50eefdc23c2e80ddf4791af6aa86c273dd971b20fa045fe7133@70.49.88.93:30301", 
	"enode://2c07d8755690f92c6bc1f4c60cd1fbb12543e02e7d270ea3148350c375c254617482bacc66a73862d4ca58b6c54b05c3804b9f99a582d5acf32f5db9d8e2dd57@54.245.184.155:30301",
	
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestnetBootnodes = []string{

}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	
}
