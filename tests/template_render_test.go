package jinja_go_tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joram/jinja-go"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	err := filepath.Walk("templates/", TemplateRenderTest)
	if err != nil {
		t.Errorf("failed to compile template: ", err)
	}
}

func TemplateRenderTest(path string, f os.FileInfo, err error) error {
	if !strings.HasSuffix(path, ".html") {
		return nil
	}
	fmt.Printf("checking rendering of: %s\n", path)
	content := readFileContent(path)
	template := jinja_go.NewTemplate()
	template.Compile(content)
	context, contextPath := getContext(path)
	expectedCompiledString, pyErr := GetPythonOutput(path, contextPath)

	// act
	compiledString := template.Render(context)

	if pyErr != nil {
		return pyErr
	}
	if expectedCompiledString != compiledString {
		fmt.Printf("expected:\t`%s`\nactual:\t\t`%s`\n", expectedCompiledString, compiledString)
		return errors.New("did not compile the same as python")
	}
	return nil
}

func GetPythonOutput(path, contextPath string) (string, error) {
	cmd := exec.Command("bin/render_jinja.py", path, contextPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func getContext(templatePath string) (map[string]interface{}, string) {
	contextPath := strings.Replace(templatePath, "templates", "context", 1)
	contextPath = strings.Replace(contextPath, ".html", ".json", 1)
	contextContent := readFileContentAsBytes(contextPath)

	context := map[string]interface{}{}
	err := json.Unmarshal(contextContent, &context)
	if err != nil {
		panic(err)
	}
	return context, contextPath
}
