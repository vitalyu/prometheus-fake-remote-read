package main

import (
	"encoding/json"
	"os"

	"github.com/prometheus/common/model"
)

type Series struct {
	Interval model.Duration `json:"interval"`
	Series   string         `json:"series"`
	Values   string         `json:"values"`
}

type Configuration struct {
	LogLevel    string   `json:"log_level"`
	InputSeries []Series `json:"input_series"`
}

func (c *Configuration) LoadConfig(fname string) (err error) {
	file, err := os.Open(fname)
	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		return
	}
	return
}

func (c *Configuration) SaveConfig(fname string) (err error) {
	file, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return
	}
	err = os.WriteFile(fname, file, 0644)
	return
}
