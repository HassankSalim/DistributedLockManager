package util

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	response := make(map[string]interface{})
	w.WriteHeader(statusCode)
	if is200Response(statusCode) {
		response["success"] = true
		response["data"] = data
	} else {
		response["success"] = false
		response["err"] = data
	}
	res, _ := json.Marshal(response)
	w.Write(res)
}

func is200Response(statusCode int) bool {
	return 100 < statusCode && statusCode < 300
}