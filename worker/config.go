package worker

import (
	"encoding/json"
	"io/ioutil"
)

const (
	defaultMaxMsgSize = 1024 * 1024 * 4
)

type Config struct {
	HTTPEndpoint    string
	GRPCEndpoint    string
	MetricsEndpoint string
	MaxMsgSize      int // in MB
}

func LoadConfig(fn string) (*Config, error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	c := Config{}
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	if c.MaxMsgSize <= 0 {
		c.MaxMsgSize = defaultMaxMsgSize
	} else {
		c.MaxMsgSize = c.MaxMsgSize * 1024 * 1024
	}

	return &c, nil
}
