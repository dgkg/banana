package handler

type ErrorHandlerValidation struct {
	Err        error
	Message    string
	Entity     string
	StatusCode int
}

func (e *ErrorHandlerValidation) Error() string {
	return e.Message
}

func NewErrorValidation(entity, message string, errOrigin error) *ErrorHandlerValidation {
	return &ErrorHandlerValidation{
		Err:        errOrigin,
		Message:    message,
		Entity:     entity,
		StatusCode: 400,
	}
}

func NewErrorAutorization(entity, uuidUser string) *ErrorHandlerValidation {
	return &ErrorHandlerValidation{
		Err:        nil,
		Message:    "user not authorized",
		Entity:     entity + " " + uuidUser,
		StatusCode: 401,
	}
}
