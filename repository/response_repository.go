package repository

// Response Repository
type Response struct {
	Status  int
	Message string
	Data    interface{}
}

// Responses DATA
func Responses(status int, message string, data interface{}) Response {
	var response Response
	response.Status = status
	response.Message = message
	response.Data = data
	return response
}
