package domain

import (
	"time"

	"github.com/tiramiseb/go-gandi-livedns"
	"github.com/tiramiseb/go-gandi-livedns/client"
)

type Domain struct {
	client client.Gandi
}

func New(apikey string, config *gandi.Config) *Domain {
	client := client.New(apikey,  config)
	client.SetEndpoint("domain/")
	return &Domain{client: *client}
}

func NewFromClient(g client.Gandi) *Domain {
	g.SetEndpoint("domain/")
	return &Domain{client: g}
}

type Contact struct {
	Country string `json:"country"`
	Email string `json:"email"`
	Family string `json:"family"`
	Given string `json:"given"`
	StreetAddr string `json:"streetaddr"`
	ContactType int `json:"type"`
	BrandNumber string `json:"brand_number,omitempty"`
	City string `json:"city,omitempty"`
	DataObfuscated *bool `json:"data_obfuscated,omitempty"`
	Fax string `json:"fax,omitempty"`
	Language string `json:"lang,omitempty"`
	MailObfuscated *bool `json:"mail_obfuscated,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	OrgName string `json:"orgname,omitempty"`
	Phone string `json:"phone,omitempty"`
	Siren string `json:"siren,omitempty"`
	State string `json:"state,omitempty"`
	Validation string `json:"validation,omitempty"`
	Zip string `json:"zip,omitempty"`
}

type DomainResponseDates struct {
	RegistryCreatedAt time.Time `json:"registry_created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AuthInfoExpiresAt time.Time `json:"authinfo_expires_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	DeletesAt time.Time `json:"deletes_at,omitempty"`
	HoldBeginsAt time.Time `json:"hold_begins_at,omitempty"`
	HoldEndsAt time.Time `json:"hold_ends_at,omitempty"`
	PendingDeleteEndsAt time.Time `json:"pending_delete_ends_at,omitempty"`
	RegistryEndsAt time.Time `json:"registry_ends_at,omitempty"`
	RenewBeginsAt time.Time `json:"renew_begins_at,omitempty"`
	RenewEndsAt time.Time `json:"renew_ends_at,omitempty"`
}

type NameServer struct {
	Current string `json:"string"`
	Hosts []string `json:"hosts,omitempty"`
}

type DomainResponse struct {
	AutoRenew *bool `json:"autorenew"`
	Dates DomainResponseDates `json:"dates"`
	DomainOwner string `json:"domain_owner"`
	FQDN string `json:"fqdn"`
	FQDNUnicode string `json:"fqdn_unicode"`
	Href string `json:"href"`
	ID string `json:"id"`
	NameServer NameServer `json:"nameserver"`
	OrgaOwner string `json:"orga_owner"`
	Owner string `json:"owner"`
	Status []string `json:"status"`
	TLD string `json:"tld"`
	SharingID string `json:"sharing_id,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

func (g *Domain) ListDomains() (domains []DomainResponse, err error) {
	_, err = g.client.Get("domains", nil, &domains)
	return
}

func (g *Domain) GetNameServers(fqdn string) (nameservers []string, err error) {
	_, err = g.client.Get("domains/" + fqdn + "/nameservers", nil, &nameservers)
	return
}
