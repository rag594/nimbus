package alerts

import "time"

type VMAlerts struct {
	Status string `json:"status"`
	Data   struct {
		Alerts []struct {
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
		} `json:"alerts"`
	} `json:"data"`
}
