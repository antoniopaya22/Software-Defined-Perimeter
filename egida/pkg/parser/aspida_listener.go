// Code generated from Aspida.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Aspida
import "github.com/antlr/antlr4/runtime/Go/antlr"

// AspidaListener is a complete listener for a parse tree produced by AspidaParser.
type AspidaListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterHosts is called when entering the hosts production.
	EnterHosts(c *HostsContext)

	// EnterTasks is called when entering the tasks production.
	EnterTasks(c *TasksContext)

	// EnterVariables is called when entering the variables production.
	EnterVariables(c *VariablesContext)

	// EnterMain_content is called when entering the main_content production.
	EnterMain_content(c *Main_contentContext)

	// EnterNameMain is called when entering the nameMain production.
	EnterNameMain(c *NameMainContext)

	// EnterConnectionMain is called when entering the connectionMain production.
	EnterConnectionMain(c *ConnectionMainContext)

	// EnterDescriptionMain is called when entering the descriptionMain production.
	EnterDescriptionMain(c *DescriptionMainContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterConnection is called when entering the connection production.
	EnterConnection(c *ConnectionContext)

	// EnterConnectionLocal is called when entering the connectionLocal production.
	EnterConnectionLocal(c *ConnectionLocalContext)

	// EnterConnectionSSH is called when entering the connectionSSH production.
	EnterConnectionSSH(c *ConnectionSSHContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterTContent is called when entering the tContent production.
	EnterTContent(c *TContentContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterTSections is called when entering the tSections production.
	EnterTSections(c *TSectionsContext)

	// EnterTPoints is called when entering the tPoints production.
	EnterTPoints(c *TPointsContext)

	// EnterTControls is called when entering the tControls production.
	EnterTControls(c *TControlsContext)

	// EnterTExclusions is called when entering the tExclusions production.
	EnterTExclusions(c *TExclusionsContext)

	// EnterTTags is called when entering the tTags production.
	EnterTTags(c *TTagsContext)

	// EnterSections is called when entering the sections production.
	EnterSections(c *SectionsContext)

	// EnterPoints is called when entering the points production.
	EnterPoints(c *PointsContext)

	// EnterControls is called when entering the controls production.
	EnterControls(c *ControlsContext)

	// EnterExclusions is called when entering the exclusions production.
	EnterExclusions(c *ExclusionsContext)

	// EnterTags is called when entering the tags production.
	EnterTags(c *TagsContext)

	// EnterIfStat is called when entering the ifStat production.
	EnterIfStat(c *IfStatContext)

	// EnterElifStat is called when entering the elifStat production.
	EnterElifStat(c *ElifStatContext)

	// EnterElseStat is called when entering the elseStat production.
	EnterElseStat(c *ElseStatContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterVars_content is called when entering the vars_content production.
	EnterVars_content(c *Vars_contentContext)

	// EnterVarDef is called when entering the varDef production.
	EnterVarDef(c *VarDefContext)

	// EnterVarObjDef is called when entering the varObjDef production.
	EnterVarObjDef(c *VarObjDefContext)

	// EnterComp_op is called when entering the comp_op production.
	EnterComp_op(c *Comp_opContext)

	// EnterStringVal is called when entering the StringVal production.
	EnterStringVal(c *StringValContext)

	// EnterNumberVal is called when entering the NumberVal production.
	EnterNumberVal(c *NumberValContext)

	// EnterTrueVal is called when entering the TrueVal production.
	EnterTrueVal(c *TrueValContext)

	// EnterFalseVal is called when entering the FalseVal production.
	EnterFalseVal(c *FalseValContext)

	// EnterNullVal is called when entering the NullVal production.
	EnterNullVal(c *NullValContext)

	// EnterArrayVal is called when entering the ArrayVal production.
	EnterArrayVal(c *ArrayValContext)

	// EnterStr_array is called when entering the str_array production.
	EnterStr_array(c *Str_arrayContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitHosts is called when exiting the hosts production.
	ExitHosts(c *HostsContext)

	// ExitTasks is called when exiting the tasks production.
	ExitTasks(c *TasksContext)

	// ExitVariables is called when exiting the variables production.
	ExitVariables(c *VariablesContext)

	// ExitMain_content is called when exiting the main_content production.
	ExitMain_content(c *Main_contentContext)

	// ExitNameMain is called when exiting the nameMain production.
	ExitNameMain(c *NameMainContext)

	// ExitConnectionMain is called when exiting the connectionMain production.
	ExitConnectionMain(c *ConnectionMainContext)

	// ExitDescriptionMain is called when exiting the descriptionMain production.
	ExitDescriptionMain(c *DescriptionMainContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitConnection is called when exiting the connection production.
	ExitConnection(c *ConnectionContext)

	// ExitConnectionLocal is called when exiting the connectionLocal production.
	ExitConnectionLocal(c *ConnectionLocalContext)

	// ExitConnectionSSH is called when exiting the connectionSSH production.
	ExitConnectionSSH(c *ConnectionSSHContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitTContent is called when exiting the tContent production.
	ExitTContent(c *TContentContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitTSections is called when exiting the tSections production.
	ExitTSections(c *TSectionsContext)

	// ExitTPoints is called when exiting the tPoints production.
	ExitTPoints(c *TPointsContext)

	// ExitTControls is called when exiting the tControls production.
	ExitTControls(c *TControlsContext)

	// ExitTExclusions is called when exiting the tExclusions production.
	ExitTExclusions(c *TExclusionsContext)

	// ExitTTags is called when exiting the tTags production.
	ExitTTags(c *TTagsContext)

	// ExitSections is called when exiting the sections production.
	ExitSections(c *SectionsContext)

	// ExitPoints is called when exiting the points production.
	ExitPoints(c *PointsContext)

	// ExitControls is called when exiting the controls production.
	ExitControls(c *ControlsContext)

	// ExitExclusions is called when exiting the exclusions production.
	ExitExclusions(c *ExclusionsContext)

	// ExitTags is called when exiting the tags production.
	ExitTags(c *TagsContext)

	// ExitIfStat is called when exiting the ifStat production.
	ExitIfStat(c *IfStatContext)

	// ExitElifStat is called when exiting the elifStat production.
	ExitElifStat(c *ElifStatContext)

	// ExitElseStat is called when exiting the elseStat production.
	ExitElseStat(c *ElseStatContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitVars_content is called when exiting the vars_content production.
	ExitVars_content(c *Vars_contentContext)

	// ExitVarDef is called when exiting the varDef production.
	ExitVarDef(c *VarDefContext)

	// ExitVarObjDef is called when exiting the varObjDef production.
	ExitVarObjDef(c *VarObjDefContext)

	// ExitComp_op is called when exiting the comp_op production.
	ExitComp_op(c *Comp_opContext)

	// ExitStringVal is called when exiting the StringVal production.
	ExitStringVal(c *StringValContext)

	// ExitNumberVal is called when exiting the NumberVal production.
	ExitNumberVal(c *NumberValContext)

	// ExitTrueVal is called when exiting the TrueVal production.
	ExitTrueVal(c *TrueValContext)

	// ExitFalseVal is called when exiting the FalseVal production.
	ExitFalseVal(c *FalseValContext)

	// ExitNullVal is called when exiting the NullVal production.
	ExitNullVal(c *NullValContext)

	// ExitArrayVal is called when exiting the ArrayVal production.
	ExitArrayVal(c *ArrayValContext)

	// ExitStr_array is called when exiting the str_array production.
	ExitStr_array(c *Str_arrayContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)
}
