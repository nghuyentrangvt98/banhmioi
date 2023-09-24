package model

type Exception struct {
	Message string `json:"message,omitempty"`
}

func (e Exception) Error() string {
	return e.Message
}
