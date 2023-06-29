package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/salesof7/go_walletcore-fc/internal/database"
	"github.com/salesof7/go_walletcore-fc/internal/event"
	"github.com/salesof7/go_walletcore-fc/internal/usecase/create_account"
	"github.com/salesof7/go_walletcore-fc/internal/usecase/create_client"
	"github.com/salesof7/go_walletcore-fc/internal/usecase/create_transaction"
	"github.com/salesof7/go_walletcore-fc/internal/web"
	"github.com/salesof7/go_walletcore-fc/internal/web/webserver"
	"github.com/salesof7/go_walletcore-fc/pkg/events"
	"github.com/salesof7/go_walletcore-fc/pkg/uow"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)
	webserver.Start()
}
