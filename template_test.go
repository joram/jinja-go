package jinja_go

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func GetPythonOutput(path string) (string, error) {
	contextPath := strings.Replace(path, ".html", "_context.json", 1)
	cmd := exec.Command("tests/bin/render_jinja.py", path, contextPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func TemplateCompileTest(path string, f os.FileInfo, err error) error {
	if !strings.HasSuffix(path, ".html") {
		return nil
	}
	fmt.Printf("test compiling %s\n", path)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	content := string(bytes)
	template := NewTemplate()
	template.Compile(content)

	treeString, _ := template.JSONTree()
	println(treeString)

	// TODO: test against path_tree.json loaded as string
	return nil
}

func TemplateRenderTest(path string, f os.FileInfo, err error) error {
	if !strings.HasSuffix(path, ".html") {
		return nil
	}
	fmt.Printf("test rendering %s\n", path)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	content := string(bytes)
	template := NewTemplate()
	template.Compile(content)

	s, err := GetPythonOutput(path)
	// TODO: compare against template.Render(context)

	println("output: " + s)
	if err != nil {
		return err
	}
	return nil
}

func TestCompile(t *testing.T) {
	err := filepath.Walk("tests/templates/", TemplateCompileTest)
	if err != nil {
		t.Errorf("failed to compile template: ", err)
	}
}

func TestRender(t *testing.T) {
	err := filepath.Walk("tests/templates/", TemplateRenderTest)
	if err != nil {
		t.Errorf("failed to compile template: ", err)
	}
}
