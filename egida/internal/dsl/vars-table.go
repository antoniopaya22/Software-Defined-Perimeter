package dsl

type VariablesTable struct {
	Table map[string]interface{}
	ActualScope int
	ActualObj []map[string]interface{}
}

func NewVariablesTable() VariablesTable {
	var temp []map[string]interface{}
	return VariablesTable{Table: make(map[string]interface{}), ActualScope: 0, ActualObj: temp}
}

func (t *VariablesTable) NewVariable(id string, value interface{}) bool {
	if t.ActualScope == 0 {
		if _, ok := t.Table[id]; ok {
			return false
		}
		t.Table[id] = value
		return true
	} else {
		if _, ok := t.ActualObj[len(t.ActualObj)-1][id]; ok {
			return false
		}
		t.ActualObj[len(t.ActualObj)-1][id] = value
		return true
	}
}

func (t *VariablesTable) NewObject(id string) bool {
	if _, ok := t.Table[id]; ok {
		return false
	}
	t.ActualObj = append(t.ActualObj, make(map[string]interface{}))
	t.ActualScope = t.ActualScope + 1
	return true
}

func (t *VariablesTable) EndObject(id string) bool {
	t.ActualScope -= 1
	if t.ActualScope == 0 {
		t.Table[id] = t.ActualObj[0]
	} else {
		t.ActualObj[len(t.ActualObj)-2][id] = t.ActualObj[len(t.ActualObj)-1]
	}
	if len(t.ActualObj) > 0 {
		t.ActualObj = t.ActualObj[:len(t.ActualObj)-1]
	}
	return true
}