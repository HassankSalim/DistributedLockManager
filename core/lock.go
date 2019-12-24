package core

import (
	"github.com/HassankSalim/DistributedLockManager/backendstore"
	"github.com/google/uuid"
	"time"
)

const DefaultTTl  = 360

type Core struct {
	dbs []backendstore.Store
}

func (c Core) AcquireLock(key string) (int, string) {
	ch := make(chan bool)
	var token uuid.UUID
	var err error
	if token, err = uuid.NewUUID(); err != nil {
		return 0, ""
	}
	defer close(ch)
	startTime := time.Now().Unix()
	for _, db := range c.dbs {
		go db.SetIfNotExists(key, token.String(), DefaultTTl, ch)
	}
	if c.hasQuorum(&ch) {
		return int(time.Now().Unix() - startTime), token.String()
	}
	return 0, ""
}

func (c Core) ReleaseLock(key, token string) bool {
	ch := make(chan struct{})
	defer close(ch)
	for _, db := range c.dbs {
		go db.DelIfKeyHasVal(key, token, ch)
	}
	for range c.dbs {
		<-ch
	}
	return true
}

func (c Core) hasQuorum(ch *chan bool) bool {
	majority := len(c.dbs) / 2 + 1
	successCount := 0
	for range c.dbs {
		if <-*ch {
			successCount++
		}
	}
	return successCount >= majority
}

func NewCore(dbs ...backendstore.Store) Core {
	return Core{
		dbs:dbs,
	}
}