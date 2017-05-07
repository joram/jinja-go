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
	println("")

	t3 := "initial text, {{ your_name }} final text."
	template3 := Template{NewDefaultConfig(), "", RootNode{}, []INode{}}
	template3.Compile(t3)
	println("")

	t4 := "initial text, {# you smell bad #} final text."
	template4 := Template{NewDefaultConfig(), "", RootNode{}, []INode{}}
	template4.Compile(t4)
	println("")

}
