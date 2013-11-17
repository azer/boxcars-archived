package JSONConfig

import (
	"encoding/json"
	"io/ioutil"
	. "github.com/azer/on-change"
	. "github.com/azer/debug"
)

type (
	Document map[string]map[string]string
	RawDocument map[string]interface{}
)

type JSONConfig struct {
	Filename string
	Document Document
	callback func(map[string]map[string]string)
}

func (config *JSONConfig) Load () {
	raw, err := Read(config.Filename)

	if err != nil {
		Debug("Failed to read and parse %s", config.Filename)
		return
	}

	config.Document = Normalize(raw)

	go config.callback(config.Document)
}

func (config *JSONConfig) EnableAutoReload () {
	OnChange(config.Filename, config.Load)
}

func NewJSONConfig (filename string, callback func(map[string]map[string]string)) *JSONConfig {
	Debug("Creating a new JSON config from %s", filename)

	config := &JSONConfig{filename, nil, callback}
	config.Load()

	return config
}

func Normalize (raw RawDocument) Document {
	Debug("Loading...")

	config := make(Document)

	for hostname, options := range raw {
		switch t := options.(type) {
		case string:
			config[hostname] = make(map[string]string)
			config[hostname]["/"] = t
		case map[string]interface{}:
			config[hostname] = make(map[string]string)

			for path, uri := range t {
				if str, ok := uri.(string); ok {
					config[hostname][path] = str
				}
			}
		}
	}

	return config
}

func Read (filename string) (RawDocument, error) {
	Debug("Reading %s", filename)

	raw := make(RawDocument)
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &raw)

	if err != nil {
		return nil, err
	}

	return raw, nil
}
