package alerts

import (
	"fmt"
	"strings"
	"time"
)

type VMAlertsResponse struct {
	Status string     `json:"status"`
	Data   AlertsData `json:"data"`
}

type AlertsData struct {
	Alerts []Alert `json:"alerts"`
}

type Alert struct {
	State       string            `json:"state"`
	Name        string            `json:"name"`
	Value       string            `json:"value"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	ActiveAt    time.Time         `json:"activeAt"`
	Id          string            `json:"id"`
	RuleId      string            `json:"rule_id"`
	GroupId     string            `json:"group_id"`
	Expression  string            `json:"expression"`
	Source      string            `json:"source"`
	Restored    bool              `json:"restored"`
}

func (a Alert) FormatLabels() string {
	var labelsBuilder strings.Builder
	for key, value := range a.Labels {
		labelsBuilder.WriteString(fmt.Sprintf("%s_%s\n", key, value))
	}
	return labelsBuilder.String()
}

func (a Alert) FormatAnnotations() string {
	var annotationsBuilder strings.Builder
	for key, value := range a.Annotations {
		annotationsBuilder.WriteString(fmt.Sprintf("%s_%s\n", key, value))
	}
	return annotationsBuilder.String()
}

func (a Alert) GetActivateAt() string {
	return a.ActiveAt.Local().Format("2006-01-02 15:04:05")
}
