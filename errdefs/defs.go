package errdefs // import "github.com/ZhanLiangUF/go-flights/errdefs"

// implment these errors to pass into httputils.MakeErrorHandler
type ErrNotFound interface {
	NotFound()
}

// ErrInvalidParameter signals that the user input is invalid
type ErrInvalidParameter interface {
	InvalidParameter()
}
