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

func IsNotFound(err error) bool {
	_, ok := getImplementer(err).(ErrNotFound)
	return ok
}

func IsInvalidParameter(err error) bool {
	_, ok := getImplementer(err).(ErrInvalidParameter)
	return ok
}
