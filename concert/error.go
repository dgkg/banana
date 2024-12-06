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
