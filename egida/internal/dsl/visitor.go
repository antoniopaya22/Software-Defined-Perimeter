package dsl

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/antonioalfa22/egida/pkg/parser"
	"github.com/antonioalfa22/go-utils/collections"
	"strings"
)

type Task struct {
	Type string
	Tasks interface{}
}

type InfoData struct {
	Name string
	Connection string
	Description string
}

type Visitor struct{
	*antlr.BaseParseTreeVisitor
	VarTable VariablesTable
	TkTable TasksTable
	HostGroup string
	Info	InfoData
}

func NewVisitor() Visitor {
	return Visitor{
		VarTable: NewVariablesTable(),
		TkTable: NewTasksTable(),
		Info: InfoData{},
	}
}

func (v *Visitor) Visit(node antlr.ParseTree) interface{} {
	res := node.Accept(v)
	return res
}


func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	ctx.Main().Accept(v)
	v.HostGroup = ctx.Hosts().Accept(v).(string)
	ctx.Tasks().Accept(v)
	if ctx.Variables() != nil {
		ctx.Variables().Accept(v)
	}
	return nil
}

func (v *Visitor) VisitMain(ctx *parser.MainContext) interface{} {
	ctx.Main_content().Accept(v)
	return nil
}

func (v *Visitor) VisitHosts(ctx *parser.HostsContext) interface{} {
	return strings.Replace(ctx.STRING().GetText(), "\"", "", -1)
}

func (v *Visitor) VisitMain_content(ctx *parser.Main_contentContext) interface{} {
	for _, mp := range ctx.AllMain_prop() {
		mp.Accept(v)
	}
	return nil
}

func (v *Visitor) VisitNameMain(ctx *parser.NameMainContext) interface{} {
	v.Info.Name = ctx.Name().Accept(v).(string)
	return nil
}

func (v *Visitor) VisitConnectionMain(ctx *parser.ConnectionMainContext) interface{} {
	v.Info.Connection = ctx.Connection().Accept(v).(string)
	return nil
}

func (v *Visitor) VisitDescriptionMain(ctx *parser.DescriptionMainContext) interface{} {
	v.Info.Description = ctx.Description().Accept(v).(string)
	return nil
}

func (v *Visitor) VisitName(ctx *parser.NameContext) interface{} {
	return ctx.Value().Accept(v)
}

func (v *Visitor) VisitConnection(ctx *parser.ConnectionContext) interface{} {
	return ctx.Connection_type().Accept(v)
}

func (v *Visitor) VisitConnectionLocal(ctx *parser.ConnectionLocalContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitConnectionSSH(ctx *parser.ConnectionSSHContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitDescription(ctx *parser.DescriptionContext) interface{} {
	return ctx.Value().Accept(v)
}

func (v *Visitor) VisitTasks(ctx *parser.TasksContext) interface{} {
	ctx.Tasks_content().Accept(v)
	return nil
}

func (v *Visitor) VisitTContent(ctx *parser.TContentContext) interface{} {
	for _, c := range ctx.AllTasks_prop() {
		task := c.Accept(v).(Task)
		switch task.Type {
		case "sections":
			collections.ForEach(task.Tasks, func(x string) {v.TkTable.AddSection(x)})
		case "controls":
			collections.ForEach(task.Tasks, func(x string) {v.TkTable.AddControl(x)})
		case "points":
			collections.ForEach(task.Tasks, func(x string) {v.TkTable.AddPoint(x)})
		case "exclusions":
			collections.ForEach(task.Tasks, func(x string) {v.TkTable.AddExclusionPoint(x)})
		case "tags":
			collections.ForEach(task.Tasks, func(x string) {v.TkTable.AddTag(x)})
		}
	}
	return nil
}

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	temp := ctx.IfStat().Accept(v).(bool)
	for _, ei := range ctx.AllElifStat() {
		if !temp{
			temp = ei.Accept(v).(bool)
		}
	}
	if !temp {
		ctx.ElseStat().Accept(v)
	}
	return nil
}

func (v *Visitor) VisitTSections(ctx *parser.TSectionsContext) interface{} {
	return Task{Tasks: ctx.Sections().Accept(v), Type: "sections"}
}

func (v *Visitor) VisitTPoints(ctx *parser.TPointsContext) interface{} {
	return Task{Tasks: ctx.Points().Accept(v), Type: "points"}
}

func (v *Visitor) VisitTControls(ctx *parser.TControlsContext) interface{} {
	return Task{Tasks: ctx.Controls().Accept(v), Type: "controls"}
}

func (v *Visitor) VisitTExclusions(ctx *parser.TExclusionsContext) interface{} {
	return Task{Tasks: ctx.Exclusions().Accept(v), Type: "exclusions"}
}

func (v *Visitor) VisitTTags(ctx *parser.TTagsContext) interface{} {
	return Task{Tasks: ctx.Tags().Accept(v), Type: "tags"}
}

func (v *Visitor) VisitSections(ctx *parser.SectionsContext) interface{} {
	return ctx.Str_array().Accept(v)
}

func (v *Visitor) VisitPoints(ctx *parser.PointsContext) interface{} {
	return ctx.Str_array().Accept(v)
}

func (v *Visitor) VisitControls(ctx *parser.ControlsContext) interface{} {
	return ctx.Str_array().Accept(v)
}

func (v *Visitor) VisitExclusions(ctx *parser.ExclusionsContext) interface{} {
	return ctx.Str_array().Accept(v)
}

func (v *Visitor) VisitTags(ctx *parser.TagsContext) interface{} {
	return ctx.Str_array().Accept(v)
}

func (v *Visitor) VisitIfStat(ctx *parser.IfStatContext) interface{} {
	com := ctx.Comparison().Accept(v).(bool)
	if com {
		ctx.Tasks_content().Accept(v)
		return true
	}
	return false
}

func (v *Visitor) VisitElifStat(ctx *parser.ElifStatContext) interface{} {
	com := ctx.Comparison().Accept(v).(bool)
	if com {
		ctx.Tasks_content().Accept(v)
		return true
	}
	return false
}

func (v *Visitor) VisitElseStat(ctx *parser.ElseStatContext) interface{} {
	ctx.Tasks_content().Accept(v)
	return nil
}

func (v *Visitor) VisitComparison(ctx *parser.ComparisonContext) interface{} {
	com := Comparation{
		Value1: ctx.Value(0).Accept(v),
		Operator: ctx.Comp_op().Accept(v).(string),
		Value2: ctx.Value(1).Accept(v),
		Host: v.HostGroup,
	}
	return com.Compare()
}

// =========================================== VARIABLES

func (v *Visitor) VisitVariables(ctx *parser.VariablesContext) interface{} {
	ctx.Vars_content().Accept(v)
	return nil
}

func (v *Visitor) VisitVars_content(ctx *parser.Vars_contentContext) interface{} {
	for _, c := range ctx.AllVars_prop() {
		c.Accept(v)
	}
	return nil
}

func (v *Visitor) VisitVarDef(ctx *parser.VarDefContext) interface{} {
	val := ctx.Value().Accept(v)
	v.VarTable.NewVariable(strings.Replace(ctx.STRING().GetText(), "\"", "", -1), val)
	return nil
}

func (v *Visitor) VisitVarObjDef(ctx *parser.VarObjDefContext) interface{} {
	v.VarTable.NewObject(strings.Replace(ctx.STRING().GetText(), "\"", "", -1))
	ctx.Vars_content().Accept(v)
	v.VarTable.EndObject(strings.Replace(ctx.STRING().GetText(), "\"", "", -1))
	return nil
}

// =========================================== VALUES

func (v *Visitor) VisitComp_op(ctx *parser.Comp_opContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitArrayVal(ctx *parser.ArrayValContext) interface{} {
	value := ctx.Array().Accept(v)
	return value
}

func (v *Visitor) VisitArray(ctx *parser.ArrayContext) interface{} {
	var temp []interface{}
	for _, c := range ctx.AllValue() {
		temp = append(temp, c.Accept(v))
	}
	return temp
}

func (v *Visitor) VisitStringVal(ctx *parser.StringValContext) interface{} {
	return strings.Replace(ctx.STRING().GetText(), "\"", "", -1)
}

func (v *Visitor) VisitNumberVal(ctx *parser.NumberValContext) interface{} {
	return ctx.NUMBER().GetText()
}

func (v *Visitor) VisitTrueVal(ctx *parser.TrueValContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitFalseVal(ctx *parser.FalseValContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitNullVal(ctx *parser.NullValContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitStr_array(ctx *parser.Str_arrayContext) interface{} {
	var temp []string
	for _, s := range ctx.AllSTRING(){
		temp = append(temp, strings.Replace(s.GetText(), "\"", "", -1))
	}
	return temp
}
