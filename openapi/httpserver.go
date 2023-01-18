// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package openapi

import (
	"net/http"
	"ubiqnet/service"

	"github.com/gin-gonic/gin"
)

type HttpServer struct{}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

type HandlerFunc func(c *gin.Context) error

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(c)
	}
}

func (server *HttpServer) InitHttpServer() {

	gin.SetMode("release")
	r := gin.Default()
	r.Use(Cors())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error":  404,
			"result": nil,
		})
	})
	r.GET("/", indexHandler)
	r.GET("/address", walletAddressHandler)
	r.GET("/chequebook/address", contractAddressHandler)
	r.GET("/exportKey", exportKey)
	r.GET("/harvest/status", mineStatusHandler)
	r.GET("/wallet/balance", walletBalanceHandler)
	r.GET("/chequebook/balance", contractBalanceHandler)

	r.POST("/harvest/deposit", mineDepositHandler)
	r.GET("/harvest/withdraw", mineWithDrawHandler)
	r.GET("/harvest/cashout", cashOutHandler)

	r.GET("/harvest/redeem", undepositHandler)
	r.GET("/status", statusHandler)
	r.GET("/upgrade", UpgradeHandler)
	r.GET("/hardInfo", GetHardInfoHandler)
	r.POST("/hardInfo", PostHardInfoHandler)
	r.POST("/nft/mint", MintNft)
	r.POST("nft/put", PutNft)
	r.POST("/nft/down", DeleteNft)
	r.GET("/nft/list", GetNftList)
	r.GET("/nft", GetAllNftList)
	r.GET("/nft/getHardDetail", GetNftDetail)
	r.GET("/delOrder", DelOrder)
	r.GET("/releaseAll", ReleaseAll)
	r.Run(service.Node.DebugApi)
}
