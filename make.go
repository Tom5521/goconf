package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"runtime"
)

type prefs struct {
	ID         string
	Path       string
	File       *os.File
	ConfigFile string
	Map        map[string]any
}

func newPrefs(id string) (*prefs, error) {
	var err error
	p := new(prefs)
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
	p.Map, err = p.read()
	if err != nil {
		return p, err
	}

	return p, nil
}

func (p *prefs) path() (string, error) {
	var path string
	usr, err := user.Current()
	if err != nil {
		return path, err
	}

	basePath := usr.HomeDir + "/"

	unixOS := []string{
		"netbsd",
		"linux",
		"openbsd",
		"darwin",
	}
	isUnix := func() bool {
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

	return path, nil
}

func (p *prefs) new() error {
	if _, err := os.Stat(p.Path); os.IsNotExist(err) {
		err = os.MkdirAll(p.Path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	err := p.write("{}")
	if err != nil {
		return err
	}
	return nil
}

func (p *prefs) write(data string) error {
	err := os.WriteFile(p.ConfigFile, []byte(data), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (p *prefs) read() (map[string]any, error) {
	var m map[string]any
	var err error
	p.File, err = os.Open(p.ConfigFile)
	if err != nil {
		return m, err
	}
	f, err := os.ReadFile(p.ConfigFile)
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(f, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}
