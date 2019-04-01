package dockerutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func parepare() (docker string, err error) {
	buff, err := ioutil.ReadFile("docker_template")
	if err != nil {
		return "", fmt.Errorf("read file error %v", err)
	}

	ret := string(buff)

	// Replace placeholders
	repo := os.Getenv("GIT_REPO_PATH")
	dockerFile := strings.ReplaceAll(ret, "{{GIT_REPO_PATH}}", repo)

	return dockerFile, nil
}

// GenerateDockerfile generate a `Dockerfile`
func GenerateDockerfile() error {
	ret, err := parepare()
	if err != nil {
		return err
	}

	f, err := os.Create("Dockerfile")
	if err != nil {
		return fmt.Errorf("Create `Dockerfile` error %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(ret)
	if err != nil {
		return fmt.Errorf("Write string to `Dockerfile` error %v", err)
	}

	f.Sync()

	return nil
}
