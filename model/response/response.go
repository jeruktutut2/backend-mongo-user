package response

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func ToResponse(data interface{}, errors interface{}) (response Response) {
	response.Data = data
	response.Error = errors
	return response
}
