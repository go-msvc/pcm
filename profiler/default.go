package profiler

import (
	"github.com/go-msvc/errors"
	"github.com/go-msvc/utils/keyvalue"
)

type inMemConstructor struct {
	Store keyvalue.Config `json:"store"`
}

func (c inMemConstructor) Validate() error {
	if c.Store == nil {
		return errors.Errorf("missing store")
	}
	if err := c.Store.Validate(); err != nil {
		return errors.Wrapf(err, "invalid store")
	}
	return nil
}

func (c inMemConstructor) Create() (Profiler, error) {
	if c.Store == nil {
		return nil, errors.Errorf("store not configured")
	}
	kv, err := c.Store.Create()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create store")
	}
	return profiler{
		Store: kv,
	}, nil
}
