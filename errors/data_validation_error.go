package errors

type DataValidationError struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Value int    `json:"value,omitempty"`
}
