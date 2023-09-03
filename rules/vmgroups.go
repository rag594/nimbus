package rules

import (
	"time"

	"github.com/rag594/nimbus/alerts"
)

type VMGroupResponse struct {
	Status string    `json:"status"`
	Data   GroupData `json:"data"`
}

type GroupData struct {
	Groups []Group `json:"groups"`
}

type Group struct {
	Name           string    `json:"name"`
	Rules          []Rule    `json:"rules"`
	Interval       int       `json:"interval"`
	LastEvaluation time.Time `json:"lastEvaluation"`
	Type           string    `json:"type"`
	Id             string    `json:"id"`
	File           string    `json:"file"`
	Concurrency    int       `json:"concurrency"`
}

type Rule struct {
	State             string            `json:"state"`
	Name              string            `json:"name"`
	Query             string            `json:"query"`
	Duration          int               `json:"duration"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	LastError         string            `json:"lastError"`
	EvaluationTime    float64           `json:"evaluationTime"`
	LastEvaluation    time.Time         `json:"lastEvaluation"`
	Health            string            `json:"health"`
	Type              string            `json:"type"`
	DatasourceType    string            `json:"datasourceType"`
	LastSamples       int               `json:"lastSamples"`
	LastSeriesFetched int               `json:"lastSeriesFetched"`
	Id                string            `json:"id"`
	GroupId           string            `json:"group_id"`
	Debug             bool              `json:"debug"`
	MaxUpdatesEntries int               `json:"max_updates_entries"`
	Alerts            []alerts.Alert    `json:"alerts,omitempty"`
}
