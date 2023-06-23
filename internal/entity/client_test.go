package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John", "john@mail.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John", client.Name)
	assert.Equal(t, "john@mail.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John", "john@mail.com")
	err := client.Update("John updated", "john@mail.com")
	assert.Nil(t, err)
	assert.Equal(t, "John updated", client.Name)
	assert.Equal(t, "john@mail.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John", "john@mail.com")
	err := client.Update("", "john@mail.com")
	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John", "john@mail.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
