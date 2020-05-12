package collector

import "github.com/prometheus/client_golang/prometheus"

// Metrics tracks all the contextual metrics for this exporter
type Metrics struct {
	issue *prometheus.Desc
}

type jiraIssue struct {
	Issues []struct {
		Fields struct {
			Assignee struct {
				DisplayName string `json:"displayName"`
			} `json:"assignee"`
			Created string `json:"created"`
			Project struct {
				Name string `json:"name"`
			} `json:"project"`
			Status struct {
				Name string `json:"name"`
			} `json:"status"`
			Summary   string `json:"summary"`
			Issuetype struct {
				Name string `json:"name"`
			} `json:"issuetype"`
		} `json:"fields"`
		Key string `json:"key"`
	} `json:"issues"`
	MaxResults int `json:"maxResults"`
	StartAt    int `json:"startAt"`
	Total      int `json:"total"`
}
