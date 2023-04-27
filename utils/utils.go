package utils

import (
	"encoding/json"
	"net/http"
)

func BodyDecoder(w http.ResponseWriter, r *http.Request) *json.Decoder {
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // body의 최대 크기
	dec := json.NewDecoder(r.Body)                   // decoder
	dec.DisallowUnknownFields()
	return dec
}
