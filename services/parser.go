package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Parser parse uploaded scripts
// Uploaded scripts are required to compose like every function
// are in `global` scope.
// No sub directories in uploaded scripts
type Parser struct {
	sourceDir string
	destDir   string
}

// NewParser return `Parser` instance
func NewParser(sourceDir string, destDir string) (p *Parser, err error) {
	var source string
	if sourceDir == "" {
		dir, err := filepath.Abs("../store/temp")
		if err != nil {
			return nil, err
		}
		source = dir
	}

	var dest string
	if destDir == "" {
		dir, err := filepath.Abs("../store")
		if err != nil {
			return nil, err
		}
		dest = dir
	}

	return &Parser{
		sourceDir: source,
		destDir:   dest,
	}, nil
}

// RunParser start parse uploaded scripts
func (p *Parser) RunParser(uploadedPath string) error {
	files, err := ioutil.ReadDir(p.sourceDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		err = p.processFile(file)
		// Stops at the first error
		if err != nil {
			return err
		}
	}

	return nil
}

// processFile process files
func (p *Parser) processFile(file os.FileInfo) error {
	fileName := file.Name()
	buff, err := ioutil.ReadFile(fileName)
	fmt.Printf("file name: %s\n", string(buff))

	// Create file in destination directory
	f, err := os.Create(fmt.Sprintf("%s/%s", p.destDir, file.Name()))
	if err != nil {
		return fmt.Errorf("Create file: %s error %v", fileName, err)
	}

	raw := string(buff) // File content
	lines := strings.Split(raw, "\n")

	var builder strings.Builder
	reg, _ := regexp.Compile(`^function\s[A-Za-z0-9]+\s*\([A-Za-z0-9\s,]*\)\s+\{`)
	for _, l := range lines {
		var funName string
		if reg.MatchString(l) {
			funName = l[len("function"):len(l)]
			funName = strings.TrimLeft(l, " ")
			builder.WriteString(funName)
		} else {
			builder.WriteString(l)
		}
	}

	_, err = f.WriteString(builder.String())

	if err != nil {
		return fmt.Errorf("Write file: %s error %v", fileName, err)
	}

	return nil
}
