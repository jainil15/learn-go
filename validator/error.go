package validator

type ValidationError map[string][]string

func (v ValidationError) Add(key, message string) {
	v[key] = append(v[key], message)
}

func (v ValidationError) IsEmpty() bool {
	return len(v) == 0
}
