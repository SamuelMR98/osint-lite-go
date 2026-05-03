package utils

// Status code to string mapping
var statusText = map[int]string{
	200: "OK",
	301: "Moved Permanently",
	302: "Found",
	403: "Forbidden",
	404: "Not Found",
	500: "Internal Server Error",
}

// GetStatusText returns the status text for a given status code
func GetStatusText(code int) string {
	if text, exists := statusText[code]; exists {
		return text
	}
	return "Unknown Status"
}