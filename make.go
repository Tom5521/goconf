package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"runtime"
)

const modePerm = os.ModePerm

type prefs struct {
	ID         string
	Path       string
	File       *os.File
	ConfigFile string
	fmap       map[string]any
}

func newPrefs(id string) (p *prefs, err error) {
	p = new(prefs)
	p.ID = id
	p.Path, err = p.path()
	if err != nil {
		return nil, err
	}
	p.ConfigFile = p.Path + "config.json"
	if _, err = os.Stat(p.ConfigFile); os.IsNotExist(err) {
		err = p.new()
		if err != nil {
			return p, err
		}
	}
	p.fmap, err = p.read()
	if err != nil {
		return p, err
	}

	return p, nil
}

func (p *prefs) path() (path string, err error) {
	var (
		usr      *user.User
		basePath string
		isUnix   bool
	)
	usr, err = user.Current()
	if err != nil {
		return
	}

	basePath = usr.HomeDir + "/"

	unixOS := [...]string{
		"netbsd",
		"linux",
		"openbsd",
		"darwin",
	}
	isUnix = func() bool {
		for _, o := range unixOS {
			if runtime.GOOS == o {
				return true
			}
		}
		return false
	}()

	if isUnix {
		path = fmt.Sprintf("%s/.config/%s/", basePath, p.ID)
	}
	if runtime.GOOS == "windows" {
		path = fmt.Sprintf("%s\\AppData\\Roaming\\%s\\", basePath, p.ID)
	}

	return
}

func (p *prefs) new() (err error) {
	if _, err = os.Stat(p.Path); os.IsNotExist(err) {
		err = os.MkdirAll(p.Path, modePerm)
		if err != nil {
			return err
		}
	}
	err = p.writestr("{}")
	if err != nil {
		return err
	}
	return nil
}

func (p *prefs) writestr(data string) error {
	return os.WriteFile(p.ConfigFile, []byte(data), modePerm)
}

func (p *prefs) write(data []byte) error {
	return os.WriteFile(p.ConfigFile, data, modePerm)
}

func (p *prefs) read() (m map[string]any, err error) {
	p.File, err = os.Open(p.ConfigFile)
	if err != nil {
		return
	}
	f, err := os.ReadFile(p.ConfigFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(f, &m)
	if err != nil {
		return
	}
	return
}
