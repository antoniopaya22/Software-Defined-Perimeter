package dsl

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/antonioalfa22/egida/internal/info"
	"github.com/antonioalfa22/go-utils/collections"
)

type Comparation struct {
	Host     string
	Value1   interface{}
	Operator string
	Value2   interface{}
}

func (c Comparation) Compare() bool {
	query := strings.Split(c.Value1.(string), ".")[0]
	value := strings.Split(c.Value1.(string), ".")[1]
	switch strings.ToLower(query) {
	case "services":
		if strings.ToLower(c.Value2.(string)) == "stopped" {
			alls, _ := info.GetStoppedServices([]string{c.Host})
			result := alls[0]
			return collections.Find(result.Lines, func(x string) bool {
				return strings.Split(x, "->")[0] == value
			}) != nil
		} else if strings.ToLower(c.Value2.(string)) == "running" {
			alls, _ := info.GetRunningServices([]string{c.Host})
			result := alls[0]
			return collections.Find(result.Lines, func(x string) bool {
				return strings.Split(x, "->")[0] == value
			}) != nil
		} else {
			alls, _ := info.GetAllServices([]string{c.Host})
			result := alls[0]
			return collections.Find(result.Lines, func(x string) bool {
				return strings.Split(x, "->")[0] == value
			}) != nil
		}
	case "packages":
		if strings.ToLower(c.Value2.(string)) == "installed" {
			alls, _ := info.GetAllPackages([]string{c.Host})
			result := alls[0]
			return collections.Find(result.Lines, func(x string) bool {
				return x == value
			}) != nil
		}
	case "hardscores":
		lines, _ := info.GetLynisScore([]string{c.Host})
		re := regexp.MustCompile("[0-9]+")
		score, _ := strconv.ParseFloat(re.FindAllString(lines[0].Lines[0], -1)[0], 64)
		value, _ := strconv.ParseFloat(c.Value2.(string), 64)
		switch c.Operator {
		case "==":
			return score == value
		case ">=":
			return score >= value
		case "<=":
			return score <= value
		case ">":
			return score > value
		case "<":
			return score < value
		case "!=":
			return score != value
		default:
			return false
		}
	case "machine":
		// Get Open ports
		if value == "open_port" {
			p, _ := strconv.ParseInt(c.Value2.(string),10,  64)
			ports, _ := info.GetMachineOpenPorts(c.Host)
			return collections.Find(ports, func(x int64) bool {
				return x == p
			}) != nil
		} else { // Machine Has GUI
			alls, _ := info.GetMachineHasGUI([]string{c.Host})
			value, _ := strconv.ParseBool(c.Value2.(string))
			result := alls[0]
			return result == value
		}
	}
	return false
}
