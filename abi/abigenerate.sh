#!/bin/sh
abigen --abi ERC20.abi --pkg abi --type Ubq --out ubq.go
abigen --abi ERC4907DEMO.abi --pkg abi --type Nft --out nft.go
abigen --abi MinerContract.abi --pkg abi --type Miner --out miner.go
abigen --abi MinerFactory.abi --pkg abi --type MinerFactory --out factory.go
abigen --abi MyApp.abi --pkg abi --type App --out app.go
abigen --abi MyTemplate.abi --pkg abi --type Template --out template.go
abigen --abi Trade.abi --pkg abi --type Trade --out trade.go
abigen --abi MyStack.abi --pkg abi --type Pool --out pool.go