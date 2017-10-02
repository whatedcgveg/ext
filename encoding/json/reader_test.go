package json_test

import (
	"bytes"
	"testing"

	"github.com/whatedcgveg/v2ray-core/testing/assert"
	. "github.com/whatedcgveg/ext/encoding/json"
)

func TestReader(t *testing.T) {
	assert := assert.On(t)

	data := []struct {
		input  string
		output string
	}{
		{
			`
content #comment 1
#comment 2
content 2`,
			`
content content 2`},
		{`content`, `content`},
		{" ", " "},
		{`con/*abcd*/tent`, "content"},
		{`
text // adlkhdf /*
//comment adfkj
text 2*/`, `
text text 2*`},
		{`"//"content`, `"//"content`},
		{`abcd'//'abcd`, `abcd'//'abcd`},
		{`"\""`, `"\""`},
		{`\"/*abcd*/\"`, `\"\"`},
	}

	for _, testCase := range data {
		reader := &Reader{
			Reader: bytes.NewReader([]byte(testCase.input)),
		}

		actual := make([]byte, 1024)
		n, err := reader.Read(actual)
		assert.Error(err).IsNil()
		assert.String(string(actual[:n])).Equals(testCase.output)
	}
}
