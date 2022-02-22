package model

type Response struct {
	StatusCode   int         `json:"statusCode"`
	ErrorMessage string      `json:"errorMessage"`
	Result       interface{} `json:"result"`
}

type Request struct {
	Content string
}