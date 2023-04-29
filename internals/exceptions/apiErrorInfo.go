package exceptions

// apiErrorInfo is the information needed to describe an API exception
// This is, status code and exception message
type apiErrorInfo struct {
	status int
	msg    string
}

func (e apiErrorInfo) Error() string {
	return e.msg
}

func (e apiErrorInfo) APIError() (int, string) {
	return e.status, e.msg
}

func NewApiError(status int, msg string) error {
	return apiErrorInfo{status: status, msg: msg}
}
