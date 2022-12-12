package dsl

import (
	"fmt"
	"github.com/antonioalfa22/go-utils/collections"
)

type TasksTable struct {
	Sections []string
	Controls []string
	Points []string
	Exclusions []string
	Tags []string
}

func NewTasksTable() TasksTable {
	return TasksTable{
		Sections: []string{},
		Controls: []string{},
		Points: []string{},
		Exclusions: []string{},
		Tags: []string{},
	}
}


func (t *TasksTable) AddSection(section string) {
	if collections.Find(t.Sections, func(x string) bool {return x == section}) == nil {
		t.Sections = append(t.Sections, section)
	} else {
		fmt.Println("Section", section, "ya existe")
	}
}

func (t *TasksTable) AddControl(control string) {
	if collections.Find(t.Controls, func(x string) bool {return x == control}) == nil {
		t.Controls = append(t.Controls, control)
	} else {
		fmt.Println("Control", control, "ya existe")
	}
}

func (t *TasksTable) AddPoint(point string) {
	if collections.Find(t.Points, func(x string) bool {return x == point}) == nil {
		t.Points = append(t.Points, point)
	} else {
		fmt.Println("Point", point, "ya existe")
	}
}

func (t *TasksTable) AddExclusionPoint(point string) {
	if collections.Find(t.Exclusions, func(x string) bool {return x == point}) == nil {
		t.Exclusions = append(t.Exclusions, point)
	} else {
		fmt.Println("Exclusion Point", point, "ya existe")
	}
}

func (t *TasksTable) AddTag(tag string) {
	if collections.Find(t.Tags, func(x string) bool {return x == tag}) == nil {
		t.Points = append(t.Points, tag)
	} else {
		fmt.Println("Tag", tag, "ya existe")
	}
}

func (t TasksTable) GetSections() []string {
	return t.Sections
}

func (t TasksTable) GetControls() []string {
	return t.Controls
}

func (t TasksTable) GetPoints() []string {
	return collections.Filter(t.Points, func(point string) bool {
		return collections.Find(t.Exclusions, func(x string) bool {return x == point}) == nil
	}).([]string)
}

func (t TasksTable) GetExclusions() []string {
	return t.Exclusions
}

func (t TasksTable) GetTags() []string {
	return t.Tags
}