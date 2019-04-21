package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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

func TestFindJsFuncName(t *testing.T) {
	r, _ := regexp.Compile(`^function\s[A-Za-z0-9]+\s*\([A-Za-z0-9\s,]*\)\s+\{`)
	matched := r.MatchString(`function aaa(p1, p2) {
		console.log('hello world');
	}`)

	if matched == false {
		t.Error("Not matched function")
	}

	fn := `function aaa(p1, p2) {`
	ret := fn[len("function"):len(fn)]
	ret = strings.TrimLeft(ret, " ")

	if ret != "aaa" {
		t.Error("Retrive function name error")
	}
}

func TestGetFuncName(t *testing.T) {
	content := "function yo(name) {\nreturn `yo ${name}!`\n}"
	funName := content[len("function"):strings.Index(content, "(")]
	if strings.TrimSpace(funName) != "yo" {
		t.Errorf("Func name: `%s` does not found", funName)
	}
}

func TestFileExt(t *testing.T) {
	const (
		rightFile  = "Hello.js"
		wrongFile  = "hello.txt"
		wrongFile2 = "bro"
	)

	p, err := NewParser("", "")
	if err != nil {
		t.Error(err)
	}

	if ok, err := p.checkFileExt(rightFile, ""); ok == false || err != nil {
		t.Error("Should be right")
	}

	if ok, err := p.checkFileExt(wrongFile, ""); ok == true || err != nil {
		t.Error("Should be right")
	}

	if ok, err := p.checkFileExt(wrongFile2, ""); ok == true || err != nil {
		t.Error("Should be right")
	}
}

func TestFileContent(t *testing.T) {
	const (
		header  = "module.exports = {"
		content = `function yo(name) {
			return "yo " + name!";
		}`
		footer = "}"
	)

	var builder strings.Builder
	builder.WriteString(header)

	var funName string
	reg, _ := regexp.Compile(`^function\s[A-Za-z0-9]+\s*\([A-Za-z0-9\s,]*\)\s+\{`)
	if reg.MatchString(content) {
		funName = content[len("function"):len(content)]
		funName = strings.TrimSpace(funName)
	}
	newContent := funName
	builder.WriteString(fmt.Sprintf("\n%s\n", newContent))
	builder.WriteString(footer)
	ret := builder.String()

	if len(ret) <= 0 {
		t.Error("Does not generate new content")
	}
}

func TestManifestParser(t *testing.T) {
	parser, _ := NewParser("./", "")
	err := parser.parseMenfest("appmanifest.json")

	if err != nil {
		t.Error("Parase manifest error")
	}
}
