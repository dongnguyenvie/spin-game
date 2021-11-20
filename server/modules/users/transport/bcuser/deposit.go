package bcuser

import (
	"context"
	"fmt"
	"log"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	contracts "nolan/spin-game/components/constracts"
	"nolan/spin-game/components/pubsub"
	userstorage "nolan/spin-game/modules/users/storage"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	goecommon "github.com/ethereum/go-ethereum/common"
)

func Deposit(appctx appctx.AppContext) {
	client := appctx.GetEthClient()

	smAddress := appctx.GetSmartContractAddr()
	pb := appctx.GetPubsub()

	contractAddress := goecommon.HexToAddress(smAddress)
	spinContract, err := contracts.NewSpincontract(contractAddress, client)
	if err != nil {
		panic(err)
	}

	db := appctx.GetMaiDBConnection()
	userStore := userstorage.NewSQLStore(db)
	// walletRepo := walletrepo.NewWalletRepo(walletStore)

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
					return
				}
				result, error := userStore.FindDataWithCondition(context.Background(), map[string]interface{}{"wallet_address": vLog.Sender.String()})

				if error != nil {
					panic("wallet address not found")
				}

				message := pubsub.NewMessage(common.UserDeposit{
					Tx:     tx,
					UserId: result.Id,
					Amount: vLog.Amount,
				})
				pb.Publish(context.Background(), common.ChannelDepositBC, message)

				fmt.Println("tx", tx)
				fmt.Println("wallet_address", vLog.Sender) // pointer to event log
				fmt.Println("amount", vLog.Amount)         // pointer to event log
			}
		}
	}()
}
