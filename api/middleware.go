package api // import "github.com/zhanlianguf/go-flights/api"

import (
	"github.com/ZhanLiangUF/go-flights/api/httputils"
)

func (s *Server) handlerWithGlobalMiddleWare(handler httputils.APIFunc) httputils.APIFunc {
	m := s.middleware.WrapHandler(handler)
	return m
}
