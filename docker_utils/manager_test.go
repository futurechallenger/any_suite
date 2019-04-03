package dockerutils

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestDockerfile(t *testing.T) {
	ret, err := parepare()
	if err != nil {
		t.Errorf("Get docker file string ERROR %v\n", err)
	}

	fmt.Printf("%s\n", ret)
}

func TestGeneratedDockerFile(t *testing.T) {
	err := GenerateDockerfile()
	if err != nil {
		t.Errorf("Generate docker file error > %v\n", err)
	}

	_, err = os.Stat("Dockerfile")
	if os.IsNotExist(err) || err != nil {
		t.Errorf("Dockerfile does not exist")
	}
}

func TestCreateImage(t *testing.T) {
	ret, err := CreateImage()
	if err != nil {
		t.Errorf("Create docker image error %v\n", err)
	}
	fmt.Printf("Execute docker command output %v\n", ret)

	_, err = exec.Command("docker", "images").Output()
	if err != nil {
		t.Errorf("Docker may not installed on this computer")
	}
}
