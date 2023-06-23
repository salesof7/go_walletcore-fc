package gateway

import "github.com/salesof7/go_walletcore-fc/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
