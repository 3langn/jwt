package helper

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

type EmptyObject struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	return Response{
		Status:  false,
		Message: message,
		Data:    data,
		Errors:  err,
	}
}
