package conf

import (
	"encoding/json"

	msg "github.com/Tom5521/GoNotes/pkg/messages"
)

func (p *prefs) writeMap() {
	d, err := json.MarshalIndent(p.Map, "", "	")
	if err != nil {
		msg.Error(err)
	}
	err = p.write(string(d))
	if err != nil {
		msg.Error(err)
	}
}

func (p *prefs) Set(key string, value any) {
	p.Map[key] = value
	p.writeMap()
}

func (p *prefs) Read(key string) any {
	return p.Map[key]
}

func (p *prefs) Bool(key string) bool {
	v, ok := p.Map[key]
	if !ok {
		v = false
		p.Set(key, v)
	}
	return v.(bool)
}

func (p *prefs) SetBool(key string, value bool) {
	p.Set(key, value)
}

func (p *prefs) BoolList(key string) []bool {
	v, ok := p.Map[key]
	if !ok {
		v = []bool(nil)
		p.Set(key, v)
	}
	return v.([]bool)
}

func (p *prefs) SetBoolList(key string, value []bool) {
	p.Set(key, value)
}

func (p *prefs) Float(key string) float64 {
	v, ok := p.Map[key]
	if !ok {
		v = 0.0
		p.Set(key, v)
	}
	return v.(float64)
}

func (p *prefs) SetFloat(key string, value float64) {
	p.Set(key, value)
}

func (p *prefs) FloatList(key string) []float64 {
	v, ok := p.Map[key]
	if !ok {
		v = []float64(nil)
		p.Set(key, v)
	}
	return v.([]float64)
}

func (p *prefs) SetFloatList(key string, value []float64) {
	p.Set(key, value)
}

func (p *prefs) Int(key string) int {
	v, ok := p.Map[key]
	if !ok {
		v = 0
		p.Set(key, v)
	}
	return v.(int)
}

func (p *prefs) SetInt(key string, value int) {
	p.Set(key, value)
}

func (p *prefs) IntList(key string) []int {
	v, ok := p.Map[key]
	if !ok {
		v = []int(nil)
		p.Set(key, v)
	}
	return v.([]int)
}

func (p *prefs) SetIntList(key string, value []int) {
	p.Set(key, value)
}

func (p *prefs) String(key string) string {
	v, ok := p.Map[key]
	if !ok {
		v = ""
		p.Set(key, v)
	}
	return v.(string)
}

func (p *prefs) SetString(key string, value string) {
	p.Set(key, value)
}

func (p *prefs) StringList(key string) []string {
	v, ok := p.Map[key]
	if !ok {
		v = []string(nil)
		p.Set(key, v)
	}
	return v.([]string)
}

func (p *prefs) SetStringList(key string, value []string) {
	p.Set(key, value)
}

func (p *prefs) RemoveValue(key string) {
	delete(p.Map, key)
	p.writeMap()
}
