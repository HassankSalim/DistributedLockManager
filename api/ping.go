package api

import (
	"net/http"
)

func (api ApiClient) Ping(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("pong"))
}
