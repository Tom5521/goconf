package conf

import (
	"encoding/json"
	"fmt"
)

func (p *prefs) writeMap() {
	d, err := json.MarshalIndent(p.fmap, "", "	")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	err = p.write(d)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func (p *prefs) Map() map[string]any {
	return p.fmap
}

func (p *prefs) Set(key string, value any) {
	p.fmap[key] = value
	p.writeMap()
}

func (p *prefs) Read(key string) any {
	return p.fmap[key]
}

func (p *prefs) Bool(key string) bool {
	v, ok := p.fmap[key]
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
	v, ok := p.fmap[key]
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
	v, ok := p.fmap[key]
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
	v, ok := p.fmap[key]
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
	v, ok := p.fmap[key]
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
	v, ok := p.fmap[key]
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
	v, ok := p.fmap[key]
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
	v, ok := p.fmap[key]
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
	delete(p.fmap, key)
	p.writeMap()
}
