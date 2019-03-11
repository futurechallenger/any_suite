package utils

// Condition means condition ? ok : defaultVal
func Condition(condition func() bool, ok interface{}, defaultVal interface{}) interface{} {
	if condition() == true {
		return ok
	}

	return defaultVal
}
