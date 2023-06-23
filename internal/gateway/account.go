package gateway

import "github.com/salesof7/go_walletcore-fc/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
