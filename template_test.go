package jinja_go

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func visit(path string, f os.FileInfo, err error) error {
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
	return nil
}

func TestCalcDesiredNumDynos(t *testing.T) {
	err := filepath.Walk("tests/templates/", visit)
	if err != nil {
		t.Errorf("failed to compile template: ", err)
	}
}
