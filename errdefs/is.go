package errdefs // import "github.com/ZhanLiangUF/go-flights/errdefs"

type causer interface {
	Cause() error
}

func getImplementer(err error) error {
	switch e := err.(type) {
	case
		ErrNotFound,
		ErrInvalidParameter:
		return err
	case causer:
		return getImplementer(e.Cause())
	default:
		return err
	}
}

// IsNotFound returns if the error passed in is an ErrNotFound
func IsNotFound(err error) bool {
	_, ok := getImplementer(err).(ErrNotFound)
	return ok
}

// IsInvalidParameter returns if error passed in is ErrInvalidParameter
func IsInvalidParameter(err error) bool {
	_, ok := getImplementer(err).(ErrInvalidParameter)
	return ok
}
