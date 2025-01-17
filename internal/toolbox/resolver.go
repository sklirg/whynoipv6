package toolbox

import (
	"errors"
	"net"
	"strings"
	"time"

	"github.com/miekg/dns"
)

const (
	// DefaultTimeout stops slow dns's
	DefaultTimeout = 20 * time.Second
	// MaxReturnedIPAddressesCount stops overflow
	MaxReturnedIPAddressesCount = 64
)

// Errors returned by the verification/validation methods at all levels.
var (
	ErrNoResult       = errors.New("requested RR not found")
	ErrNsNotAvailable = errors.New("no name server to answer the question")
	ErrInvalidQuery   = errors.New("invalid query input")
)

var resolver *Resolver

// Resolver contains the client configuration for github.com/miekg/dns,
// the instantiated client and the func that performs the actual queries.
// queryFn can be used for mocking the actual DNS lookups in the test suite.
type Resolver struct {
	Query           func(string, uint16) (*dns.Msg, error)
	dnsClient       *dns.Client
	dnsClientConfig *dns.ClientConfig
}

// NewDNSMessage creates and initializes a dns.Msg object.
// It returns a pointer to the created object.
func (s *Service) NewDNSMessage() *dns.Msg {
	dnsMessage := &dns.Msg{
		MsgHdr: dns.MsgHdr{
			RecursionDesired: true,
		},
	}
	return dnsMessage
}

// NewResolver initializes the package Resolver instance using the default
// dnsClientConfig.
func (s *Service) NewResolver() (res *Resolver, err error) {
	resolver = &Resolver{}
	resolver.dnsClient = &dns.Client{
		ReadTimeout: DefaultTimeout,
	}
	resolver.dnsClientConfig, err = dns.ClientConfigFromReader(strings.NewReader(s.Nameserver))
	if err != nil {
		return nil, err
	}
	resolver.Query = s.localQuery
	return resolver, nil
}

// localQuery takes a query name (qname) and query type (qtype) and
// performs a DNS lookup by calling dnsClient.Exchange.
// It returns the answer in a *dns.Msg (or nil in case of an error, in which
// case err will be set accordingly.)
func (s *Service) localQuery(qname string, qtype uint16) (*dns.Msg, error) {
	dnsMessage := s.NewDNSMessage()
	dnsMessage.SetQuestion(dns.Fqdn(qname), qtype)

	if resolver.dnsClientConfig == nil {
		return nil, errors.New("dns client not initialized")
	}

	for _, server := range resolver.dnsClientConfig.Servers {
		r, _, err := resolver.dnsClient.Exchange(dnsMessage, net.JoinHostPort(server, resolver.dnsClientConfig.Port))
		if err != nil {
			return nil, err
		}
		if r == nil || r.Rcode == dns.RcodeNameError || r.Rcode == dns.RcodeSuccess {
			return r, err
		}
	}
	return nil, ErrNsNotAvailable
}

// QueryResult is the result from a dns lookup
type QueryResult struct {
	IPv6  bool
	Rcode int
}
