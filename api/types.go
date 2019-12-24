package api

import "github.com/HassankSalim/DistributedLockManager/core"

type ApiClient struct {
	core core.Core
}

type AcquireLockResponse struct {
	Key string
	TTL int
	Token string
}

func NewClient(core core.Core) ApiClient {
	return ApiClient{
		core: core,
	}
}
