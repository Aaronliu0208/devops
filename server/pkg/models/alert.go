package models

import "time"

// AlertState type for alert status
type AlertState string

const (
	// AlertStateActive 报警激活
	AlertStateActive AlertState = "active"
	// AlertStateResovled 报警解决
	AlertStateResovled AlertState = "resovled"
)

//Raw Info for alert
type Raw map[string]interface{}

// Alert 报警信息
type Alert struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	StartsAt    time.Time  `json:"start_at"`
	EndsAt      time.Time  `json:"end_at,omitempty"`
	Source      string     `json:"source"`
	Description string     `json:"desc,omitempty"`
	Tags        Raw        `json:"raw,inline,omitempty"`
	State       AlertState `json:"state,omitempty"`
	RawID       string     `json:"raw_id,omitempty"`
}
