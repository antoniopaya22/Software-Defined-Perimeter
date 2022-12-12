// Code generated from Aspida.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Aspida
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseAspidaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAspidaVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitMain(ctx *MainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitHosts(ctx *HostsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitTasks(ctx *TasksContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitVariables(ctx *VariablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitMain_content(ctx *Main_contentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitNameMain(ctx *NameMainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitConnectionMain(ctx *ConnectionMainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitDescriptionMain(ctx *DescriptionMainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitConnection(ctx *ConnectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitConnectionLocal(ctx *ConnectionLocalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitConnectionSSH(ctx *ConnectionSSHContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitDescription(ctx *DescriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitTContent(ctx *TContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitTSections(ctx *TSectionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitTPoints(ctx *TPointsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitTControls(ctx *TControlsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitTExclusions(ctx *TExclusionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitTTags(ctx *TTagsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitSections(ctx *SectionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitPoints(ctx *PointsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitControls(ctx *ControlsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitExclusions(ctx *ExclusionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitTags(ctx *TagsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitIfStat(ctx *IfStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitElifStat(ctx *ElifStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitElseStat(ctx *ElseStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitVars_content(ctx *Vars_contentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitVarDef(ctx *VarDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitVarObjDef(ctx *VarObjDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitComp_op(ctx *Comp_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitStringVal(ctx *StringValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitNumberVal(ctx *NumberValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitTrueVal(ctx *TrueValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitFalseVal(ctx *FalseValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitNullVal(ctx *NullValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitArrayVal(ctx *ArrayValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitStr_array(ctx *Str_arrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAspidaVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}
