package utils

import "fmt"

// URLEncode generate querystring from a map
func URLEncode(params map[string]string) string {
	if params == nil {
		return ""
	}

	queryString := ""
	count := 0
	for key, val := range params {
		if count == 0 {
			queryString = fmt.Sprintf("%s=%s", key, val)
		} else {
			queryString = fmt.Sprintf("%s&%s=%s", queryString, key, val)
		}

		count = count + 1
	}

	return queryString
}
