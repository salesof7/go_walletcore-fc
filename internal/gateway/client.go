package gateway

import "github.com/salesof7/go_walletcore-fc/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
