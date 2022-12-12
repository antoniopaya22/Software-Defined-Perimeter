// Code generated from Aspida.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Aspida
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAspidaListener is a complete listener for a parse tree produced by AspidaParser.
type BaseAspidaListener struct{}

var _ AspidaListener = &BaseAspidaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAspidaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAspidaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAspidaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAspidaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseAspidaListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseAspidaListener) ExitProgram(ctx *ProgramContext) {}

// EnterMain is called when production main is entered.
func (s *BaseAspidaListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseAspidaListener) ExitMain(ctx *MainContext) {}

// EnterHosts is called when production hosts is entered.
func (s *BaseAspidaListener) EnterHosts(ctx *HostsContext) {}

// ExitHosts is called when production hosts is exited.
func (s *BaseAspidaListener) ExitHosts(ctx *HostsContext) {}

// EnterTasks is called when production tasks is entered.
func (s *BaseAspidaListener) EnterTasks(ctx *TasksContext) {}

// ExitTasks is called when production tasks is exited.
func (s *BaseAspidaListener) ExitTasks(ctx *TasksContext) {}

// EnterVariables is called when production variables is entered.
func (s *BaseAspidaListener) EnterVariables(ctx *VariablesContext) {}

// ExitVariables is called when production variables is exited.
func (s *BaseAspidaListener) ExitVariables(ctx *VariablesContext) {}

// EnterMain_content is called when production main_content is entered.
func (s *BaseAspidaListener) EnterMain_content(ctx *Main_contentContext) {}

// ExitMain_content is called when production main_content is exited.
func (s *BaseAspidaListener) ExitMain_content(ctx *Main_contentContext) {}

// EnterNameMain is called when production nameMain is entered.
func (s *BaseAspidaListener) EnterNameMain(ctx *NameMainContext) {}

// ExitNameMain is called when production nameMain is exited.
func (s *BaseAspidaListener) ExitNameMain(ctx *NameMainContext) {}

// EnterConnectionMain is called when production connectionMain is entered.
func (s *BaseAspidaListener) EnterConnectionMain(ctx *ConnectionMainContext) {}

// ExitConnectionMain is called when production connectionMain is exited.
func (s *BaseAspidaListener) ExitConnectionMain(ctx *ConnectionMainContext) {}

// EnterDescriptionMain is called when production descriptionMain is entered.
func (s *BaseAspidaListener) EnterDescriptionMain(ctx *DescriptionMainContext) {}

// ExitDescriptionMain is called when production descriptionMain is exited.
func (s *BaseAspidaListener) ExitDescriptionMain(ctx *DescriptionMainContext) {}

// EnterName is called when production name is entered.
func (s *BaseAspidaListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseAspidaListener) ExitName(ctx *NameContext) {}

// EnterConnection is called when production connection is entered.
func (s *BaseAspidaListener) EnterConnection(ctx *ConnectionContext) {}

// ExitConnection is called when production connection is exited.
func (s *BaseAspidaListener) ExitConnection(ctx *ConnectionContext) {}

// EnterConnectionLocal is called when production connectionLocal is entered.
func (s *BaseAspidaListener) EnterConnectionLocal(ctx *ConnectionLocalContext) {}

// ExitConnectionLocal is called when production connectionLocal is exited.
func (s *BaseAspidaListener) ExitConnectionLocal(ctx *ConnectionLocalContext) {}

// EnterConnectionSSH is called when production connectionSSH is entered.
func (s *BaseAspidaListener) EnterConnectionSSH(ctx *ConnectionSSHContext) {}

// ExitConnectionSSH is called when production connectionSSH is exited.
func (s *BaseAspidaListener) ExitConnectionSSH(ctx *ConnectionSSHContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseAspidaListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseAspidaListener) ExitDescription(ctx *DescriptionContext) {}

// EnterTContent is called when production tContent is entered.
func (s *BaseAspidaListener) EnterTContent(ctx *TContentContext) {}

// ExitTContent is called when production tContent is exited.
func (s *BaseAspidaListener) ExitTContent(ctx *TContentContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseAspidaListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseAspidaListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterTSections is called when production tSections is entered.
func (s *BaseAspidaListener) EnterTSections(ctx *TSectionsContext) {}

// ExitTSections is called when production tSections is exited.
func (s *BaseAspidaListener) ExitTSections(ctx *TSectionsContext) {}

// EnterTPoints is called when production tPoints is entered.
func (s *BaseAspidaListener) EnterTPoints(ctx *TPointsContext) {}

// ExitTPoints is called when production tPoints is exited.
func (s *BaseAspidaListener) ExitTPoints(ctx *TPointsContext) {}

// EnterTControls is called when production tControls is entered.
func (s *BaseAspidaListener) EnterTControls(ctx *TControlsContext) {}

// ExitTControls is called when production tControls is exited.
func (s *BaseAspidaListener) ExitTControls(ctx *TControlsContext) {}

// EnterTExclusions is called when production tExclusions is entered.
func (s *BaseAspidaListener) EnterTExclusions(ctx *TExclusionsContext) {}

// ExitTExclusions is called when production tExclusions is exited.
func (s *BaseAspidaListener) ExitTExclusions(ctx *TExclusionsContext) {}

// EnterTTags is called when production tTags is entered.
func (s *BaseAspidaListener) EnterTTags(ctx *TTagsContext) {}

// ExitTTags is called when production tTags is exited.
func (s *BaseAspidaListener) ExitTTags(ctx *TTagsContext) {}

// EnterSections is called when production sections is entered.
func (s *BaseAspidaListener) EnterSections(ctx *SectionsContext) {}

// ExitSections is called when production sections is exited.
func (s *BaseAspidaListener) ExitSections(ctx *SectionsContext) {}

// EnterPoints is called when production points is entered.
func (s *BaseAspidaListener) EnterPoints(ctx *PointsContext) {}

// ExitPoints is called when production points is exited.
func (s *BaseAspidaListener) ExitPoints(ctx *PointsContext) {}

// EnterControls is called when production controls is entered.
func (s *BaseAspidaListener) EnterControls(ctx *ControlsContext) {}

// ExitControls is called when production controls is exited.
func (s *BaseAspidaListener) ExitControls(ctx *ControlsContext) {}

// EnterExclusions is called when production exclusions is entered.
func (s *BaseAspidaListener) EnterExclusions(ctx *ExclusionsContext) {}

// ExitExclusions is called when production exclusions is exited.
func (s *BaseAspidaListener) ExitExclusions(ctx *ExclusionsContext) {}

// EnterTags is called when production tags is entered.
func (s *BaseAspidaListener) EnterTags(ctx *TagsContext) {}

// ExitTags is called when production tags is exited.
func (s *BaseAspidaListener) ExitTags(ctx *TagsContext) {}

// EnterIfStat is called when production ifStat is entered.
func (s *BaseAspidaListener) EnterIfStat(ctx *IfStatContext) {}

// ExitIfStat is called when production ifStat is exited.
func (s *BaseAspidaListener) ExitIfStat(ctx *IfStatContext) {}

// EnterElifStat is called when production elifStat is entered.
func (s *BaseAspidaListener) EnterElifStat(ctx *ElifStatContext) {}

// ExitElifStat is called when production elifStat is exited.
func (s *BaseAspidaListener) ExitElifStat(ctx *ElifStatContext) {}

// EnterElseStat is called when production elseStat is entered.
func (s *BaseAspidaListener) EnterElseStat(ctx *ElseStatContext) {}

// ExitElseStat is called when production elseStat is exited.
func (s *BaseAspidaListener) ExitElseStat(ctx *ElseStatContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseAspidaListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseAspidaListener) ExitComparison(ctx *ComparisonContext) {}

// EnterVars_content is called when production vars_content is entered.
func (s *BaseAspidaListener) EnterVars_content(ctx *Vars_contentContext) {}

// ExitVars_content is called when production vars_content is exited.
func (s *BaseAspidaListener) ExitVars_content(ctx *Vars_contentContext) {}

// EnterVarDef is called when production varDef is entered.
func (s *BaseAspidaListener) EnterVarDef(ctx *VarDefContext) {}

// ExitVarDef is called when production varDef is exited.
func (s *BaseAspidaListener) ExitVarDef(ctx *VarDefContext) {}

// EnterVarObjDef is called when production varObjDef is entered.
func (s *BaseAspidaListener) EnterVarObjDef(ctx *VarObjDefContext) {}

// ExitVarObjDef is called when production varObjDef is exited.
func (s *BaseAspidaListener) ExitVarObjDef(ctx *VarObjDefContext) {}

// EnterComp_op is called when production comp_op is entered.
func (s *BaseAspidaListener) EnterComp_op(ctx *Comp_opContext) {}

// ExitComp_op is called when production comp_op is exited.
func (s *BaseAspidaListener) ExitComp_op(ctx *Comp_opContext) {}

// EnterStringVal is called when production StringVal is entered.
func (s *BaseAspidaListener) EnterStringVal(ctx *StringValContext) {}

// ExitStringVal is called when production StringVal is exited.
func (s *BaseAspidaListener) ExitStringVal(ctx *StringValContext) {}

// EnterNumberVal is called when production NumberVal is entered.
func (s *BaseAspidaListener) EnterNumberVal(ctx *NumberValContext) {}

// ExitNumberVal is called when production NumberVal is exited.
func (s *BaseAspidaListener) ExitNumberVal(ctx *NumberValContext) {}

// EnterTrueVal is called when production TrueVal is entered.
func (s *BaseAspidaListener) EnterTrueVal(ctx *TrueValContext) {}

// ExitTrueVal is called when production TrueVal is exited.
func (s *BaseAspidaListener) ExitTrueVal(ctx *TrueValContext) {}

// EnterFalseVal is called when production FalseVal is entered.
func (s *BaseAspidaListener) EnterFalseVal(ctx *FalseValContext) {}

// ExitFalseVal is called when production FalseVal is exited.
func (s *BaseAspidaListener) ExitFalseVal(ctx *FalseValContext) {}

// EnterNullVal is called when production NullVal is entered.
func (s *BaseAspidaListener) EnterNullVal(ctx *NullValContext) {}

// ExitNullVal is called when production NullVal is exited.
func (s *BaseAspidaListener) ExitNullVal(ctx *NullValContext) {}

// EnterArrayVal is called when production ArrayVal is entered.
func (s *BaseAspidaListener) EnterArrayVal(ctx *ArrayValContext) {}

// ExitArrayVal is called when production ArrayVal is exited.
func (s *BaseAspidaListener) ExitArrayVal(ctx *ArrayValContext) {}

// EnterStr_array is called when production str_array is entered.
func (s *BaseAspidaListener) EnterStr_array(ctx *Str_arrayContext) {}

// ExitStr_array is called when production str_array is exited.
func (s *BaseAspidaListener) ExitStr_array(ctx *Str_arrayContext) {}

// EnterArray is called when production array is entered.
func (s *BaseAspidaListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseAspidaListener) ExitArray(ctx *ArrayContext) {}
