// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package openapi

import (
	"math/big"
	"ubiqnet/constants"
	"ubiqnet/service"
	"ubiqnet/store/localstore"
	"ubiqnet/tools"

	version "ubiqnet"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
)

// @Summary Welcom
// @Accept  json
// @Produce json
// @Success 200 string  "Welcome to nctr"
// @Router / [get]
func indexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Welcome to Ubiqnet",
	})
}

func statusHandler(c *gin.Context) {
	contractAddress := localstore.Get(constants.ContractAddressKey)
	walletAddress := service.Node.WalletAddress

	info, err := rpcService.GetAllInfo(contractAddress)
	if err != nil {

		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}
	amount, err := rpcService.BalanceOfMiner(contractAddress)
	if err != nil {

		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}

	temp := big.NewInt(0)
	withdraw := temp.Sub(amount, info.DepositAmount)
	if withdraw.Cmp(big.NewInt(0)) <= 0 {
		withdraw = big.NewInt(0)
	}
	data := StatusData{
		Status:          1,
		WalletAddress:   walletAddress,
		NodeType:        service.Node.NodeType,
		Version:         version.CurrentVersionNumber,
		ContractAddress: contractAddress,
		PledgeBalance:   info.DepositAmount,
		TotalBalance:    amount,
		RewardBalance:   withdraw,
		PublicKey:       hexutil.Encode(service.Node.PublicKey),
	}
	c.JSON(200, gin.H{
		"result": data,
		"error":  nil,
	})

}

func PostHardInfoHandler(c *gin.Context) {
	tools.DeviceInfo.Status = 0
	go tools.GetHardwareData()
	c.JSON(200, gin.H{
		"result": nil,
		"error":  nil,
	})

}

func GetHardInfoHandler(c *gin.Context) {

	deviceInfo := tools.DeviceInfo
	c.JSON(200, gin.H{
		"result": deviceInfo,
		"error":  nil,
	})
}
