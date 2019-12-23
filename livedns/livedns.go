package livedns

import "github.com/tiramiseb/go-gandi-livedns/client"

type LiveDNS struct {
	client client.Gandi
}

func New(g client.Gandi) *LiveDNS {
	g.SetEndpoint("livedns/")
	return &LiveDNS{client: g}
}
