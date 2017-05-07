package jinja_go

import (
	"testing"
)

func TestCalcDesiredNumDynos(t *testing.T) {

	t1 := "initial text, {% if foo == bar %}FOO{% else %}BAR{% endif %} final text."
	template := Template{NewDefaultConfig(), "", RootNode{}, []INode{}}
	template.Compile(t1)
	println("")

	t2 := "initial text, {% if foo == bar %}FOO{% endif %} final text."
	template2 := Template{NewDefaultConfig(), "", RootNode{}, []INode{}}
	template2.Compile(t2)

}
