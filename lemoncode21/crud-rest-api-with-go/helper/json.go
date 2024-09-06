package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(r *http.Request, resul interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(resul)
	PanicIfError(err)
}

func WriteResponseBody(write http.ResponseWriter, response interface{}) {
	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err := encoder.Encode(response)
	PanicIfError(err)
}
