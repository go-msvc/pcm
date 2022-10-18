package main

import (
	"github.com/go-msvc/utils/gsm"
)

type icapConfig struct{}

func (c icapConfig) Validate() error { return nil }

func (c icapConfig) Create() (any, error) {
	return icapClient{}, nil
}

type IcapClient interface {
	Get(msisdn string) (IcapResponse, error)
}

type icapClient struct{}

func (icapClient icapClient) Get(msisdn string) (IcapResponse, error) {
	return IcapResponse{}, nil
}

type IcapResponse struct {
	Status string
	Imsi   gsm.Imsi
}
