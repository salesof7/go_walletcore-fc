package create_transaction

import (
	"context"
	"testing"

	"github.com/salesof7/go_walletcore-fc/internal/entity"
	"github.com/salesof7/go_walletcore-fc/internal/event"
	"github.com/salesof7/go_walletcore-fc/internal/usecase/mocks"
	"github.com/salesof7/go_walletcore-fc/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("client1", "client1@mail.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("client2", "client2@mail.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()
	ctx := context.Background()

	u := NewCreateTransactionUseCase(mockUow, dispatcher, event)
	output, err := u.Execute(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
