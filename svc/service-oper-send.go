package main

import (
	"bytes"
	"context"

	"github.com/go-msvc/pcm"
	gsmmap "github.com/go-msvc/utils/gsm/map"
	"github.com/go-msvc/utils/limiter"
)

//operSend() fails if cannot deliver for any reason
//succeeds only when delivered
func operSend(ctx context.Context, req pcm.SendRequest) error {
	//e.g. limit sender to e.g. {12/day} and {1/60s}
	if limiterA, ok := ctx.Value("limiter_a").(limiter.Limiter); ok {
		if !limiterA.Allow("pcm_from_" + req.ANumber.String()) {
			return pcm.LimitedA
		}
	}

	//e.g. limit sender to recipient e.g. to {1/300s}
	if limiterAB, ok := ctx.Value("limiter_ab").(limiter.Limiter); ok {
		if !limiterAB.Allow("pcm_from_" + req.ANumber.String() + "_to_" + req.BNumber.String()) {
			return pcm.LimitedA2B
		}
	}

	//load profile of sender
	profiler := ctx.Value(ctxProfiler{}).(Profiler)
	aProfile, err := profiler.Get(req.ANumber)
	if err != nil {
		return pcm.FailedGetProfileA
	}

	//if ussd.imsi != profile.imsi then get imsi and status from icap
	//(skipped if USSD does not profile IMSI)
	icap := ctx.Value(ctxIcap{}).(IcapClient)
	if req.AImsi != "" && aProfile.Imsi != req.AImsi {
		icapResponse, err := icap.Get(req.ANumber.String())
		if err != nil {
			return pcm.FailedIcapA
		}
		if icapResponse.Status != "ACTIVE" { //icap.Active {
			return pcm.IcapANotActive
		}
		aProfile.Imsi = icapResponse.Imsi
		if err := profiler.Set(aProfile); err != nil {
			return pcm.FailedSetProfileA
		}
	}

	bProfile, err := profiler.Get(req.BNumber)
	if err != nil {
		return pcm.FailedGetProfileB
	}
	if bProfile.BlockedPcm {
		return pcm.BlockedByB
	}

	gsmMapClient := ctx.Value(ctxGsm{}).(gsmmap.Client)

	//see if B is available, where it is and what network it belongs to
	sriSmResB, err := gsmMapClient.DoSriSm(
		ctx,
		gsmmap.SriSmRequest{
			Msisdn: req.BNumber,
		},
	)
	if err != nil {
		return pcm.FailedSriSmB
	}

	switch sriSmResB.Status {
	case gsmmap.SubscriberStatusUnknown: //todo
	case gsmmap.SubscriberStatusAbsent: //todo
	case gsmmap.SubscriberStatusAvailable: //todo
	default: //todo
	}

	config := ctx.Value("pcm").(serviceConfig)
	//see if allowed to this subscriber's network
	if allowedToBNetwork, ok := config.AllowToImsiNetwork[sriSmResB.ImsiNetwork]; !ok || !allowedToBNetwork {
		return pcm.NotAllowedToBImsiNetwork
	}
	//see if allowed to network where subscriber is currently registered
	if allowedToBNetwork, ok := config.AllowToVlrNetwork[sriSmResB.VlrNetwork]; !ok || !allowedToBNetwork {
		return pcm.NotAllowedToBVlrNetwork
	}

	//todo: get advert for SMS to B
	adverts := ctx.Value("adverts").(AdvertsClient)
	bSmsAdvert, err := adverts.Get("B", "SMS")
	if err != nil {
		bSmsAdvert = Advert{}
	}

	buffer := bytes.NewBuffer(nil)
	if err := config.tmplBSms.Execute(
		buffer,
		map[string]interface{}{
			"ANUMBER": req.ANumber,
			"ANAME":   aProfile.Name,
			"BNUMBER": req.BNumber,
			"AD":      bSmsAdvert.Text,
		}); err != nil {
		return pcm.FailedSmsRenderB
	}

	//deliver the SMS to recipient in the GSM network
	if _ /*fwdSmResA*/, err := gsmMapClient.DoFwdSm(
		ctx,
		gsmmap.FwdSmRequest{
			VlrGt: sriSmResB.VlrGt,
			Imsi:  sriSmResB.Imsi,
			Text:  buffer.String(),
		}); err != nil {
		return pcm.FailedFwdSmB
	}

	//todo: get advert for SMS to A
	aSmsAdvert, err := adverts.Get("A", "SMS")
	if err != nil {
		aSmsAdvert = Advert{}
	}

	//buffer := bytes.NewBuffer(nil)
	buffer.Reset()
	if err := config.tmplBSms.Execute(
		buffer,
		map[string]interface{}{
			"ANUMBER": req.ANumber,
			"ANAME":   aProfile.Name,
			"BNUMBER": req.BNumber,
			"AD":      aSmsAdvert.Text,
		}); err != nil {
		return pcm.FailedSmsRenderA
	}

	//todo: put in SMSC for proper retry because A might be on a call and not be notified...
	//could just submit with smpp into external SMSC and use own SMSC

	//see if B is available, where it is and what network it belongs to
	sriSmResA, err := gsmMapClient.DoSriSm(
		ctx,
		gsmmap.SriSmRequest{
			Msisdn: req.ANumber,
		},
	)
	if err != nil {
		return pcm.FailedSriSmA
	}

	switch sriSmResA.Status {
	case gsmmap.SubscriberStatusUnknown: //todo
	case gsmmap.SubscriberStatusAbsent: //todo
	case gsmmap.SubscriberStatusAvailable: //todo
	default: //todo
	}

	if _, err := gsmMapClient.DoFwdSm(
		ctx,
		gsmmap.FwdSmRequest{
			VlrGt: sriSmResA.VlrGt,
			Imsi:  sriSmResA.Imsi,
			Text:  buffer.String(),
		},
	); err != nil {
		return pcm.FailedFwdSmA
	}

	//todo: audit record
	return pcm.ResultSuccess
} //operSend()
