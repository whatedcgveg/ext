package conf_test

import (
	"encoding/json"
	"testing"

	v2net "github.com/whatedcgveg/v2ray-core/common/net"
	"github.com/whatedcgveg/v2ray-core/testing/assert"
	. "github.com/whatedcgveg/ext/tools/conf"
)

func TestDnsConfigParsing(t *testing.T) {
	assert := assert.On(t)

	rawJson := `{
    "servers": ["8.8.8.8"]
  }`

	jsonConfig := new(DnsConfig)
	err := json.Unmarshal([]byte(rawJson), jsonConfig)
	assert.Error(err).IsNil()

	config := jsonConfig.Build()
	assert.Int(len(config.NameServers)).Equals(1)
	dest := config.NameServers[0].AsDestination()
	assert.Destination(dest).IsUDP()
	assert.Address(dest.Address).Equals(v2net.IPAddress([]byte{8, 8, 8, 8}))
	assert.Port(dest.Port).Equals(v2net.Port(53))
}
