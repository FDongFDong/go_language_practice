package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"reflect"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BlockListner() error {
	client, err := ethclient.Dial("WSSRPC")
	// 에러 시
	if err != nil {
		log.Fatal(err)
	}
	// 블록 이벤트 받아오기
	// SubscribeNewHead : 헤더 정보가 새로 생겼을 때 받아오기
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("##############")
			fmt.Println("Block Hash : ", block.Hash().Hex())
			fmt.Println("Block Number : ", block.Number().Uint64())
			fmt.Println("Block Timestamp : ", block.Time())
			fmt.Println("Block Nonce : ", block.Nonce())
			fmt.Println("Total Transaction : ", len(block.Transactions()))
			fmt.Println("##############")
			baseFee := block.BaseFee()

			if len(block.Uncles()) > 0 {
				// 하나의 블록에 최대 두개의 uncle까지 들어갈 수 있기 때문에 2개의 uncle을 처리해주기 위함
				for _, uncle := range block.Uncles() {
					// uncle 수수료 계산
					uncleFee := float64((uncle.Number.Uint64()+8-block.Number().Uint64())*2) / 8.0
					fmt.Println("Uncle Block Length : ", len(block.Uncles()))
					// 블록을 채굴한 사람과 uncle 블록을 채굴한 사람이 다른 경우가 많다.
					fmt.Println("Uncle Miner Address : ", uncle.Coinbase.Hex())
					fmt.Println("Uncle Block Number :", uncle.Number.Uint64())
					fmt.Println("Uncle Block Reward :", uncleFee)
				}
			}
			// 블록안에 트랜잭션 모두 가져오기
			for _, tx := range block.Transactions() {
				fmt.Println("*************")
				fmt.Println("Transaction Hash ", tx.Hash().Hex())
				if tx.To() != nil {
					fmt.Println("To Address : ", tx.To())
				} else {
					fmt.Println("To Address : Contract Creation")
					//TODO
					contractAddress := GetContractAddress(client, tx.Hash())
					fmt.Println("To Address : ", contractAddress)
				}
				fmt.Println("Transfer Value(wei) : " + tx.Value().String())
				fmt.Println("Transaction nonce : ", tx.Nonce())
				// 예상 Gas Limit
				fmt.Println("Transaction Gas Limit : ", tx.Gas())
				// real Gas Limit
				realGasLimit := GetRealGasUsed(client, tx.Hash())
				fmt.Println("Transation Real Gas Limit : ", realGasLimit)
				fmt.Println("Transaction GasFeeCap : ", tx.GasFeeCap().Uint64())
				fmt.Println("Transaction GasTipCap : ", tx.GasTipCap().Uint64())
				// Gas Price
				realGasPrice := GetRealGasPrice(baseFee.Uint64(), tx.GasFeeCap().Uint64(), tx.GasTipCap().Uint64())
				fmt.Println("Transcation Real Gas Price : ", realGasPrice)
				fmt.Println("Transaction Input Data : ", hex.EncodeToString(tx.Data()))

				// ERC20 데이터 가져오기
				if len(tx.Data()) != 0 {

					to, value := ERC20Transaction(hex.EncodeToString(tx.Data()))

					if to != "" {
						// value 값은 ERC20 토큰마다 사용하는 Decimal 값이 다르므로 다르게 처리해줘야한다.
						symbol, name, decimal := GetContractInfo(client, tx.To())
						fmt.Println("ERC20 Transfer To Address : ", to)
						fmt.Println("ERC20 Contarct Name : ", name)
						fmt.Println("ERC20 Contarct symbol : ", symbol)
						fmt.Println("ERC20 Contarct decimal : ", decimal)
						fmt.Println("ERC20 Transfer To Address : ", to)
						fmt.Println("ERC20 Transfer Value : ", value) // value / (10 ** decimal)

					}
				}
			}
		}
	}

}
