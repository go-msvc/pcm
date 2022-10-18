package main

//todo - make this a ms with req/res structs and an interface to import and config to use

type advertsConfig struct{}

func (c advertsConfig) Validate() error { return nil }

type AdvertsClient interface {
	Get(party string, media string) (Advert, error)
}

type Advert struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
