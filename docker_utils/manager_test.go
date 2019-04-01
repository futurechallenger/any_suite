package dockerutils

import (
	"fmt"
	"testing"
)

func TestDockerfile(t *testing.T) {
	ret, err := parepare()
	if err != nil {
		t.Errorf("Get docker file string ERROR %v\n", err)
	}

	fmt.Printf("%s\n", ret)
}
