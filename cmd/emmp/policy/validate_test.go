package policy

import (
	"strings"
	"testing"
)

func TestFormatValidateResult_Pass(t *testing.T) {
	result := &ValidateResult{Valid: true}
	out := FormatValidateResult(result)
	if !strings.Contains(out, "PASS") {
		t.Errorf("output = %q", out)
	}
}

func TestFormatValidateResult_Fail(t *testing.T) {
	result := &ValidateResult{
		Valid: false,
		Errors: []ValidationError{
			{Line: 3, Message: "syntax error", Code: "SYNTAX_ERROR"},
		},
		Warnings: []ValidationError{
			{Message: "no rules", Code: "EMPTY_POLICY"},
		},
	}
	out := FormatValidateResult(result)
	if !strings.Contains(out, "FAIL") {
		t.Error("should contain FAIL")
	}
	if !strings.Contains(out, "line 3") {
		t.Error("should contain line number")
	}
	if !strings.Contains(out, "SYNTAX_ERROR") {
		t.Error("should contain error code")
	}
	if !strings.Contains(out, "WARN") {
		t.Error("should contain warning")
	}
}

func TestFormatTestResult_Pass(t *testing.T) {
	result := &TestResult{
		Total: 10, Passed: 10, Failed: 0,
		CoveragePct: 85.0, CoverageGatePassed: true,
	}
	out := FormatTestResult(result)
	if !strings.Contains(out, "PASS") {
		t.Error("should contain PASS")
	}
	if !strings.Contains(out, "10/10") {
		t.Error("should contain pass count")
	}
}

func TestFormatTestResult_Fail(t *testing.T) {
	result := &TestResult{
		Total: 5, Passed: 3, Failed: 2,
		CoveragePct: 60.0, CoverageGatePassed: false,
		Failures: []TestFailure{
			{Name: "admin deny", ExpectedDecision: "PERMIT", ActualDecision: "DENY"},
		},
	}
	out := FormatTestResult(result)
	if !strings.Contains(out, "FAIL") {
		t.Error("should contain FAIL")
	}
	if !strings.Contains(out, "80%") {
		t.Error("should warn about coverage gate")
	}
	if !strings.Contains(out, "admin deny") {
		t.Error("should list failure")
	}
}
