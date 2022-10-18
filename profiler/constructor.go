package profiler

import "github.com/go-msvc/utils/ms"

type Constructor interface {
	ms.Config
	Create() (any, error)
}
