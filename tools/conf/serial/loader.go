package serial

import (
	"encoding/json"
	"io"

	"github.com/whatedcgveg/v2ray-core"
	json_reader "github.com/whatedcgveg/ext/encoding/json"
	"github.com/whatedcgveg/ext/tools/conf"
)

func LoadJSONConfig(reader io.Reader) (*core.Config, error) {
	jsonConfig := &conf.Config{}
	decoder := json.NewDecoder(&json_reader.Reader{
		Reader: reader,
	})

	if err := decoder.Decode(jsonConfig); err != nil {
		return nil, newError("failed to read config file").Base(err)
	}

	pbConfig, err := jsonConfig.Build()
	if err != nil {
		return nil, newError("failed to parse json config").Base(err)
	}

	return pbConfig, nil
}
