package redis

// import (
// 	"context"
// 	"marketing/util/conf"
// 	"testing"
// 	"time"
// )

// func TestRedisLock(t *testing.T) {
// 	conf.Init()
// 	Init()
// 	c := context.Background()
// 	key := "test"
// 	locked, lockVal, err := Lock(c, key, time.Second)
// 	if err != nil {
// 		t.Errorf("lock failed: %v", err)
// 		return
// 	}
// 	if !locked {
// 		t.Errorf("lock failed, not locked")
// 		return
// 	}
// 	if lockVal == "" {
// 		t.Errorf("lockVal is empty")
// 		return
// 	}
// 	// unlock key not exist
// 	err = Unlock(c, "key not exist", false, lockVal)
// 	if err != nil {
// 		t.Errorf("unlock failed: %v", err)
// 		return
// 	}
// 	// unlock not locked
// 	err = Unlock(c, key, false, lockVal)
// 	if err != nil {
// 		t.Errorf("unlock failed: %v", err)
// 		return
// 	}
// 	// unlock other lockval
// 	err = Unlock(c, key, true, "other lockval")
// 	if err == nil {
// 		t.Errorf("unlock others lock!")
// 		return
// 	}
// 	// unlock success
// 	err = Unlock(c, key, locked, lockVal)
// 	if err != nil {
// 		t.Errorf("unlock failed: %v", err)
// 		return
// 	}
// 	// duplicate unlock
// 	err = Unlock(c, key, locked, lockVal)
// 	if err != nil {
// 		t.Errorf("duplicate unlock failed: %v", err)
// 		return
// 	}
// 	// unlock expired lock
// 	time.Sleep(time.Second)
// 	err = Unlock(c, key, locked, lockVal)
// 	if err != nil {
// 		t.Errorf("unlock failed: %v", err)
// 		return
// 	}
// }
