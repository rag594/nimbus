package models

import "time"

type VMAlerts struct {
	Status string `json:"status"`
	Data   struct {
		Alerts []struct {
			State  string `json:"state"`
			Name   string `json:"name"`
			Value  string `json:"value"`
			Labels struct {
				Alertgroup string `json:"alertgroup"`
				Alertname  string `json:"alertname"`
				Alerttype  string `json:"alerttype"`
				Pg         string `json:"pg"`
				Product    string `json:"product"`
				Team       string `json:"team"`
			} `json:"labels"`
			Annotations struct {
				Description string `json:"description"`
				Mode        string `json:"mode"`
				Summary     string `json:"summary"`
			} `json:"annotations"`
			ActiveAt   time.Time `json:"activeAt"`
			Id         string    `json:"id"`
			RuleId     string    `json:"rule_id"`
			GroupId    string    `json:"group_id"`
			Expression string    `json:"expression"`
			Source     string    `json:"source"`
			Restored   bool      `json:"restored"`
		} `json:"alerts"`
	} `json:"data"`
}
