package main

import (
	"github.com/go-msvc/errors"
	"github.com/go-msvc/utils/gsm"
	"github.com/go-msvc/utils/keyvalue"
)

type profilerConfig struct {
	StoreConfig keyvalue.Config `json:"store"`
}

func (c profilerConfig) Validate() error {
	if c.StoreConfig == nil {
		return errors.Errorf("missing store")
	}
	if err := c.StoreConfig.Validate(); err != nil {
		return errors.Wrapf(err, "invalid store")
	}
	return nil
}

func (c profilerConfig) Create() (any, error) {
	if c.StoreConfig == nil {
		return nil, errors.Errorf("store not configured")
	}
	kv, err := c.StoreConfig.Create()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create store")
	}
	return Profiler{
		Store: kv,
	}, nil
}

type Profiler struct {
	Store keyvalue.Store
}

func (p Profiler) Get(msisdn gsm.IsdnAddr) (SubscriberProfile, error) {
	sub, err := p.Store.Get("pcm_" + msisdn.String())
	if err != nil || sub == nil {
		sub = SubscriberProfile{
			Imsi:           "",
			Name:           "",
			BlockedPcm:     false,
			BlockedAdverts: false,
			MmsCapable:     false,
			BlockedMms:     false,
		} //default values
	}
	return sub.(SubscriberProfile), nil
}

func (p Profiler) Set(sub SubscriberProfile) error {
	return p.Store.Set("pcm_"+sub.Msisdn.String(), sub)
}

type SubscriberProfile struct {
	Msisdn         gsm.IsdnAddr `json:"msisdn"`
	Imsi           gsm.Imsi     `json:"imsi"`
	Name           string       `json:"name"`
	BlockedPcm     bool         `json:"blocked_pcm"`
	BlockedAdverts bool         `json:"blocked_adverts"`
	MmsCapable     bool         `json:"mms_capable"` //imported from MMS system data
	BlockedMms     bool         `json:"blocked_mms"`
}
