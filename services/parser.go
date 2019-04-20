package services

import (
	"encoding/json"
	"fmt"
	"int_ecosys/models"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// Parser parse uploaded scripts
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
	f.WriteString("module.exports = {\n")

	for _, file := range files {
		fmt.Printf("File name: %v\n", file.Name())
		err = p.processFile(f, file)
		// Stops at the first error
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			continue
		}
	}

	f.WriteString("}\n")

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
	// TODO: Parse with AST tree in node
	// Or code like inner function can be a problem
	reg, _ := regexp.Compile(`^function\s[A-Za-z0-9]+\s*\([A-Za-z0-9\s,]*\)\s+\{`)
	for _, l := range lines {
		fmt.Printf("O LINE: %s\n", l)

		var fn string
		if reg.MatchString(l) {
			fn = l[len("function"):len(l)]
			fn = strings.TrimLeft(fn, " ")
			builder.WriteString(fn)
		} else {
			builder.WriteString(l)
		}
	}
	builder.WriteString(",\n")

	to := builder.String()
	fmt.Printf("LINE: %s\n", to)
	_, err = f.WriteString(builder.String())

	if err != nil {
		return fmt.Errorf("Write file: %s error %v", fileName, err)
	}

	return nil
}

// Parse manifest file
func (p *Parser) parseMenfest(fileName string) error {
	var m models.Manifest
	buff, err := ioutil.ReadFile(path.Join(p.sourceDir, fileName))
	json.Unmarshal(buff, &m)

	fmt.Printf("===>String %s, Manifest: %v\n", string(buff), m)

	return err
}
