package exceptions

import "net/http"

var (
	NotFound            = apiErrorInfo{http.StatusNotFound, "not found"}
	WrongIdType         = apiErrorInfo{http.StatusBadRequest, "parameter id must be a integer"}
	ReservationCollides = apiErrorInfo{http.StatusConflict, "unable to set the new reservation. There is not enough space"}
	WrongEmailLogin     = apiErrorInfo{http.StatusBadRequest, "wrong email at login"}
	WrongPasswordLogin  = apiErrorInfo{http.StatusBadRequest, "wrong password at login"}
)
