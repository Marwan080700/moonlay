package result

type SuccessResult struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}