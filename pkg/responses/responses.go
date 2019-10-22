package responses

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, code int, message string) {
	Response(w, code, map[string]string{"error": message})
}

func Response(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
