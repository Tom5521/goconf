package conf

import (
	"encoding/json"
	"os"
)

type Preferences struct {
	id   string
	path string

	file    *os.File
	decoder *json.Decoder
	encoder *json.Encoder

	values map[string]any

	configFile string
}

func New(id string) (p *Preferences, err error) {
	p = &Preferences{
		id: id,
	}
	err = p.findPath()
	if err != nil {
		return
	}

	err = p.loadConfigFile()

	return
}

func (p Preferences) ConfigFile() string {
	return p.configFile
}

func (p Preferences) Values() map[string]any {
	return p.values
}

func (p *Preferences) Close() error {
	return p.file.Close()
}

func (p *Preferences) Save() error {
	return p.encoder.Encode(p.values)
}

func (p *Preferences) RemoveValue(key string) {
	delete(p.values, key)
}
