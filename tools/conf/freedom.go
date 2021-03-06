package conf

import (
	"net"
	"strings"

	v2net "github.com/whatedcgveg/v2ray-core/common/net"
	"github.com/whatedcgveg/v2ray-core/common/protocol"
	"github.com/whatedcgveg/v2ray-core/common/serial"
	"github.com/whatedcgveg/v2ray-core/proxy/freedom"
)

type FreedomConfig struct {
	DomainStrategy string  `json:"domainStrategy"`
	Timeout        *uint32 `json:"timeout"`
	Redirect       string  `json:"redirect"`
}

func (v *FreedomConfig) Build() (*serial.TypedMessage, error) {
	config := new(freedom.Config)
	config.DomainStrategy = freedom.Config_AS_IS
	domainStrategy := strings.ToLower(v.DomainStrategy)
	if domainStrategy == "useip" || domainStrategy == "use_ip" {
		config.DomainStrategy = freedom.Config_USE_IP
	}
	config.Timeout = 600
	if v.Timeout != nil {
		config.Timeout = *v.Timeout
	}
	if len(v.Redirect) > 0 {
		host, portStr, err := net.SplitHostPort(v.Redirect)
		if err != nil {
			return nil, newError("invalid redirect address: ", v.Redirect, ": ", err).Base(err)
		}
		port, err := v2net.PortFromString(portStr)
		if err != nil {
			return nil, newError("invalid redirect port: ", v.Redirect, ": ", err).Base(err)
		}
		if len(host) == 0 {
			host = "127.0.0.1"
		}
		config.DestinationOverride = &freedom.DestinationOverride{
			Server: &protocol.ServerEndpoint{
				Address: v2net.NewIPOrDomain(v2net.ParseAddress(host)),
				Port:    uint32(port),
			},
		}
	}
	return serial.ToTypedMessage(config), nil
}
