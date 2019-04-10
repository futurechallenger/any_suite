package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	targetFile, err := filepath.Abs("../store/tmp/file_list.txt")
	if err != nil {
		t.Error("Can not get target file")
	}

	dir, err := ioutil.ReadDir("./")
	if err != nil {
		t.Error("Read dir failed")
	}

	var f *os.File
	// if _, err := os.Stat(targetFile); os.IsNotExist(err) {
	// Create a new file
	f, err = os.Create(targetFile)
	if err != nil {
		t.Errorf("Create target file error %v", err)
	}
	// }

	for _, file := range dir {
		fmt.Printf("File name:- %s\n", file.Name())
		// err = ioutil.WriteFile(targetFile, []byte(file.Name()), 0644)
		_, err = f.WriteString(fmt.Sprintf("%s\n", file.Name()))
		if err != nil {
			t.Errorf("Write file failed: %v", err)
			return
		}
	}
}

func TestFileContent(t *testing.T) {
	const (
		header  = "module.exports = {"
		content = "function yo(name) {return `yo ${name}!`}"
		footer  = "}"
	)

	var builder strings.Builder
	builder.WriteString(header)

}
