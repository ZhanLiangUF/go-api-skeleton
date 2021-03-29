package errdefs //// import "github.com/ZhanLiangUF/go-api-skeleton/errdefs"

type errNotFound struct{ error }

// NotFound is a helper to create an error of the class with the same name from any error type
func NotFound(err error) error {
	if err == nil || IsNotFound(err) {
		return err
	}
	return errNotFound{err}
}
