package services

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// Container deals with Docker containers
type Container struct {
}

// CheckInstalled used to check if Docker is installed
func (c *Container) CheckInstalled() (installed int16, err error) {
	fmt.Printf("Check docker install on %s\n", runtime.GOOS)

	// cmd := exec.Command("docker", "-v")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	//
	// if err := cmd.Run(); err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	//  }

	bs, err := exec.Command("docker", "-v").Output()
	if err != nil {
		return 0, err
	}

	output := string(bs)
	if strings.Contains(output, "Docker") == true {
		fmt.Println("Docker is installed")
	}
	fmt.Printf("Check result %s\n", output)

	return 0, nil
}
