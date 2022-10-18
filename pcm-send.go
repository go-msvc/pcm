package pcm

import "github.com/go-msvc/utils/gsm"

type SendRequest struct {
	ANumber gsm.IsdnAddr `json:"anumber"`
	AImsi   gsm.Imsi     `json:"aimsi"`
	BNumber gsm.IsdnAddr `json:"bnumber"`
}

func (req SendRequest) Validate() error {
	return nil
}
