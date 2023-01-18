// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package openapi

import (
	"ubiqnet/service"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

func walletAddressHandler(c *gin.Context) {
	data := WalletData{
		WalletAddress: service.Node.WalletAddress,
	}
	c.JSON(200, gin.H{
		"error":  nil,
		"result": &data,
	})
}
func exportKey(c *gin.Context) {
	privateKey := service.Node.PrivateKey
	privateKeyBytes := crypto.FromECDSA(privateKey)

	priv := common.Bytes2Hex(privateKeyBytes)
	c.JSON(200, gin.H{
		"address":    service.Node.WalletAddress,
		"privateKey": priv,
	})
}

func walletBalanceHandler(c *gin.Context) {

	eth, _ := rpcService.BalanceOfEth()

	ubq := rpcService.BalanceOfUbq()
	if service.Node.NetWorkId == 10 {
		c.JSON(200, gin.H{
			"xdai": eth,
			"ubq":  ubq,
		})

		return
	}
	c.JSON(200, gin.H{
		"eth": eth,
		"ubq": ubq,
	})

}
