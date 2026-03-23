package policy

import (
	"strings"
	"testing"
)

func TestFormatScanResult_Clean(t *testing.T) {
	result := &ScanResult{Summary: ScanSummary{}}
	out := FormatScanResult(result)
	if !strings.Contains(out, "PASS") {
		t.Errorf("clean scan should show PASS: %q", out)
	}
}

func TestFormatScanResult_WithFindings(t *testing.T) {
	result := &ScanResult{
		Findings: []ScanFinding{
			{Code: "BROAD_PERMIT", Severity: "HIGH", Message: "overly broad", Line: 5},
			{Code: "SHADOWED_RULE", Severity: "WARN", Message: "shadowed"},
		},
		Summary: ScanSummary{High: 1, Warn: 1, Total: 2},
	}
	out := FormatScanResult(result)
	if !strings.Contains(out, "2 total") {
		t.Error("should show total")
	}
	if !strings.Contains(out, "line 5") {
		t.Error("should show line number")
	}
	if !strings.Contains(out, "BROAD_PERMIT") {
		t.Error("should show finding code")
	}
}

func TestFormatLintResult_Clean(t *testing.T) {
	result := &LintResult{Valid: true}
	out := FormatLintResult(result)
	if !strings.Contains(out, "PASS") {
		t.Errorf("clean lint should show PASS: %q", out)
	}
}

func TestFormatLintResult_WithFindings(t *testing.T) {
	result := &LintResult{
		Valid: false,
		Findings: []LintFinding{
			{Severity: "ERROR", Message: "missing header", File: "pool.cedar", Line: 1},
			{Severity: "WARN", Message: "unused entity"},
		},
	}
	out := FormatLintResult(result)
	if !strings.Contains(out, "2 findings") {
		t.Error("should show count")
	}
	if !strings.Contains(out, "pool.cedar:1") {
		t.Error("should show file:line")
	}
}
