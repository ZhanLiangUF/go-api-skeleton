package api // import "github.com/zhanlianguf/go-api-skeleton/api"

import (
	"github.com/ZhanLiangUF/go-api-skeleton/api/httputils"
)

// handlerWithGlobalMiddlewares wraps handler function with server's global middleware
func (s *Server) handlerWithGlobalMiddlewares(handler httputils.APIFunc) httputils.APIFunc {
	m := s.middleware.WrapHandler(handler)
	return m
}
