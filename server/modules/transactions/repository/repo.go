package transactionrepo

type TransactionStorage interface{}

type transactionRepo struct {
	storage TransactionStorage
}

func NewTransactionRepo(storage TransactionStorage) *transactionRepo {
	return &transactionRepo{storage: storage}
}
