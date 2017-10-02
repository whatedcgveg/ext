package conf

import (
	"github.com/whatedcgveg/v2ray-core/common/serial"
)

type Buildable interface {
	Build() (*serial.TypedMessage, error)
}
