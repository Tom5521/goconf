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
		p.path = fmt.Sprintf("%s/.config/%s/", usr.HomeDir, p.id)
	} else if runtime.GOOS == "windows" {
		p.path = fmt.Sprintf("%s\\AppData\\Roaming\\%s\\", usr.HomeDir, p.id)
	}

	return
}

func (p *Preferences) loadConfigFile() (err error) {
	p.configFile = p.path + "config.json"

	p.file, err = os.OpenFile(p.configFile, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(p.path, os.ModePerm)
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

	p.decoder = json.NewDecoder(p.file)
	p.encoder = json.NewEncoder(p.file)

	return p.decoder.Decode(&p.values)
}
