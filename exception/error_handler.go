package exception

import (
	"fmt"
	"net/http"

	"github.com/jeruktutut2/backend-mongo-user/helper"
	"github.com/jeruktutut2/backend-mongo-user/model/web"
)

func ErrorHandler(writer http.ResponseWriter, err error) {
	if badRequestError(writer, err) {
		return
	} else if internalServerError(writer, err) {
		return
	}
}

func badRequestError(writer http.ResponseWriter, err error) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Data:   nil,
			Errors: exception.Message,
		}
		errWriteToResponseBody := helper.WriteToResponseBody(writer, webResponse)
		if errWriteToResponseBody != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusInternalServerError)
			webResponse = web.WebResponse{
				Data:   nil,
				Errors: "Internal Server Error",
			}
			errWriteToResponseBody = helper.WriteToResponseBody(writer, webResponse)
			fmt.Println(errWriteToResponseBody)
		}
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, err error) bool {
	exception, ok := err.(InternalServerError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)

		webResponse := web.WebResponse{
			Data:   nil,
			Errors: exception.Message,
		}
		errWriteToResponseBody := helper.WriteToResponseBody(writer, webResponse)
		if errWriteToResponseBody != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusInternalServerError)
			webResponse = web.WebResponse{
				Data:   nil,
				Errors: "Internal Server Error",
			}
			errWriteToResponseBody = helper.WriteToResponseBody(writer, webResponse)
			fmt.Println(errWriteToResponseBody)
		}
		return true
	} else {
		return false
	}
}
