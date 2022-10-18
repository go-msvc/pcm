package main

import (
	_ "github.com/go-msvc/rest-utils/server"
	gsmmap "github.com/go-msvc/utils/gsm/map"
	"github.com/go-msvc/utils/limiter"
	"github.com/go-msvc/utils/ms"
)

type ctxPcmService struct{}
type ctxPrmService struct{}
type ctxLimiterA struct{}
type ctxLimiterAB struct{}
type ctxAdverts struct{}
type ctxProfiler struct{}
type ctxIcap struct{}
type ctxGsm struct{}

func main() {
	ms := ms.New("pcm",
		//all named config are accessible to handler from ctx
		//all configs with Create() methods will also be called during ms.Configure()
		//and the output will also be accessible from ms.Configured(ctx,name)
		ms.WithConfig("pcm", ctxPcmService{}, &serviceConfig{}), //specific to sending pcm
		ms.WithConfig("prm", ctxPrmService{}, &serviceConfig{}), //specific to sending prm
		ms.WithConfig("limiter_a", ctxLimiterA{}, limiter.Config{}),
		ms.WithConfig("limiter_ab", ctxLimiterAB{}, limiter.Config{}),
		ms.WithConfig("adverts", ctxAdverts{}, advertsConfig{}),
		ms.WithConfig("profiler", ctxProfiler{}, profilerConfig{}),
		ms.WithConfig("icap", ctxIcap{}, icapConfig{}),
		// ms.WithConfig("sms_sender", sms_sender.Config),
		ms.WithConfig("gsm", ctxGsm{}, gsmmap.Config{}),
		ms.WithOper("send", operSend),
		//ms.WithOper("get_profile", operGetProfile),
		//ms.WithOper("upd_profile", operUpdProfile),
	)
	ms.Configure()
	ms.Serve()
}
