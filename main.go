package conf

func New(id string) (Preferences, error) {
	p, err := newPrefs(id)
	if err != nil {
		return p, err
	}

	return p, nil
}
