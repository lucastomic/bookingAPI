package exceptions

// ApiError is an error instance with API information, as an status code and a specific message
type ApiError interface {
	// ApiError returns the status code of the exceptio as first value and the
	// exception's message as second one
	APIError() (int, string)
	// Error returns the exception's message
	Error() string
}
