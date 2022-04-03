package request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeruktutut2/backend-mongo-user/exception"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	fmt.Println("result:", result)
	exception.PanicIfError(err)
}
