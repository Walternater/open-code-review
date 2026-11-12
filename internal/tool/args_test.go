// SPDX-License-Identifier: Apache-2.0
// Copyright 2026 alibaba/open-code-review Contributors

package tool

import (
	"strings"
	"testing"
)

func TestBoolToolArg(t *testing.T) {
	for _, tc := range []struct {
		name        string
		args        map[string]any
		wantValue   bool
		wantPresent bool
		wantErr     string
	}{
		{name: "absent", args: map[string]any{}},
		{name: "nil", args: map[string]any{"flag": nil}},
		{name: "bool true", args: map[string]any{"flag": true}, wantValue: true, wantPresent: true},
		{name: "bool false", args: map[string]any{"flag": false}, wantPresent: true},
		{name: "string true", args: map[string]any{"flag": "true"}, wantValue: true, wantPresent: true},
		{name: "string true uppercase", args: map[string]any{"flag": "TRUE"}, wantValue: true, wantPresent: true},
		{name: "string false padded", args: map[string]any{"flag": " false "}, wantPresent: true},
		{name: "empty string is absent", args: map[string]any{"flag": ""}},
		{name: "non-boolean string", args: map[string]any{"flag": "yes"}, wantErr: `flag must be a boolean, got "yes"`},
		{name: "number", args: map[string]any{"flag": float64(1)}, wantErr: "flag must be a boolean, got float64"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, present, err := boolToolArg(tc.args, "flag")
			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("boolToolArg() error = %v, want containing %q", err, tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("boolToolArg() error = %v", err)
			}
			if got != tc.wantValue || present != tc.wantPresent {
				t.Fatalf("boolToolArg() = (%v, %v), want (%v, %v)", got, present, tc.wantValue, tc.wantPresent)
			}
		})
	}
}
