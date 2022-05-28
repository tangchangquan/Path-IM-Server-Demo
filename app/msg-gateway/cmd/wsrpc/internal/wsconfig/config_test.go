package wsconfig

import (
	"github.com/Path-IM/Path-IM-Server-Demo/common/xconf"
	"testing"
)

func TestConfig(t *testing.T) {
	conf := &Config{}
	xconf.MustLoad("", conf)
}
