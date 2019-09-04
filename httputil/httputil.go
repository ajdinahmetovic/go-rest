package httputil

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

//Response struct
type Response struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

//WriteError sends error response
func WriteError(w http.ResponseWriter, message string) {

	res := Response{
		ID:      rand.Intn(1000),
		Message: message,
	}

	v, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(v)

}

//WriteResponse ...
func WriteResponse(w http.ResponseWriter, message string) {

	res := Response{
		ID:      rand.Intn(1000),
		Message: message,
	}

	v, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(v)

}
