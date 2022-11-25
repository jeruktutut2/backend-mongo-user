package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) (err error) {
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) (err error) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(response)
	if err != nil {
		return err
	}
	return nil
}
