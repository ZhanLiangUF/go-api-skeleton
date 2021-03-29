package errdefs // import "github.com/ZhanLiangUF/go-api-skeleton/errdefs"
import "net/http"

// GetHTTPErrorStatusCode returns status code from error.
func GetHTTPErrorStatusCode(err error) int {
	if err == nil {
		// log internal server error
	}

	var statusCode int

	switch {
	case IsNotFound(err):
		statusCode = http.StatusNotFound
	case IsInvalidParameter(err):
		statusCode = http.StatusBadRequest
	case IsConflict(err):
		statusCode = http.StatusConflict
	case IsUnauthorized(err):
		statusCode = http.StatusUnauthorized
	case IsUnavailable(err):
		statusCode = http.StatusServiceUnavailable
	case IsForbidden(err):
		statusCode = http.StatusForbidden
	case IsNotModified(err):
		statusCode = http.StatusNotModified
	case IsNotImplemented(err):
		statusCode = http.StatusNotImplemented
	case IsSystem(err) || IsUnknown(err) || IsDataLoss(err) || IsDeadline(err) || IsCancelled(err):
		statusCode = http.StatusInternalServerError
	}

	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	return statusCode
}
