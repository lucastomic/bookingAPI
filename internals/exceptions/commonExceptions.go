package exceptions

import "net/http"

var (
	NotFound    = apiErrorInfo{http.StatusNotFound, "not found"}
	WrongIdType = apiErrorInfo{http.StatusBadRequest, "parameter id must be a integer"}
)
