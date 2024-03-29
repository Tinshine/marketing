package redis

import (
	"context"
	"fmt"
	"marketing/consts/errs"
	"math/rand"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func Lock(ctx context.Context, key string, expir time.Duration) (locked bool, lockVal string, err error) {
	lockVal = strconv.FormatInt(rand.Int63(), 10)
	locked, err = rdb.SetNX(ctx, key, lockVal, expir).Result()
	if err != nil {
		return false, "", errors.WithMessage(errs.RedisException, err.Error())
	}
	return locked, lockVal, nil
}

func Unlock(ctx context.Context, key string, locked bool, lockVal string) error {
	if !locked {
		return nil
	}
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return errors.WithMessage(errs.RedisException, fmt.Sprintf("rdb.Get, %v", err))
	}
	if val != lockVal {
		return errors.WithMessage(errs.UnlockOthers, "lockVal not match")
	}
	_, err = rdb.Del(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return errors.WithMessage(errs.RedisException, fmt.Sprintf("rdb.Del, %v", err))
	}
	return nil
}
