package errdefs // import "github.com/ZhanLiangUF/go-flights/errdefs"
import "net/http"

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
	}
	return statusCode
}
