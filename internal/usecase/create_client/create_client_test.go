package create_client

import (
	"testing"

	"github.com/salesof7/go_walletcore-fc/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	u := NewCreateClientUseCase(m)
	output, err := u.Execute(CreateClientInputDTO{
		Name:  "John",
		Email: "john@mail.com",
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "John", output.Name)
	assert.Equal(t, "john@mail.com", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
