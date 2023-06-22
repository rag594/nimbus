package models

import "time"

type VMGroups struct {
	Status string `json:"status"`
	Data   struct {
		Groups []struct {
			Name  string `json:"name"`
			Rules []struct {
				State       string            `json:"state"`
				Name        string            `json:"name"`
				Query       string            `json:"query"`
				Duration    int               `json:"duration"`
				Labels      map[string]string `json:"labels"`
				Annotations struct {
					Summary string `json:"summary"`
				} `json:"annotations"`
				LastError         string    `json:"lastError"`
				EvaluationTime    float64   `json:"evaluationTime"`
				LastEvaluation    time.Time `json:"lastEvaluation"`
				Health            string    `json:"health"`
				Type              string    `json:"type"`
				DatasourceType    string    `json:"datasourceType"`
				LastSamples       int       `json:"lastSamples"`
				LastSeriesFetched int       `json:"lastSeriesFetched"`
				Id                string    `json:"id"`
				GroupId           string    `json:"group_id"`
				Debug             bool      `json:"debug"`
				MaxUpdatesEntries int       `json:"max_updates_entries"`
				Alerts            []struct {
					State  string `json:"state"`
					Name   string `json:"name"`
					Value  string `json:"value"`
					Labels struct {
						Alertgroup string `json:"alertgroup"`
						Alertname  string `json:"alertname"`
						Hpa        string `json:"hpa"`
						Namespace  string `json:"namespace"`
						Severity   string `json:"severity"`
						Squadcast  string `json:"squadcast"`
						Team       string `json:"team"`
					} `json:"labels"`
					Annotations struct {
						Summary string `json:"summary"`
					} `json:"annotations"`
					ActiveAt   time.Time `json:"activeAt"`
					Id         string    `json:"id"`
					RuleId     string    `json:"rule_id"`
					GroupId    string    `json:"group_id"`
					Expression string    `json:"expression"`
					Source     string    `json:"source"`
					Restored   bool      `json:"restored"`
				} `json:"alerts,omitempty"`
			} `json:"rules"`
			Interval       int       `json:"interval"`
			LastEvaluation time.Time `json:"lastEvaluation"`
			Type           string    `json:"type"`
			Id             string    `json:"id"`
			File           string    `json:"file"`
			Concurrency    int       `json:"concurrency"`
		} `json:"groups"`
	} `json:"data"`
}
