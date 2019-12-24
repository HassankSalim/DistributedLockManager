package main

import (
	"github.com/HassankSalim/DistributedLockManager/api"
	"github.com/HassankSalim/DistributedLockManager/backendstore"
	"github.com/HassankSalim/DistributedLockManager/core"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func addRoutes(api api.ApiClient) *mux.Router {
	handler := mux.NewRouter()
	handler.HandleFunc("/api/ping", api.Ping)
	handler.HandleFunc("/api/v1/lock/{key}", api.AcquireLock).Methods("GET")
	handler.HandleFunc("/api/v1/lock/{key}/{token}", api.ReleaseLock).Methods("DELETE")
	return handler
}

func main() {
	redisAddrOne := "172.17.0.2:6379"
	redisAddrTwo := "172.17.0.3:6379"
	redisDb := 0
	rdOne := backendstore.NewRedisStore(redisAddrOne, redisDb)
	rdTwo := backendstore.NewRedisStore(redisAddrTwo, redisDb)
	core := core.NewCore([]backendstore.Store{rdOne, rdTwo}...)
	api := api.NewClient(core)
	err := http.ListenAndServe(":8080", addRoutes(api))
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
