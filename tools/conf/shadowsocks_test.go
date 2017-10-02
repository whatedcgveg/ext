package conf_test

import (
	"encoding/json"
	"testing"

	"github.com/whatedcgveg/v2ray-core/proxy/shadowsocks"
	"github.com/whatedcgveg/v2ray-core/testing/assert"
	. "github.com/whatedcgveg/ext/tools/conf"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	assert := assert.On(t)

	rawJson := `{
    "method": "aes-128-cfb",
    "password": "v2ray-password"
  }`

	rawConfig := new(ShadowsocksServerConfig)
	err := json.Unmarshal([]byte(rawJson), rawConfig)
	assert.Error(err).IsNil()

	ts, err := rawConfig.Build()
	assert.Error(err).IsNil()
	iConfig, err := ts.GetInstance()
	assert.Error(err).IsNil()
	config := iConfig.(*shadowsocks.ServerConfig)

	rawAccount, err := config.User.GetTypedAccount()
	assert.Error(err).IsNil()

	account, ok := rawAccount.(*shadowsocks.ShadowsocksAccount)
	assert.Bool(ok).IsTrue()

	assert.Int(account.Cipher.KeySize()).Equals(16)
	assert.Bytes(account.Key).Equals([]byte{160, 224, 26, 2, 22, 110, 9, 80, 65, 52, 80, 20, 38, 243, 224, 241})
}
