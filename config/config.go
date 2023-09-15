package config

import (
	"flag"
	"strconv"
)

var (
	port = flag.Int("port", 9180, "runner port")
)

type Config struct {
	HTTPAddr string
}

func NewConfig() (*Config, error) {

	flag.Parse()
	config := &Config{
		HTTPAddr: ":" + strconv.Itoa(*port),
	}
	return config, nil
}
