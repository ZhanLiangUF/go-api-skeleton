package errdefs // import "github.com/ZhanLiangUF/go-flights/errdefs"

type ErrNotFound interface {
	NotFound()
}

// ErrInvalidParameter signals that the user input is invalid
type ErrInvalidParameter interface {
	InvalidParameter()
}
