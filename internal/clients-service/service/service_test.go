package service

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/pedidopago/trainingsvc-clients/protos/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestService(t *testing.T) (*Service, sqlmock.Sqlmock) {
	rdb, mock, err := sqlmock.New()
	require.NoError(t, err)

	db := sqlx.NewDb(rdb, "sqlmock")
	service := &Service{
		db: db,
	}
	return service, mock
}

func TestNewClient(t *testing.T) {
	service, mock := newTestService(t)

	mock.ExpectExec("INSERT INTO clients.*").WillReturnResult(sqlmock.NewResult(0, 1))
	resp, err := service.NewClient(context.Background(), &pb.NewClientRequest{
		Name:     "Test",
		Birthday: time.Now().UnixNano(),
		Score:    0,
	})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetClients(t *testing.T) {
	//FIXME: escrever teste
	t.Fail()
}
