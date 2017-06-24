package jinja_go_tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func readFileContent(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	return content

}

func readFileContentAsBytes(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes
}

func getNumberedPaths(pathTemplate, defaultContent string) []string {
	contextPaths := []string{}
	i := 0
	for {
		contextPath := fmt.Sprintf(pathTemplate, i)
		_, err := os.Stat(contextPath)
		if os.IsNotExist(err) {
			break
		}
		contextPaths = append(contextPaths, contextPath)
		i += 1
	}

	if len(contextPaths) == 0 {
		contextPath := fmt.Sprintf(pathTemplate, 0)
		contextDir := path.Dir(contextPath)
		os.MkdirAll(contextDir, 0755)
		ioutil.WriteFile(contextPath, []byte(defaultContent), 0755)
		contextPaths = append(contextPaths, contextPath)
	}

	return contextPaths
}
