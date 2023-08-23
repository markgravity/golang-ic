package errorhelpers

import (
	"net/http"
	"os"
	"strings"
)

func GetErrorCode(err error) string {
	if isNotFound(err) {
		return NotFoundCode
	}

	if isTimeout(err) {
		return TimeoutCode
	}

	return InternalServerErrorCode
}

func GetErrorStatusCode(err error) int {
	switch GetErrorCode(err) {
	case NotFoundCode:
		return http.StatusNotFound
	case TimeoutCode:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

func isNotFound(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(strings.ToUpper(err.Error()), NotFound)
}

func isTimeout(err error) bool {
	if err == nil {
		return false
	}

	return os.IsTimeout(err)
}
