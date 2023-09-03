package rules

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type rulesTable struct {
	table table.Table
	rules *VMGroupResponse
}

func NewRulesTable(rules *VMGroupResponse, columns ...interface{}) rulesTable {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New(columns...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	return rulesTable{table: tbl, rules: rules}
}

func (t *rulesTable) ListRulesByGroupName(name string) table.Table {
	if name == "" {
		return nil
	}

	for _, vmGroup := range t.rules.Data.Groups {
		if vmGroup.Name == name {
			for _, rule := range vmGroup.Rules {
				t.table.AddRow(rule.Name, rule.State, rule.Health)
			}
		}
	}

	return t.table
}

func (t *rulesTable) ListRulesByLabels(label, value string) table.Table {
	if label == "" && value == "" {
		return nil
	}

	for _, vmGroup := range t.rules.Data.Groups {
		for _, rule := range vmGroup.Rules {
			val, ok := rule.Labels[label]
			if ok && val == value {
				t.table.AddRow(rule.Name, rule.State, rule.Health)
			}
		}
	}

	return t.table
}
