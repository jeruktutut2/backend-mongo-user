package response

import (
	"encoding/json"
	"net/http"

	"github.com/jeruktutut2/backend-mongo-user/exception"
)

type UserLoginResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func ToUserLoginResponse(id string, username string) (userLoginResponse UserLoginResponse) {
	userLoginResponse.Id = id
	userLoginResponse.Username = username
	return userLoginResponse
}

func ResponseHttp(w http.ResponseWriter, httpStatusCode int, data interface{}, errors interface{}) {
	response := ToResponse(data, errors)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	err := json.NewEncoder(w).Encode(response)
	exception.PanicIfError(err)
}
