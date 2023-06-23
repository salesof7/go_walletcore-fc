package createtransaction

import (
	"github.com/salesof7/go_walletcore-fc/internal/entity"
	"github.com/salesof7/go_walletcore-fc/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (u *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := u.AccountGateway.FindByID(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}
	accountTo, err := u.AccountGateway.FindByID(input.AccountIDTo)
	if err != nil {
		return nil, err
	}
	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}
	err = u.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}
	return &CreateTransactionOutputDTO{ID: transaction.ID}, nil
}
