package policy

import "fmt"

// TestResult holds the result of a policy test run.
type TestResult struct {
	Total              int
	Passed             int
	Failed             int
	CoveragePct        float64
	CoverageGatePassed bool
	Failures           []TestFailure
}

// TestFailure describes a failed test case.
type TestFailure struct {
	TestCaseID       string
	Name             string
	ExpectedDecision string
	ActualDecision   string
}

// FormatTestResult formats the test result for CLI output.
func FormatTestResult(result *TestResult) string {
	status := "PASS"
	if result.Failed > 0 {
		status = "FAIL"
	}

	out := fmt.Sprintf("%s: %d/%d tests passed (%.1f%% coverage)\n",
		status, result.Passed, result.Total, result.CoveragePct)

	if !result.CoverageGatePassed {
		out += fmt.Sprintf("  WARNING: coverage %.1f%% below 80%% gate\n", result.CoveragePct)
	}

	for _, f := range result.Failures {
		out += fmt.Sprintf("  FAIL: %s — expected %s, got %s\n",
			f.Name, f.ExpectedDecision, f.ActualDecision)
	}

	return out
}
