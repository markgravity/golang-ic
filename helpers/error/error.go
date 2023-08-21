package errorhelpers

import (
	"net/http"
	"os"
	"strings"
)

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(strings.ToUpper(err.Error()), NotFound)
}

func IsTimeout(err error) bool {
	if err == nil {
		return false
	}

	return os.IsTimeout(err)
}

func GetErrorCode(err error) string {
	if IsNotFound(err) {
		return NotFoundCode
	}

	if IsTimeout(err) {
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
