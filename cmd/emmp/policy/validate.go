package policy

import "fmt"

// ValidateResult holds the result of a policy validation.
type ValidateResult struct {
	Valid    bool
	Errors   []ValidationError
	Warnings []ValidationError
}

// ValidationError is a validation error with location.
type ValidationError struct {
	Line    int
	Column  int
	Message string
	Code    string
}

// FormatValidateResult formats the validation result for CLI output.
func FormatValidateResult(result *ValidateResult) string {
	if result.Valid {
		return "PASS: policy is valid"
	}

	out := "FAIL: policy validation failed\n"
	for _, e := range result.Errors {
		if e.Line > 0 {
			out += fmt.Sprintf("  ERROR [line %d]: %s (%s)\n", e.Line, e.Message, e.Code)
		} else {
			out += fmt.Sprintf("  ERROR: %s (%s)\n", e.Message, e.Code)
		}
	}
	for _, w := range result.Warnings {
		out += fmt.Sprintf("  WARN: %s (%s)\n", w.Message, w.Code)
	}
	return out
}
