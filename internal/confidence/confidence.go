package confidence

import "github.com/SamuelMR98/osint-lite-go/internal/providers"

// ConfidenceLevel represents the confidence level of a username check result.
func Score(evidence []providers.Evidence) int {
	score := 0

	for _, item := range evidence {
		score += item.Weight
	}

	if score < 0 {
		return 0
	}

	if score > 100 {
		return 100
	}

	return score
}

func Label(score int) string {
	switch {
	case score >= 80:
		return "High"
	case score >= 50:
		return "Medium"
	case score >= 20:
		return "Low"
	default:
		return "Unknown"
	}
}
