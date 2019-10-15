package httputil

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

//Response struct
type Response struct {
	ID      int         `json:"id"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//WriteError sends error response
func WriteError(w http.ResponseWriter, err error, status int) {
	res := Response{
		ID:      rand.Intn(1000),
		Message: err.Error(),
	}
	v, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(v)
}

//WriteResponse ...
func WriteResponse(w http.ResponseWriter, message string, data interface{}) {
	res := Response{
		ID:      rand.Intn(1000),
		Message: message,
		Data:    data,
	}
	v, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(v)
}

//EnableCors func
func EnableCors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
