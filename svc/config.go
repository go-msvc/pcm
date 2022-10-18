package main

import (
	"text/template"

	"github.com/go-msvc/errors"
	"github.com/go-msvc/utils/gsm"
)

type serviceConfig struct {
	AllowToImsiNetwork map[*gsm.Network]bool
	AllowToVlrNetwork  map[*gsm.Network]bool

	ASmsTemplate string `json:"a_sms_template" doc:"Template for SMS sent to sender"`
	tmplASms     *template.Template
	BSmsTemplate string `json:"b_sms_template" doc:"Template for SMS sent to recipient"`
	tmplBSms     *template.Template
}

func (c *serviceConfig) Validate() error {
	var err error
	c.tmplASms, err = template.New("a_sms").Parse(c.ASmsTemplate)
	if err != nil {
		return errors.Wrapf(err, "invalid template a_sms_template")
	}
	c.tmplBSms, err = template.New("b_sms").Parse(c.BSmsTemplate)
	if err != nil {
		return errors.Wrapf(err, "invalid template b_sms_template")
	}
	return nil
}

func (c serviceConfig) Loaded() {}
