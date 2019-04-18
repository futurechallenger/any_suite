package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// Parser parse uploaded scripts
// Uploaded scripts are required to compose like every function
// are in `global` scope.
// No sub directories in uploaded scripts
// NOTE:
// 1. Check if the file is `.js`
// 2. Put all code in one file maybe the best way, what about the name conflicts
// 3. Upload multiple files one time
type Parser struct {
	sourceDir string
	destDir   string
}

// NewParser return `Parser` instance
func NewParser(sourceDir string, destDir string) (p *Parser, err error) {
	var source string
	if sourceDir == "" {
		dir, err := filepath.Abs("./store/tmp")
		if err != nil {
			return nil, err
		}
		source = dir
	}

	var dest string
	if destDir == "" {
		dir, err := filepath.Abs("./store")
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
func (p *Parser) RunParser() error {
	files, err := ioutil.ReadDir(p.sourceDir)
	if err != nil {
		return err
	}

	// Prepare to parse files
	// Create dest file here if it does not exist
	// Create file in destination directory
	const destFileName = "dest.js"
	f, err := os.Create(path.Join(p.destDir, destFileName))
	if err != nil {
		return fmt.Errorf("Create dest file: %s error %v", destFileName, err)
	}

	for _, file := range files {
		fmt.Printf("File name: %v\n", file.Name())
		err = p.processFile(f, file)
		// Stops at the first error
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			continue
		}
	}

	return nil
}

// checkFileExt checks file's extenstion is equals to the given one
// Default is `.js`
func (p *Parser) checkFileExt(fileName string, fileExt string) (bool, error) {
	if fileName == "" {
		return false, fmt.Errorf("`fileName` is invalid")
	}

	if fileExt == "" {
		fileExt = "js"
	}

	return strings.Index(fileName, fileExt) >= 0, nil
}

// processFile process files
func (p *Parser) processFile(f *os.File, file os.FileInfo) error {
	fileName := file.Name()

	ok, _ := p.checkFileExt(fileName, "")
	if ok == false {
		return fmt.Errorf("This file `%s` is not the type is suppposed to be", fileName)
	}

	buff, err := ioutil.ReadFile(path.Join(p.sourceDir, fileName))
	// File content
	raw := string(buff)
	lines := strings.Split(raw, "\n")

	fmt.Printf("File Content : %s\n", raw)

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
	builder.WriteString("\n")

	_, err = f.WriteString(builder.String())

	if err != nil {
		return fmt.Errorf("Write file: %s error %v", fileName, err)
	}

	return nil
}
