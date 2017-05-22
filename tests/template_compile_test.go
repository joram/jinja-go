package jinja_go_tests

import (
	"errors"
	"fmt"
	"github.com/joram/jinja-go"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TemplateCompileTest(path string, f os.FileInfo, err error) error {
	fmt.Printf("checking compiled tree of: %s\n", path)
	if !strings.HasSuffix(path, ".html") {
		return nil
	}

	content := readFileContent(path)
	template := jinja_go.NewTemplate()
	expectedTreePath := strings.Replace(path, "templates", "compile_tree", 1)
	expectedTreePath = strings.Replace(expectedTreePath, ".html", ".json", 1)
	expectedTree := readFileContent(expectedTreePath)
	template.Compile(content)
	tree, err := template.JSONTree()

	// act
	treeString, _ := template.JSONTree()

	if err != nil {
		return err
	}
	if tree != expectedTree {
		fmt.Printf("expected:\t%s\nactual:\t\t%s\n", expectedTree, treeString)
		return errors.New("Failed to compile properly")
	}
	return nil
}

func Test(t *testing.T) {
	println("testing compiling")
	err := filepath.Walk("templates/", TemplateCompileTest)
	if err != nil {
		t.Errorf("failed to compile template: ", err)
	}
}
