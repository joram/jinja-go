package jinja_go

import (
	"testing"
)

func TestCalcDesiredNumDynos(t *testing.T) {

	t1 := "initial text, {% if foo == bar %}FOO{% else %}BAR{% endif %} final text."
	template := Template{NewDefaultConfig(), "", RootNode{}, []INode{}}
	template.Compile(t1)
	//if template.rootSection == nil {
	//	t.Errorf("Didn't compile properly")
	//}
	//if template.rootSection.startIndex != 0 {
	//	t.Errorf("%d", template.rootSection.startIndex)
	//}
	//if template.rootSection.EndIndex == 0 {
	//	t.Errorf("%d", template.rootSection.EndIndex)
	//}

}
