package conf_test

import (
	"encoding/json"
	"testing"

	"github.com/whatedcgveg/v2ray-core/testing/assert"
	"github.com/whatedcgveg/v2ray-core/transport/internet"
	"github.com/whatedcgveg/v2ray-core/transport/internet/headers/http"
	"github.com/whatedcgveg/v2ray-core/transport/internet/headers/noop"
	"github.com/whatedcgveg/v2ray-core/transport/internet/kcp"
	"github.com/whatedcgveg/v2ray-core/transport/internet/tcp"
	"github.com/whatedcgveg/v2ray-core/transport/internet/websocket"
	. "github.com/whatedcgveg/ext/tools/conf"
)

func TestTransportConfig(t *testing.T) {
	assert := assert.On(t)

	rawJson := `{
    "tcpSettings": {
      "header": {
        "type": "http",
        "request": {
          "version": "1.1",
          "method": "GET",
          "path": "/b",
          "headers": {
            "a": "b",
            "c": "d"
          }
        },
        "response": {
          "version": "1.0",
          "status": "404",
          "reason": "Not Found"
        }
      }
    },
    "kcpSettings": {
      "mtu": 1200,
      "header": {
        "type": "none"
      }
    },
    "wsSettings": {
      "path": "/t"
    }
  }`

	var transportSettingsConf TransportConfig
	assert.Error(json.Unmarshal([]byte(rawJson), &transportSettingsConf)).IsNil()

	ts, err := transportSettingsConf.Build()
	assert.Error(err).IsNil()

	assert.Int(len(ts.TransportSettings)).Equals(3)
	var settingsCount uint32
	for _, settingsWithProtocol := range ts.TransportSettings {
		rawSettings, err := settingsWithProtocol.Settings.GetInstance()
		assert.Error(err).IsNil()
		switch settings := rawSettings.(type) {
		case *tcp.Config:
			settingsCount++
			assert.Bool(settingsWithProtocol.Protocol == internet.TransportProtocol_TCP).IsTrue()
			rawHeader, err := settings.HeaderSettings.GetInstance()
			assert.Error(err).IsNil()
			header := rawHeader.(*http.Config)
			assert.String(header.Request.GetVersionValue()).Equals("1.1")
			assert.String(header.Request.Uri[0]).Equals("/b")
			assert.String(header.Request.Method.Value).Equals("GET")
			var va, vc string
			for _, h := range header.Request.Header {
				switch h.Name {
				case "a":
					va = h.Value[0]
				case "c":
					vc = h.Value[0]
				default:
					t.Error("Unknown header ", h.String())
				}
			}
			assert.String(va).Equals("b")
			assert.String(vc).Equals("d")
			assert.String(header.Response.Version.Value).Equals("1.0")
			assert.String(header.Response.Status.Code).Equals("404")
			assert.String(header.Response.Status.Reason).Equals("Not Found")
		case *kcp.Config:
			settingsCount++
			assert.Bool(settingsWithProtocol.Protocol == internet.TransportProtocol_MKCP).IsTrue()
			assert.Uint32(settings.GetMTUValue()).Equals(1200)
			rawHeader, err := settings.HeaderConfig.GetInstance()
			assert.Error(err).IsNil()
			header := rawHeader.(*noop.Config)
			assert.Pointer(header).IsNotNil()
		case *websocket.Config:
			settingsCount++
			assert.Bool(settingsWithProtocol.Protocol == internet.TransportProtocol_WebSocket).IsTrue()
			assert.String(settings.Path).Equals("/t")
		default:
			t.Error("Unknown type of settings.")
		}
	}
	assert.Uint32(settingsCount).Equals(3)
}
