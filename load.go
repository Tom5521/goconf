package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/user"
	"runtime"
)

func (p *Preferences) findPaths() (err error) {
	var usr *user.User

	usr, err = user.Current()
	if err != nil {
		return
	}

	if IsUnix {
		p.configFolder = fmt.Sprintf("%s/.config/%s/", usr.HomeDir, p.id)
	} else if runtime.GOOS == "windows" {
		p.configFolder = fmt.Sprintf("%s\\AppData\\Roaming\\%s\\", usr.HomeDir, p.id)
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
