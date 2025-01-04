package conf

import (
	"encoding/json"
	"errors"
	"os"
)

func (p *Preferences) findPaths() (err error) {
	p.configFolder, err = os.UserConfigDir()
	if err != nil {
		return
	}

	p.configFile = p.configFolder + "config.json"

	return
}

func (p *Preferences) loadConfigFile() (err error) {
	data, err := os.ReadFile(p.configFile)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(p.configFolder, os.ModePerm)
		if err != nil {
			return
		}
		err = os.WriteFile(p.configFile, []byte("{}"), 0o600)
		if err != nil {
			return
		}

		return p.loadConfigFile()
	} else if err != nil {
		return
	}

	return json.Unmarshal(data, &p.values)
}
