package model

import (
	"context"
	"marketing/consts"
	"math/rand"
)

type mockDAO struct {
	trs map[uint]*Transaction
}

func initMockDAO() *mockDAO {
	return &mockDAO{}
}

func (dao *mockDAO) SetEnv(env consts.Env) DAO {
	return dao
}

func (dao *mockDAO) CreateTx(ctx context.Context, records []*Transaction) error {
	id := uint(rand.Intn(9999) + 10000)
	for _, record := range records {
		record.ID = id
		id++
		dao.trs[record.ID] = record
	}
	return nil
}

func (dao *mockDAO) CancelTx(ctx context.Context, trId string) error {
	// todo...
	return nil
}

func (dao *mockDAO) ConfirmTx(ctx context.Context, trId string) error {
	// todo...
	return nil
}
