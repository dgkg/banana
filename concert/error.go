package concert

type ErrConcertResponse struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func (e ErrConcertResponse) Error() string {
	return e.Message
}

func NewErrNotFound(message string) ErrConcertResponse {
	return ErrConcertResponse{
		Code:      404,
		Status:    "Not Found",
		Message:   message,
		Timestamp: "now",
	}
}
