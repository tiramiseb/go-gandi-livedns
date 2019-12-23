package livedns

import "github.com/tiramiseb/go-gandi-livedns/client"

type LiveDNS struct {
	client client.Gandi
}

func New(apikey string, sharing_id string, debug bool, dry_run bool) *LiveDNS {
	client := client.New(apikey,  sharing_id,  debug, dry_run)
	client.SetEndpoint("livedns/")
	return &LiveDNS{client: *client}
}

func NewFromClient(g client.Gandi) *LiveDNS {
	g.SetEndpoint("livedns/")
	return &LiveDNS{client: g}
}
