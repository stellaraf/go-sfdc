package types

type ServerErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type QueryErrorResponse struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}
