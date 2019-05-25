package services

import (
	"any_suite/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	// ManifestName name of the app manifest
	ManifestName = "appmanifest.json"
)

// Parser parse uploaded scripts
type Parser struct {
	sourceDir      string
	destDir        string
	Manifest       models.Manifest
	manifestParser *ManifestParser
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
// Parse `manifest` and `code.js` first
func (p *Parser) RunParser() error {
	files, err := ioutil.ReadDir(p.sourceDir)
	if err != nil {
		return err
	}

	const (
		manifestFileName = "appmanifest.json"
		destFileName     = "dest.js" // All plugin files will be put into one file, it's `dest.js`
		entryFile        = "code.js" // Entry file of plugins
	)

	// Check `manifest`, `code.js` exists
	// TODO: Check meta files sepeartely
	// destExists := checkFileExists(files, destFileName, true)
	// codeExists := checkFileExists(files, entryFile, true)

	// if !destExists || !codeExists {
	// 	return fmt.Errorf("`appmanifest.json or code.js deso not exist")
	// }

	// Parse `appmanifest.joson`
	// Check if this file exists first
	// If manifest exists, store all info with this *UserID* of current user
	// In redis or some other storage
	manifestExists := checkFileExists(files, manifestFileName, true)
	if manifestExists == false {
		return fmt.Errorf("`appmanifest.json` is not uploaded")
	}

	err = p.parseManifest(destFileName)
	if err != nil {
		return err
	}

	// TODO: Check if code.js includes all method declared in manifest.json

	// Prepare to parse files
	// Create dest file here if it does not exist
	// Create file in destination directory
	f, err := os.Create(path.Join(p.destDir, destFileName))
	if err != nil {
		return fmt.Errorf("Create dest file: %s error %v", destFileName, err)
	}
	f.WriteString("module.exports = {\n")

	for _, file := range files {
		fmt.Printf("File name: %v\n", file.Name())
		// Not parse `appmanifest.jsonnn`
		if file.Name() == ManifestName {
			continue
		}

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

// Check if file is uploaded
// 	- files: uploaded files
// 	- fn: file name, check if this file is in the uploaded files
// 	- full: if it's true, check `fn` as full name,
// 		or check if `fn` is included in the file name
func checkFileExists(files []os.FileInfo, fn string, full bool) bool {
	if files == nil {
		return false
	}

	if fn == "" {
		return false
	}

	for _, fi := range files {
		fileName := fi.Name()
		if full == true && fileName == fn {
			return true
		} else if full == false && strings.Contains(fileName, fn) {
			return true
		}
	}

	return false
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
	// function calls, which all the line only contains
	callReg, _ := regexp.Compile(`^\s*[A-Za-z0-9]+\([\'\"A-Za-z0-9\s,]*\);$`)
	for _, l := range lines {
		fmt.Printf("O LINE: %s\n", l)

		if reg.MatchString(l) {
			fn := l[len("function"):len(l)]
			fn = strings.TrimLeft(fn, " ")
			builder.WriteString(fmt.Sprintf("%s\n", fn))
		} else if callReg.MatchString(l) {
			fncall := fmt.Sprintf("  this.%s\n", strings.TrimLeft(l, " "))
			builder.WriteString(fncall)
		} else {
			builder.WriteString(fmt.Sprintf("%s\n", l))
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

// Parse `appmanifest.json` file
func (p *Parser) parseManifest(fileName string) error {
	var m models.Manifest
	buff, err := ioutil.ReadFile(path.Join(p.sourceDir, fileName))
	json.Unmarshal(buff, &m)

	fmt.Printf("===>String %s, Manifest: %v\n", string(buff), m)

	p.Manifest = m
	manifestParser := NewManifestParser(&m)
	manifestParser.Run()

	return err
}
