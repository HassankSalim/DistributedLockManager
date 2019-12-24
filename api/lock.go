package api

import (
	"github.com/HassankSalim/DistributedLockManager/util"
	"github.com/gorilla/mux"
	"net/http"
)

func (api ApiClient) AcquireLock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, ok := vars["key"]
	if !ok {
		util.WriteResponse(w, "Key not present", http.StatusBadRequest)
		return
	}
	lock := AcquireLockResponse{
		Key:key,
	}
	ttl, token := api.core.AcquireLock(key)
	if token != "" {
		lock.TTL = ttl
		lock.Token = token
		util.WriteResponse(w, lock, http.StatusOK)
		return
	}
	util.WriteResponse(w, "Key already locked", http.StatusForbidden)
	return
}

func (api ApiClient) ReleaseLock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, keyExists := vars["key"]
	token, tokenExists := vars["token"]
	if !keyExists || !tokenExists {
		util.WriteResponse(w, "Key or Token not present", http.StatusBadRequest)
		return
	}
	success := api.core.ReleaseLock(key, token)
	util.WriteResponse(w, success, http.StatusOK)
}