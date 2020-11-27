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

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// CHS Foundation Go Bootnodes

	"enode://45edec81a0143394fb3d3857a51365878992acd12f7f14e1425f0bf0e6377fc44e3113bec818beb2e62d624fa91d567d0f28d669b82832795572a7ff082a428b@49.235.119.123:30303",
	"enode://575ec6f9f9e1a05b5fc55cbb4bae4998ed38f055e7db8c7bffa4fc9fad840e5b841f80afda70dff6d594d27b762c24b173d0f4b7fa0765ed811cafea15d9d047@101.32.204.15:10104",

	"enode://8ce21b55e1c20128db873b67794cbbeb679fcf437bd3596c27b168e81c70e612643551ff125c4af59e8e3c789fc9cea14997518a02b5181f04d95f822442b9d5@129.28.15.222:10101",

	"enode://8b8d72c2cb864790c825446f9c7fe79d2134fbdea65c76d84c5bf21182d704b45ed896a75241ac5f2a114ff6b1eeeea894d7e204ae24d8fde5cb47e948287821@112.17.171.62:10101",
	"enode://8b8d72c2cb864790c825446f9c7fe79d2134fbdea65c76d84c5bf21182d704b45ed896a75241ac5f2a114ff6b1eeeea894d7e204ae24d8fde5cb47e948287821@[2409:8728:4ea:23:e4a1:abff:fe3c:ac48]:10101",

	"enode://d9506147ea7db2f494bc7a6d24bb4c3371d7cc0acb89abb5b29919838e5d542ee87f5478e182ef53e1763e987e30e0aea7891570f4d6dc5b4c4650bf50772bde@112.17.171.62:10102",
	"enode://d9506147ea7db2f494bc7a6d24bb4c3371d7cc0acb89abb5b29919838e5d542ee87f5478e182ef53e1763e987e30e0aea7891570f4d6dc5b4c4650bf50772bde@[2409:8728:4ea:23:a033:3dff:feab:f6b5]:10102",

	"enode://41992cf850fda21584aa39427b75473ffa8c78c266243ed8ba7517fffcc879994393e62caf26aa6283988aba6611514abb96be380897ed659cf48d9443e808e9@112.17.171.62:10103",
	"enode://41992cf850fda21584aa39427b75473ffa8c78c266243ed8ba7517fffcc879994393e62caf26aa6283988aba6611514abb96be380897ed659cf48d9443e808e9@[2409:8728:4ea:23:dc9d:dff:fee2:ea00]:10103",

	"enode://289948b84e6b7bceb6e85ffe779563174660151abbf449c8309b3fbf96202e093c176355b2b7450366c8078d323615fec8830f56f1dd5eb2f25124320b9d1f7b@112.17.171.62:10104",
	"enode://289948b84e6b7bceb6e85ffe779563174660151abbf449c8309b3fbf96202e093c176355b2b7450366c8078d323615fec8830f56f1dd5eb2f25124320b9d1f7b@[2409:8728:4ea:23:8474:d7ff:fef2:687a]:10104",

	"enode://7f44aa20a629f43661e260b4653b623f0c32620d02b1321a6762c09707930a4ec22ec7529859bd4dfc1317e55bae6635a2a71fbc4ead0df3ec2e4a4625eac3a2@gchs.p2sp.net:30303",
	"enode://b2a13b6fb4b969354b264de1ba7e9d8c054b0c2330b4efb4df7024e615c070f782cb88568f4fb45f7783920b2826be0b761501bcc802d608f2a138df0c5c3b4c@gchs.microcai.org:30303",

	"enode://575ec6f9f9e1a05b5fc55cbb4bae4998ed38f055e7db8c7bffa4fc9fad840e5b841f80afda70dff6d594d27b762c24b173d0f4b7fa0765ed811cafea15d9d047@gchs01.superpool.io:10104",
	"enode://06f4536f3956646c49e32d82ac31721669012324c4da28210795865d2cd821ffc88737c442c5458788e0ef9252df4c549226d69092df0d035e2d5abdac7723e5@gchs02.superpool.io:10102",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://45edec81a0143394fb3d3857a51365878992acd12f7f14e1425f0bf0e6377fc44e3113bec818beb2e62d624fa91d567d0f28d669b82832795572a7ff082a428b@49.235.119.123:30303",
	"enode://575ec6f9f9e1a05b5fc55cbb4bae4998ed38f055e7db8c7bffa4fc9fad840e5b841f80afda70dff6d594d27b762c24b173d0f4b7fa0765ed811cafea15d9d047@101.32.204.15:10104",

	"enode://8ce21b55e1c20128db873b67794cbbeb679fcf437bd3596c27b168e81c70e612643551ff125c4af59e8e3c789fc9cea14997518a02b5181f04d95f822442b9d5@129.28.15.222:10101",

	"enode://8b8d72c2cb864790c825446f9c7fe79d2134fbdea65c76d84c5bf21182d704b45ed896a75241ac5f2a114ff6b1eeeea894d7e204ae24d8fde5cb47e948287821@112.17.171.62:10101",
	"enode://8b8d72c2cb864790c825446f9c7fe79d2134fbdea65c76d84c5bf21182d704b45ed896a75241ac5f2a114ff6b1eeeea894d7e204ae24d8fde5cb47e948287821@[2409:8728:4ea:23:e4a1:abff:fe3c:ac48]:10101",

	"enode://d9506147ea7db2f494bc7a6d24bb4c3371d7cc0acb89abb5b29919838e5d542ee87f5478e182ef53e1763e987e30e0aea7891570f4d6dc5b4c4650bf50772bde@112.17.171.62:10102",
	"enode://d9506147ea7db2f494bc7a6d24bb4c3371d7cc0acb89abb5b29919838e5d542ee87f5478e182ef53e1763e987e30e0aea7891570f4d6dc5b4c4650bf50772bde@[2409:8728:4ea:23:a033:3dff:feab:f6b5]:10102",

	"enode://41992cf850fda21584aa39427b75473ffa8c78c266243ed8ba7517fffcc879994393e62caf26aa6283988aba6611514abb96be380897ed659cf48d9443e808e9@112.17.171.62:10103",
	"enode://41992cf850fda21584aa39427b75473ffa8c78c266243ed8ba7517fffcc879994393e62caf26aa6283988aba6611514abb96be380897ed659cf48d9443e808e9@[2409:8728:4ea:23:dc9d:dff:fee2:ea00]:10103",

	"enode://289948b84e6b7bceb6e85ffe779563174660151abbf449c8309b3fbf96202e093c176355b2b7450366c8078d323615fec8830f56f1dd5eb2f25124320b9d1f7b@112.17.171.62:10104",
	"enode://289948b84e6b7bceb6e85ffe779563174660151abbf449c8309b3fbf96202e093c176355b2b7450366c8078d323615fec8830f56f1dd5eb2f25124320b9d1f7b@[2409:8728:4ea:23:8474:d7ff:fef2:687a]:10104",
}

// YoloV1Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// YOLOv1 ephemeral test network.
var YoloV1Bootnodes = []string{
	"enode://9e1096aa59862a6f164994cb5cb16f5124d6c992cdbf4535ff7dea43ea1512afe5448dca9df1b7ab0726129603f1a3336b631e4d7a1a44c94daddd03241587f9@35.178.210.161:30303",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".nodes.superpool.io"
}
