package models

type Variable struct {
	Name string
	Section string
	Points []string
	Tags []string
	Controls []string
}

func (v Variable) VariableInPoints(points []string) bool {
	for _, c := range points {
		if v.Contains(v.Points, "rule_"+c) {
			return true
		}
	}
	return false
}

func (v Variable) VariableInTags(tags []string) bool {
	for _, c := range tags {
		if v.Contains(v.Tags, c) {
			return true
		}
	}
	return false
}

func (v Variable) VariableInControls(controls []string) bool {
	for _, c := range controls {
		if v.Contains(v.Controls, "control_"+c) {
			return true
		}
	}
	return false
}

func (v Variable) VariableInSection(sections []string) bool {
	for _, c := range sections {
		if v.Section == "section_"+c {
			return true
		}
	}
	return false
}

func (v Variable) Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}