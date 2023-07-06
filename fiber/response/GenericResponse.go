package response

type GenericResponse[T comparable] struct {
	Status   string `json:"status"`
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Response *T     `json:"response"`
}

func OK[T comparable](response *T) *GenericResponse[T] {
	return &GenericResponse[T]{
		Status:   "Success",
		Success:  true,
		Message:  "Success",
		Response: response,
	}
}

func OKWithMessage[T comparable](response *T, message string) *GenericResponse[T] {
	return &GenericResponse[T]{
		Status:   "Success",
		Success:  true,
		Message:  message,
		Response: response,
	}
}

func Error[T comparable](response *T) *GenericResponse[T] {
	return &GenericResponse[T]{
		Status:   "Error",
		Success:  false,
		Message:  "Failed",
		Response: response,
	}
}

func ErrorWithMessage[T comparable](response *T, message string) *GenericResponse[T] {
	return &GenericResponse[T]{
		Status:   "Error",
		Success:  false,
		Message:  message,
		Response: response,
	}
}
