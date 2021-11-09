package bcwallet

import (
	"context"
	"fmt"
	"log"
	"nolan/spin-game/components/appctx"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func Deposit(appctx appctx.AppContext) {
	client := appctx.GetEthClient()

	smAddress := appctx.GetSmartContractAddr()

	contractAddress := common.HexToAddress(smAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		// TODO: handle event from contract, maybe deposit, or somthing
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case vLog := <-logs:
				fmt.Println(vLog.Address) // pointer to event log
				// var d []interface{}
				// if err := json.Unmarshal(vLog.Data, &d); err != nil {
				// 	panic(err)
				// }
				fmt.Printf("%+v\n", vLog)
				fmt.Printf("%+v\n", string(vLog.Data))
			}
		}
	}()
}
