package jinja_go_tests

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/joram/jinja-go"
)

func TemplateCompileTest(templatePath string, f os.FileInfo, err error) error {
	if !strings.HasSuffix(templatePath, ".html") {
		return nil
	}
	//if true {
	//	return nil
	//}
	fmt.Printf("checking compiled tree of: %s\n", templatePath)

	// arrange
	expectedTreePath := strings.Replace(templatePath, ".html", ".compile_tree.json", 1)
	expectedTree := readFileContent(expectedTreePath)

	content := readFileContent(templatePath)
	template := jinja_go.NewTemplate()
	template.Compile(content)

	// act
	tree, err := template.JSONTree()

	// assert
	if err != nil {
		return err
	}

	if tree != expectedTree {
		fmt.Printf("expected:\t%s\nactual:\t\t%s\n", expectedTree, tree)
		return errors.New("Failed to compile properly")
	}
	return nil
}

func Test(t *testing.T) {
	//println("testing compiling")
	//err := filepath.Walk("templates/", TemplateCompileTest)
	//if err != nil {
	//	t.Errorf("failed to compile template: ", err)
	//}
}
