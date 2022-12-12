package dsl

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/antonioalfa22/egida/pkg/ansible"
	"github.com/antonioalfa22/egida/pkg/models"
	"github.com/antonioalfa22/egida/pkg/parser"
	"github.com/antonioalfa22/go-utils/collections"
	"reflect"
	"strings"
	"unicode/utf8"
)


func ParseFile(file string)  {
	listener := &ErrorListener{Errors: 0}
	input, _ := antlr.NewFileStream(file)
	// Lexer
	lexer := parser.NewAspidaLexer(input)
	lexer.RemoveErrorListeners()
	stream := antlr.NewCommonTokenStream(lexer,antlr.TokenDefaultChannel)
	// Parser
	p := parser.NewAspidaParser(stream)
	// Listeners
	p.RemoveErrorListeners()
	p.AddErrorListener(listener)
	p.BuildParseTrees = true
	tree := p.Program()
	if listener.Errors == 0 {
		// Codegen
		codegen := NewVisitor()
		codegen.Visit(tree)
		// Run Playbook
		tags := getTags(codegen.TkTable.Sections, codegen.TkTable.Points,
			codegen.TkTable.Controls, codegen.TkTable.Tags)
		vars := collections.Map(getVars(codegen.VarTable.Table), func(x string) string {
			return "    " + x + "\n"
		}).([]string)
		hosts := codegen.HostGroup
		connection := strings.ToLower(codegen.Info.Connection)

		// Check variables
		checkVariables(codegen.TkTable.Sections, codegen.TkTable.Points,
			codegen.TkTable.Controls, codegen.TkTable.Tags, vars)
		ansible.RunDSLPlaybook(tags, vars, hosts, connection)
	}
}

func checkVariables(sections []string, points []string, controls []string, tags []string, vars []string) {
	notvars := collections.Filter(models.GetVarsList(), func(x models.Variable) bool {
		for _, v := range vars { if v == x.Name { return false } }
		return true
	}).([]models.Variable)
	for _, variable := range notvars {
		inuse := variable.VariableInPoints(points) || variable.VariableInControls(controls) ||
			variable.VariableInSection(sections) || variable.VariableInTags(tags)
		if inuse {
			fmt.Println("WARNING: Variable ", variable.Name, "not defined", "--> Its default value will be used")
		}
	}
}

func getTags(sections []string, points []string, controls []string, tags []string) []string {
	var lines []string
	for _, sec := range sections {
		lines = append(lines, "section_"+sec)
	}
	for _, p := range points {
		lines = append(lines, "rule_"+p)
	}
	for _, c := range controls {
		lines = append(lines, "control_"+c)
	}
	for _, c := range tags {
		lines = append(lines, c)
	}
	return lines
}

func getVars(all map[string]interface{}) []string {
	var lines []string
	lines = getLinesFromObject(all, 0)
	return lines
}

func getLinesFromObject(all map[string]interface{}, tab int) []string {
	var lines []string
	for key, val := range all {
		tabulador := strings.Repeat("  ", tab)
		if reflect.ValueOf(val).Kind() == reflect.Map {
			lines = append(lines, tabulador + key + ":")
			tab = tab + 1
			lines = append(lines, getLinesFromObject(val.(map[string]interface{}), tab)...)
			tab = tab - 1
		} else {
			lines = append(lines, tabulador + key + ":" + " " + getYAMLformat(val))
		}
	}
	return lines
}

func getYAMLformat(value interface{}) string {
	switch reflect.ValueOf(value).Kind() {
	case reflect.String:
		return  "'"+ fmt.Sprintf("%v", value)+ "'"
	case reflect.Slice:
		line := "["
		for _, v := range value.([]interface{}) {
			line = line + getYAMLformat(v) + ","
		}
		line = trimLastChar(line)
		line = line + "]"
		return line
	default:
		return fmt.Sprintf("%v", value)
	}
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}


