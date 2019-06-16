// Package twofer helps with distributing resources between yourself and another person.
package twofer

import "fmt"

// ShareWith returns "One for you, one for me" unless given a name, in which case it substitues that name for "you".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
