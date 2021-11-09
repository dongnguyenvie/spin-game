package bcwallet

import (
	"context"
	"fmt"
	"log"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	contracts "nolan/spin-game/components/constracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	goecommon "github.com/ethereum/go-ethereum/common"
)

func Deposit(appctx appctx.AppContext) {
	client := appctx.GetEthClient()

	smAddress := appctx.GetSmartContractAddr()

	contractAddress := goecommon.HexToAddress(smAddress)
	spinContract, err := contracts.NewSpincontract(contractAddress, client)
	if err != nil {
		panic(err)
	}

	go func() {
		defer common.AppRecover()

		logs := make(chan *contracts.SpincontractDeposit)
		sub, err := spinContract.WatchDeposit(&bind.WatchOpts{Context: context.Background()}, logs)
		if err != nil {
			log.Fatal(err)
		}

		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case vLog := <-logs:
				fmt.Printf("%v\n", vLog.Raw)
				tx := vLog.Raw.TxHash.String()

				if vLog.Raw.Removed {
					// TODO: recovery balance when the tsx is failed
				}
				// TODO: handle update balance by wallet address
				fmt.Println("tx", tx)
				fmt.Println("wallet_address", vLog.Sender) // pointer to event log
				fmt.Println("amount", vLog.Amount)         // pointer to event log
			}
		}
	}()
}
