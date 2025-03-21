package dto

type TokenResponse struct {
	Token string `json:"token"`
}

type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}
