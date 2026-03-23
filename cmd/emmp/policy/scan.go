package policy

import "fmt"

// ScanResult holds namespace scan results.
type ScanResult struct {
	Findings []ScanFinding
	Summary  ScanSummary
}

// ScanFinding is a posture finding.
type ScanFinding struct {
	Code     string
	Severity string
	Message  string
	Line     int
}

// ScanSummary aggregates findings.
type ScanSummary struct {
	Info     int
	Warn     int
	High     int
	Critical int
	Total    int
}

// FormatScanResult formats scan results for CLI output.
func FormatScanResult(result *ScanResult) string {
	if result.Summary.Total == 0 {
		return "PASS: no posture findings"
	}

	out := fmt.Sprintf("FINDINGS: %d total (%d critical, %d high, %d warn, %d info)\n",
		result.Summary.Total, result.Summary.Critical, result.Summary.High,
		result.Summary.Warn, result.Summary.Info)

	for _, f := range result.Findings {
		prefix := "  "
		if f.Line > 0 {
			prefix = fmt.Sprintf("  [line %d] ", f.Line)
		}
		out += fmt.Sprintf("%s%s: %s (%s)\n", prefix, f.Severity, f.Message, f.Code)
	}

	return out
}

// LintResult holds lint results.
type LintResult struct {
	Valid    bool
	Findings []LintFinding
}

// LintFinding is a lint finding.
type LintFinding struct {
	Severity string
	Message  string
	File     string
	Line     int
}

// FormatLintResult formats lint results for CLI output.
func FormatLintResult(result *LintResult) string {
	if result.Valid && len(result.Findings) == 0 {
		return "PASS: no lint findings"
	}

	out := fmt.Sprintf("LINT: %d findings\n", len(result.Findings))
	for _, f := range result.Findings {
		loc := ""
		if f.File != "" {
			loc = f.File
			if f.Line > 0 {
				loc = fmt.Sprintf("%s:%d", f.File, f.Line)
			}
			loc += " "
		}
		out += fmt.Sprintf("  %s%s: %s\n", loc, f.Severity, f.Message)
	}
	return out
}
