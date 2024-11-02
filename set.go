package conf

import "reflect"

func SetFor[T any](p *Preferences, key string, v T) {
	p.Set(key, v)
}

func (p *Preferences) SetValue(key string, value reflect.Value) {
	p.values[key] = value.Interface()
}

func (p *Preferences) Set(key string, value any) {
	p.values[key] = value
	if p.Autosave {
		p.autoSaveErr = p.Save()
	}
}

func (p *Preferences) SetBool(key string, value bool) {
	p.Set(key, value)
}

func (p *Preferences) SetBoolList(key string, value []bool) {
	p.Set(key, value)
}

func (p *Preferences) SetFloat(key string, value float64) {
	p.Set(key, value)
}

func (p *Preferences) SetFloatList(key string, value []float64) {
	p.Set(key, value)
}

func (p *Preferences) SetInt(key string, value int) {
	p.Set(key, value)
}

func (p *Preferences) SetIntList(key string, value []int) {
	p.Set(key, value)
}

func (p *Preferences) SetString(key string, value string) {
	p.Set(key, value)
}

func (p *Preferences) SetStringList(key string, value []string) {
	p.Set(key, value)
}
