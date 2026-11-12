// SPDX-License-Identifier: Apache-2.0
// Copyright 2026 alibaba/open-code-review Contributors

package tool

import (
	"fmt"
	"strings"
)

// boolToolArg reads an optional boolean tool argument.
//
// JSON decodes booleans as bool, but a model — or a gateway forwarding a
// model's raw completion text as arguments — can send "true"/"false" as a
// string. A bare args[name].(bool) assertion turns that into the zero value,
// which silently runs the opposite behavior instead of answering the question
// that was asked. Coerce the string form and make anything else an explicit
// error the model can act on.
func boolToolArg(args map[string]any, name string) (bool, bool, error) {
	value, present := args[name]
	if !present || value == nil {
		return false, false, nil
	}
	switch v := value.(type) {
	case bool:
		return v, true, nil
	case string:
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			return false, false, nil
		}
		switch strings.ToLower(trimmed) {
		case "true":
			return true, true, nil
		case "false":
			return false, true, nil
		}
		return false, false, fmt.Errorf("%s must be a boolean, got %q", name, v)
	default:
		return false, false, fmt.Errorf("%s must be a boolean, got %T", name, value)
	}
}
