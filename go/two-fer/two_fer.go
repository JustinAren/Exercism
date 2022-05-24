// Package twofer provides a simple hello world alternative.
package twofer

import "fmt"

// ShareWith returns One for [given name] (or [you], if empty), one for me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
