package jinja_go_tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joram/jinja-go"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	err := filepath.Walk("templates/", TemplateRenderTest)
	if err != nil {
		t.Errorf("failed to render template: ", err)
	}
}

func TemplateRenderTest(path string, f os.FileInfo, err error) error {

	if !strings.HasSuffix(path, ".html") {
		return nil
	}
	if strings.HasPrefix(path, "templates/sendwithus/") {
		return nil
	}

	template := jinja_go.NewTemplate()
	template.Compile(readFileContent(path))
	for i, contextPath := range getContextPaths(path) {
		fmt.Printf("checking rendering of %v: %s\n", i, path)

		// Arrange
		expectedCompiledString, pyErr := GetPythonOutput(path, contextPath)
		if pyErr != nil {
			fmt.Printf("Python error:\n%s\n%s\n", expectedCompiledString, pyErr)
			return pyErr
		}
		context := getContext(contextPath)

		// act
		compiledString := template.Render(context)

		// Assert
		if !compareResults(expectedCompiledString, compiledString) {
			//fmt.Printf("expected:\n`%s`\nactual:\n`%s`\n", expectedCompiledString, compiledString)
			return errors.New("did not compile the same as python")
		}

	}
	return nil
}

func compareResults(s1, s2 string) bool {

	if s1 == s2 {
		return true
	}

	s1 = strings.Replace(s1, "\r", "", -1)
	s2 = strings.Replace(s2, "\r", "", -1)
	if s1 == s2 {
		println("warning different newlines!")
		return true
	}

	for i, _ := range s1 {
		if s1[i] == s2[i] {
			continue
		}
		min := int(math.Max(0, float64(i-10)))
		max1 := int(math.Min(float64(len(s1)-1), float64(i+100)))
		max2 := int(math.Min(float64(len(s2)-1), float64(i+100)))

		fmt.Printf("difference at:\n-----------------------\n%v\n--------------------\n%v\n\n", s1[min:max1], s2[min:max2])
		break
	}

	return false
}

func GetPythonOutput(templatePath, contextPath string) (string, error) {
	cmd := exec.Command("bin/render_jinja.py", templatePath, contextPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("bin/render_jinja.py %v %v\n", templatePath, contextPath)
		return out.String(), err
	}
	return out.String(), nil
}

func getContextPaths(templatePath string) []string {
	contextPath := strings.Replace(templatePath, ".html", ".%v.context.json", 1)
	paths := getNumberedPaths(contextPath, "{}")
	return paths
}

func getContext(contextPath string) map[string]interface{} {
	contextContent := readFileContentAsBytes(contextPath)
	context := map[string]interface{}{}
	err := json.Unmarshal(contextContent, &context)
	if err != nil {
		panic(err)
	}
	return context
}
