package config

import (
	"github.com/showurl/Path-IM-Server/common/xenv"
	"testing"
)

func TestParseConfig(t *testing.T) {
	t.Setenv("ZEROIM_REDISCONFIG_CONF__HOST", "10.0.8.6:6379")
	conf := &Config{}
	xenv.Parse(conf)
	t.Log(conf.Jpns.ApnsProduction)
}
