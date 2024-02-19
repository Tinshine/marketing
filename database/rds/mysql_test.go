package rds

import (
	"context"
	"marketing/consts"
	"marketing/util/conf"
	"marketing/util/log"
	"testing"
)

func TestInit(t *testing.T) {
	conf.Init()
	log.Init()
	Init()

	c := context.Background()
	pdb := ProdDB(c)
	if pdb == nil {
		t.Errorf("ProdDB: prod db not initialized")
		return
	}
	tdb := TestDB(c)
	if tdb == nil {
		t.Errorf("TestDB: test db not initialized")
		return
	}
	ddb := DB(c, consts.Dev)
	if ddb != nil {
		t.Errorf("DB: ddb should not initialized")
		return
	}
	tdb = DB(c, consts.Test)
	if tdb == nil {
		t.Errorf("DB: test db not initialized")
		return
	}
	pdb = DB(c, consts.Prod)
	if pdb == nil {
		t.Errorf("DB: prod db not initialized")
		return
	}
}
