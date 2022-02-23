package model

type Graph struct {
	VID  string `json:"vid"`
	Tag  string `json:"tag"`
	Edge string `json:"edge"`
}

type Response struct {
	StatusCode   int         `json:"statusCode"`
	ErrorMessage string      `json:"errorMessage"`
	Result       interface{} `json:"result"`
}

type Request struct {
	Content string
}
