package jinja_go_tests

import "io/ioutil"

func readFileContent(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	return content

}
