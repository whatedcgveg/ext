package conf

import (
	"github.com/whatedcgveg/v2ray-core/common/serial"
	"github.com/whatedcgveg/v2ray-core/proxy/http"
)

type HttpServerConfig struct {
	Timeout uint32 `json:"timeout"`
}

func (v *HttpServerConfig) Build() (*serial.TypedMessage, error) {
	config := &http.ServerConfig{
		Timeout: v.Timeout,
	}

	return serial.ToTypedMessage(config), nil
}
