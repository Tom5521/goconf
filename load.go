package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/user"
	"runtime"
)

func (p *Preferences) findPath() (err error) {
	var usr *user.User

	usr, err = user.Current()
	if err != nil {
		return
	}

	if isUnix() {
		p.configFolder = fmt.Sprintf("%s/.config/%s/", usr.HomeDir, p.id)
	} else if runtime.GOOS == "windows" {
		p.configFolder = fmt.Sprintf("%s\\AppData\\Roaming\\%s\\", usr.HomeDir, p.id)
	}

	return
}

func (p *Preferences) loadConfigFile() (err error) {
	p.configFile = p.configFolder + "config.json"
	p.values = make(map[string]any)

	data, err := os.ReadFile(p.configFile)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(p.configFolder, os.ModePerm)
		if err != nil {
			return err
		}
		err = os.WriteFile(p.configFile, []byte("{}"), 0o600)
		if err != nil {
			return err
		}

		return p.loadConfigFile()
	} else if err != nil {
		return err
	}

	return json.Unmarshal(data, &p.values)
}
