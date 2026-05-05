package providers

import "time"

type Evidence struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	Weight int `json:"weight"`
}

type Result struct {
	Site       string     `json:"site"`
	Category   string     `json:"category"`
	Username   string     `json:"username"`
	URL        string     `json:"url"`
	Found      bool       `json:"found"`
	StatusCode int        `json:"status_code"`
	Score      int        `json:"score"`
	Confidence string     `json:"confidence"`
	Evidences  []Evidence `json:"evidences"`
	Error      string     `json:"error,omitempty"`
	CheckedAt  time.Time  `json:"checked_at"`
}

type Provider interface {
	Name() string
	Category() string
	URL(username string) string
	Check(username string) Result
}
