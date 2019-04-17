package services

import (
	"fmt"
	"os/exec"
)

var holder = make(chan string, 1)

// Run runs a event handler
func Run() error {
	holder <- "start"

	out, err := exec.Command("sh", "-c", "../Dragon -c=./config.json").CombinedOutput()
	fmt.Printf("Dragon output : %v", string(out))
	if err != nil {
		<-holder
		return fmt.Errorf("Execute `Dragon` error %v", err)
	}

	<-holder
	return nil
}
