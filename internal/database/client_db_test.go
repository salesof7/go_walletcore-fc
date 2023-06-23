package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/salesof7/go_walletcore-fc/internal/entity"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), email varchar(255), name varchar(255), created_at date)")
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	client, _ := entity.NewClient("John", "john@mail.com")
	err := s.clientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("John", "john@mail.com")
	s.clientDB.Save(client)

	clientDB, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}
