package policy

import (
	"strings"
	"testing"
)

func TestFormatAnalyzeResult_WhatCan(t *testing.T) {
	result := &AnalyzeResult{
		Entries: []AccessEntry{
			{Resource: "MCPServer::\"srv-1\"", Action: "InvokeTool", Decision: "PERMIT"},
		},
		Total: 1,
	}
	out := FormatAnalyzeResult(result, "what-can-principal-call")
	if !strings.Contains(out, "RESOURCE") {
		t.Error("should have RESOURCE header")
	}
	if !strings.Contains(out, "PERMIT") {
		t.Error("should show PERMIT decision")
	}
	if !strings.Contains(out, "Total: 1") {
		t.Error("should show total")
	}
}

func TestFormatAnalyzeResult_WhoCan(t *testing.T) {
	result := &AnalyzeResult{
		Entries: []AccessEntry{
			{Principal: "MCPApplication::\"app-1\"", Action: "InvokeTool", Decision: "PERMIT"},
			{Principal: "MCPApplication::\"app-2\"", Action: "InvokeTool", Decision: "PERMIT"},
		},
		Total: 2,
	}
	out := FormatAnalyzeResult(result, "who-can-call-resource")
	if !strings.Contains(out, "PRINCIPAL") {
		t.Error("should have PRINCIPAL header")
	}
	if !strings.Contains(out, "Total: 2") {
		t.Error("should show total")
	}
}

func TestFormatAnalyzeResult_Empty(t *testing.T) {
	result := &AnalyzeResult{Total: 0}
	out := FormatAnalyzeResult(result, "what-can-principal-call")
	if !strings.Contains(out, "No results") {
		t.Error("empty should say no results")
	}
}
