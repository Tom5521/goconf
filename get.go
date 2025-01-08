package conf

func (p *Preferences) Check(key string) bool {
	_, c := p.values[key]
	return c
}

func (p *Preferences) Load(key string) any {
	return p.values[key]
}

func (p *Preferences) Bool(key string) bool {
	v, ok := p.values[key]
	if !ok {
		v = false
	}
	return v.(bool)
}

func (p *Preferences) BoolList(key string) []bool {
	v, ok := p.values[key]
	if !ok {
		v = []bool(nil)
	}
	return v.([]bool)
}

func (p *Preferences) Float(key string) float64 {
	v, ok := p.values[key]
	if !ok {
		v = 0.0
	}
	return v.(float64)
}

func (p *Preferences) FloatList(key string) []float64 {
	v, ok := p.values[key]
	if !ok {
		v = []float64(nil)
	}
	return v.([]float64)
}

func (p *Preferences) Int(key string) int {
	v, ok := p.values[key]
	if !ok {
		v = 0
	}
	return v.(int)
}

func (p *Preferences) IntList(key string) []int {
	v, ok := p.values[key]
	if !ok {
		v = []int(nil)
	}
	return v.([]int)
}

func (p *Preferences) String(key string) string {
	v, ok := p.values[key]
	if !ok {
		v = ""
	}
	return v.(string)
}

func (p *Preferences) StringList(key string) []string {
	v, ok := p.values[key]
	if !ok {
		v = []string(nil)
	}
	return v.([]string)
}
