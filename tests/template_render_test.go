package jinja_go_tests

import (
	"bytes"
	"fmt"
	"github.com/joram/jinja-go"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	err := filepath.Walk("tests/templates/", TemplateRenderTest)
	if err != nil {
		t.Errorf("failed to compile template: ", err)
	}
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
	template := jinja_go.NewTemplate()
	template.Compile(content)

	s, err := GetPythonOutput(path)
	// TODO: compare against template.Render(context)

	println("output: " + s)
	if err != nil {
		return err
	}
	return nil
}

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
