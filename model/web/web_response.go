package web

type WebResponse struct {
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}
