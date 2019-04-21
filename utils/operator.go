package utils

import "reflect"

// Condition means condition ? ok : defaultVal
func Condition(condition func() bool, ok interface{}, defaultVal interface{}) interface{} {
	if condition() == true {
		return ok
	}

	return defaultVal
}

// HasElem is a util method to check if an element is in an array
func HasElem(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {

			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

// HasStringElem check if a string included in a string array
func HasStringElem(l []string, el string) int {
	if l == nil {
		return -1
	}

	for i, v := range l {
		if v == el {
			return i
		}
	}

	return -1
}
