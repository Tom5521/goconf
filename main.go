package conf

import (
	"encoding/json"
	"os"
	"reflect"
)

type Preferences struct {
	// If true, the file will be written every time a Set operation is executed.
	Autosave bool

	id           string
	configFolder string

	values map[string]any

	configFile string

	autoSaveErr error
}

func New(id string) (p *Preferences, err error) {
	p = &Preferences{
		id:       id,
		values:   make(map[string]any),
		Autosave: true,
	}
	err = p.findPaths()
	if err != nil {
		return
	}

	err = p.loadConfigFile()

	return
}

type Field struct {
	Key       string
	FieldType reflect.Type
}

func (p *Preferences) CreateNewFields(overwrite bool, fields ...Field) {
	for _, f := range fields {
		if !overwrite && p.Check(f.Key) {
			continue
		}
		p.Set(f.Key, reflect.Zero(f.FieldType).Interface())
	}
}

func (p Preferences) ConfigFile() string {
	return p.configFile
}

func (p Preferences) ID() string {
	return p.id
}

func (p Preferences) Values() map[string]any {
	return p.values
}

func (p *Preferences) Save() error {
	data, err := json.MarshalIndent(p.values, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(p.configFile, data, 0o600)
}

func (p *Preferences) Refresh() error {
	data, err := os.ReadFile(p.configFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &p.values)
}

// Returns possible errors made when autosaving.
func (p *Preferences) Error() error {
	return p.autoSaveErr
}

func (p *Preferences) RemoveValue(key string) {
	delete(p.values, key)
}
