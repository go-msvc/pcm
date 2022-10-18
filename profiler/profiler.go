package profiler

import (
	"github.com/go-msvc/utils/gsm"
	"github.com/go-msvc/utils/keyvalue"
)

type Profiler interface {
	Get(msisdn gsm.IsdnAddr) (SubscriberProfile, error)
}

type profiler struct {
	Store keyvalue.Store
}

func (p profiler) Get(msisdn gsm.IsdnAddr) (SubscriberProfile, error) {
	sub, err := p.Store.Get("pcm_" + msisdn.String())
	if err != nil {
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

func (p profiler) Set(sub SubscriberProfile) error {
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
