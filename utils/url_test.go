package utils

import (
	"testing"
)

func TestURLEncode(t *testing.T) {
	if URLEncode(map[string]string{"a": "111", "b": "222"}) != "a=111&b=222" {
		t.Error("Expected result is a=111&b=222")
	}
}
