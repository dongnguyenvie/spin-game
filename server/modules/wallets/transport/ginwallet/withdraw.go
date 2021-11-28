package ginwallet

import (
	"context"
	"math/big"
	"net/http"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/tokenprovider"
	walletbiz "nolan/spin-game/modules/wallets/biz"
	walletmodel "nolan/spin-game/modules/wallets/model"
	walletstorage "nolan/spin-game/modules/wallets/storage"

	goecommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

func Withdraw(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(*tokenprovider.TokenPayload)
		var data walletmodel.Withdraw
		err := c.ShouldBind(&data)
		if err != nil {
			panic(err)
		}

		db := appctx.GetMaiDBConnection()
		walletStore := walletstorage.NewSQLStore(db)
		biz := walletbiz.NewWithdraw(walletStore)
		err = biz.Withdraw(c, user, &data)

		if err != nil {
			panic(err)
		}

		privateKeyStr := appctx.GetAccountPrivateKey()
		privateKey, err := crypto.HexToECDSA(privateKeyStr)
		if err != nil {
			panic(err)
		}
		publicKey := privateKey.PublicKey
		fromAddress := crypto.PubkeyToAddress(publicKey)
		client := appctx.GetEthClient()
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			panic(err)
		}
		gasLimit := uint64(21000)
		gasPrice := big.NewInt(30000000000)
		toAddress := goecommon.HexToAddress(user.WalletAddress)
		chainID, err := client.NetworkID(context.Background())

		if err != nil {
			panic(err)
		}

		value := big.NewInt(data.Amount)

		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			panic(err)
		}
		data.Tx = signedTx.Hash().Hex()

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			panic(err)
		}

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
