package policy

import "fmt"

// AnalyzeResult holds access analysis results.
type AnalyzeResult struct {
	Entries []AccessEntry
	Total   int
}

// AccessEntry is a principal/action/resource/decision tuple.
type AccessEntry struct {
	Principal string
	Action    string
	Resource  string
	Decision  string
}

// FormatAnalyzeResult formats analysis results as a table.
func FormatAnalyzeResult(result *AnalyzeResult, queryType string) string {
	if result.Total == 0 {
		return "No results found."
	}

	header := ""
	switch queryType {
	case "what-can-principal-call":
		header = fmt.Sprintf("%-40s %-20s %-10s\n", "RESOURCE", "ACTION", "DECISION")
	case "who-can-call-resource":
		header = fmt.Sprintf("%-40s %-20s %-10s\n", "PRINCIPAL", "ACTION", "DECISION")
	default:
		header = fmt.Sprintf("%-40s %-20s %-40s %-10s\n", "PRINCIPAL", "ACTION", "RESOURCE", "DECISION")
	}

	out := header
	out += fmt.Sprintf("%s\n", dashLine(len(header)-1))

	for _, e := range result.Entries {
		switch queryType {
		case "what-can-principal-call":
			out += fmt.Sprintf("%-40s %-20s %-10s\n", e.Resource, e.Action, e.Decision)
		case "who-can-call-resource":
			out += fmt.Sprintf("%-40s %-20s %-10s\n", e.Principal, e.Action, e.Decision)
		default:
			out += fmt.Sprintf("%-40s %-20s %-40s %-10s\n", e.Principal, e.Action, e.Resource, e.Decision)
		}
	}

	out += fmt.Sprintf("\nTotal: %d entries\n", result.Total)
	return out
}

func dashLine(n int) string {
	if n <= 0 {
		n = 70
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = '-'
	}
	return string(b)
}
