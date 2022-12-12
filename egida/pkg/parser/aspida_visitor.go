// Code generated from Aspida.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Aspida
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by AspidaParser.
type AspidaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AspidaParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by AspidaParser#main.
	VisitMain(ctx *MainContext) interface{}

	// Visit a parse tree produced by AspidaParser#hosts.
	VisitHosts(ctx *HostsContext) interface{}

	// Visit a parse tree produced by AspidaParser#tasks.
	VisitTasks(ctx *TasksContext) interface{}

	// Visit a parse tree produced by AspidaParser#variables.
	VisitVariables(ctx *VariablesContext) interface{}

	// Visit a parse tree produced by AspidaParser#main_content.
	VisitMain_content(ctx *Main_contentContext) interface{}

	// Visit a parse tree produced by AspidaParser#nameMain.
	VisitNameMain(ctx *NameMainContext) interface{}

	// Visit a parse tree produced by AspidaParser#connectionMain.
	VisitConnectionMain(ctx *ConnectionMainContext) interface{}

	// Visit a parse tree produced by AspidaParser#descriptionMain.
	VisitDescriptionMain(ctx *DescriptionMainContext) interface{}

	// Visit a parse tree produced by AspidaParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by AspidaParser#connection.
	VisitConnection(ctx *ConnectionContext) interface{}

	// Visit a parse tree produced by AspidaParser#connectionLocal.
	VisitConnectionLocal(ctx *ConnectionLocalContext) interface{}

	// Visit a parse tree produced by AspidaParser#connectionSSH.
	VisitConnectionSSH(ctx *ConnectionSSHContext) interface{}

	// Visit a parse tree produced by AspidaParser#description.
	VisitDescription(ctx *DescriptionContext) interface{}

	// Visit a parse tree produced by AspidaParser#tContent.
	VisitTContent(ctx *TContentContext) interface{}

	// Visit a parse tree produced by AspidaParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by AspidaParser#tSections.
	VisitTSections(ctx *TSectionsContext) interface{}

	// Visit a parse tree produced by AspidaParser#tPoints.
	VisitTPoints(ctx *TPointsContext) interface{}

	// Visit a parse tree produced by AspidaParser#tControls.
	VisitTControls(ctx *TControlsContext) interface{}

	// Visit a parse tree produced by AspidaParser#tExclusions.
	VisitTExclusions(ctx *TExclusionsContext) interface{}

	// Visit a parse tree produced by AspidaParser#tTags.
	VisitTTags(ctx *TTagsContext) interface{}

	// Visit a parse tree produced by AspidaParser#sections.
	VisitSections(ctx *SectionsContext) interface{}

	// Visit a parse tree produced by AspidaParser#points.
	VisitPoints(ctx *PointsContext) interface{}

	// Visit a parse tree produced by AspidaParser#controls.
	VisitControls(ctx *ControlsContext) interface{}

	// Visit a parse tree produced by AspidaParser#exclusions.
	VisitExclusions(ctx *ExclusionsContext) interface{}

	// Visit a parse tree produced by AspidaParser#tags.
	VisitTags(ctx *TagsContext) interface{}

	// Visit a parse tree produced by AspidaParser#ifStat.
	VisitIfStat(ctx *IfStatContext) interface{}

	// Visit a parse tree produced by AspidaParser#elifStat.
	VisitElifStat(ctx *ElifStatContext) interface{}

	// Visit a parse tree produced by AspidaParser#elseStat.
	VisitElseStat(ctx *ElseStatContext) interface{}

	// Visit a parse tree produced by AspidaParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by AspidaParser#vars_content.
	VisitVars_content(ctx *Vars_contentContext) interface{}

	// Visit a parse tree produced by AspidaParser#varDef.
	VisitVarDef(ctx *VarDefContext) interface{}

	// Visit a parse tree produced by AspidaParser#varObjDef.
	VisitVarObjDef(ctx *VarObjDefContext) interface{}

	// Visit a parse tree produced by AspidaParser#comp_op.
	VisitComp_op(ctx *Comp_opContext) interface{}

	// Visit a parse tree produced by AspidaParser#StringVal.
	VisitStringVal(ctx *StringValContext) interface{}

	// Visit a parse tree produced by AspidaParser#NumberVal.
	VisitNumberVal(ctx *NumberValContext) interface{}

	// Visit a parse tree produced by AspidaParser#TrueVal.
	VisitTrueVal(ctx *TrueValContext) interface{}

	// Visit a parse tree produced by AspidaParser#FalseVal.
	VisitFalseVal(ctx *FalseValContext) interface{}

	// Visit a parse tree produced by AspidaParser#NullVal.
	VisitNullVal(ctx *NullValContext) interface{}

	// Visit a parse tree produced by AspidaParser#ArrayVal.
	VisitArrayVal(ctx *ArrayValContext) interface{}

	// Visit a parse tree produced by AspidaParser#str_array.
	VisitStr_array(ctx *Str_arrayContext) interface{}

	// Visit a parse tree produced by AspidaParser#array.
	VisitArray(ctx *ArrayContext) interface{}
}
