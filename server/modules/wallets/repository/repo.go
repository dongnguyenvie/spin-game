package walletrepo

type WalletStorage interface{}

type walletRepo struct {
	storage WalletStorage
}

func NewWalletRepo(storage WalletStorage) *walletRepo {
	return &walletRepo{storage: storage}
}
