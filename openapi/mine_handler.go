// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package openapi

import (
	"fmt"
	"math/big"
	"strconv"
	"ubiqnet/constants"
	"ubiqnet/service"
	"ubiqnet/store/localstore"

	"github.com/gin-gonic/gin"
)

var rpcService service.RpcService

type RechargeRequest struct {
	Amount *big.Int `json:"amount"`
}

func mineStatusHandler(c *gin.Context) {
	contractAddress := localstore.Get(constants.ContractAddressKey)
	res, err := rpcService.GetAllInfo(contractAddress)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		// "harvest": service.Node.MinerStatus,
		"reward":  res.AllReturnAmount,
		"pending": res.RewardAmount,
		"pledge":  res.DepositAmount,
	})
}

func mineDepositHandler(c *gin.Context) {
	var request RechargeRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}

	contractAddress := localstore.Get(constants.ContractAddressKey)

	balance := rpcService.BalanceOfUbq()
	if balance.Cmp(big.NewInt(0)) < 0 {
		c.JSON(200, gin.H{
			"error":  "ubq insufficient balance",
			"result": nil,
		})
		return
	}

	t, err := rpcService.Recharge(request.Amount, contractAddress)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"error":  nil,
		"result": t.Hash().String(),
	})

}

func mineWithDrawHandler(c *gin.Context) {

	contractAddress := localstore.Get(constants.ContractAddressKey)
	max, err := rpcService.BalanceOfMiner(contractAddress)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}
	incomeData, err := rpcService.GetAllInfo(contractAddress)
	if err != nil {
		c.JSON(500, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
	}
	pledgeNum := incomeData.DepositAmount
	temp := big.Int{}
	num := temp.Sub(max, pledgeNum)
	if num.Cmp(big.NewInt(0)) <= 0 {
		c.JSON(200, gin.H{
			"error":  "Insufficient cash balance",
			"result": nil,
		})
		return
	}

	t, err := rpcService.WithDraw(contractAddress, num)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"error":  nil,
		"result": t.Hash().String(),
	})

}

func cashOutHandler(c *gin.Context) {
	contractAddress := localstore.Get(constants.ContractAddressKey)
	amount := rpcService.GetRewardAmount(contractAddress)
	if amount == nil || amount.Cmp(big.NewInt(0)) == 0 {
		c.JSON(200, gin.H{
			"error":  "No dividends",
			"result": nil,
		})
		return
	}

	gasAmount, err := rpcService.GetGasAmount()
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}

	t, err := rpcService.GetReward(amount, contractAddress, gasAmount)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"error":  nil,
		"result": t.Hash().String(),
	})

}

func undepositHandler(c *gin.Context) {
	contractAddress := localstore.Get(constants.ContractAddressKey)
	incomeData, err := rpcService.GetAllInfo(contractAddress)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}
	if incomeData.DepositAmount.Cmp(big.NewInt(0)) == 0 {
		c.JSON(400, gin.H{
			"error":  "Please pledge first",
			"result": nil,
		})
		return
	}
	t, err := rpcService.UnDeposit(contractAddress)
	if err != nil {
		c.JSON(200, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"error":  nil,
		"result": t.Hash().String(),
	})

}

func UpgradeHandler(c *gin.Context) {

	nodeTypeStr := c.DefaultQuery("nodeType", "")
	if nodeTypeStr == "" {
		c.JSON(200, gin.H{
			"result": nil,
			"error":  "node Type is none",
		})
		return
	}
	nodeType, err := strconv.Atoi(nodeTypeStr)

	if err != nil {
		c.JSON(200, gin.H{
			"result": nil,
			"error":  err,
		})
		return
	}

	contractAddress := localstore.Get(constants.ContractAddressKey)
	ubqAmount, err := rpcService.BalanceOfMiner(contractAddress)
	if err != nil {
		c.JSON(200, gin.H{
			"result": nil,
			"error":  err,
		})
	}

	if nodeType == 2 && (ubqAmount.Cmp(big.NewInt(0)) <= 0) {
		c.JSON(200, gin.H{
			"result": nil,
			"error":  "Stable node ubq pledge amount needs to be greater than 0",
		})
		return
	}

	verifyPleage := big.NewInt(constants.PleageVerify)
	verifyPleage = verifyPleage.Mul(verifyPleage, big.NewInt(constants.PledgeSymbol))
	if nodeType == 3 && (ubqAmount.Cmp(verifyPleage) < 0) {
		c.JSON(200, gin.H{
			"result": nil,
			"error":  fmt.Sprintf("Verify nodes need to pledge at least %d ubq", constants.PleageVerify),
		})
		return
	}

	var pleageAmount int64
	if nodeType == 3 {
		pleageAmount = constants.PleageVerify

	} else if nodeType == 2 {
		pleageAmount = constants.PleageStable
	} else if nodeType == 1 {
		pleageAmount = constants.PleageStable
	}
	x := big.NewInt(pleageAmount)
	x = x.Mul(x, big.NewInt(constants.PledgeSymbol))
	realPleageAmount := big.NewInt(0).Sub(x, ubqAmount)
	if realPleageAmount.Cmp(big.NewInt(0)) > 0 {
		t, err := rpcService.Recharge(realPleageAmount, contractAddress)
		if err != nil {
			c.JSON(200, gin.H{
				"result": err,
				"error":  nil,
			})
			return
		}
		service.Node.NodeType = nodeType
		localstore.SetNodeType(nodeTypeStr)
		c.JSON(200, gin.H{
			"result": t.Hash(),
			"error":  nil,
		})
		return

	}
	if nodeType == 3 {
		service.Node.NodeType = constants.NodeTypeVerify
		localstore.SetNodeType(strconv.Itoa(constants.NodeTypeVerify))
	} else if nodeType == 2 {
		service.Node.NodeType = constants.NodeTypeStable
		localstore.SetNodeType(strconv.Itoa(constants.NodeTypeStable))
	}

	c.JSON(200, gin.H{
		"result": "Success",
		"error":  nil,
	})
}
