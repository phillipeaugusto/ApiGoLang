package utils

import (
	"ApiGoLang/core/results"
	"encoding/json"
	"net/http"
)

func ResultRequest(w http.ResponseWriter, statusCode int, result results.Result) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(result)
}
