package alerts

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type alertsTable struct {
	table  table.Table
	alerts *VMAlertsResponse
}

func NewAlertsTable(alerts *VMAlertsResponse, columns ...interface{}) alertsTable {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New(columns...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	return alertsTable{table: tbl, alerts: alerts}
}

func (t *alertsTable) ListAlertsByName(name string) table.Table {
	if name == "" {
		return nil
	}

	for _, vmAlert := range t.alerts.Data.Alerts {
		if vmAlert.Name == name {
			t.table.AddRow(vmAlert.Name, vmAlert.FormatLabels(), vmAlert.FormatAnnotations(), vmAlert.State, vmAlert.GetActivateAt())
		}
	}

	return t.table
}

func (t *alertsTable) ListAlertsByLabels(label, value string) table.Table {
	if label == "" && value == "" {
		return nil
	}

	for _, vmAlert := range t.alerts.Data.Alerts {
		val, ok := vmAlert.Labels[label]
		if ok && val == value {
			t.table.AddRow(vmAlert.Name, vmAlert.FormatLabels(), vmAlert.FormatAnnotations(), vmAlert.State, vmAlert.GetActivateAt())
		}
	}

	return t.table
}
