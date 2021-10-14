package logger

import (
	"fmt"
	"regexp"
	"strings"
)

/* Utils for Generic uber-go/zap FieldBuilder Template - ysyesilyurt 2021 */

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Color represents a text color.
type Color uint8

// Add adds the coloring to the given string.
func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}

/*
	checkIsStringOneOfOrContained checks whether given string is one of or is included in the given possible strings.
	e.g:
		checkIsStringOneOfOrContained("test-local", "local", "prod", "test") => true
		checkIsStringOneOfOrContained("dev", "local", "prod", "test") => false
		checkIsStringOneOfOrContained("development", "dev", "prod", "test") => true
		checkIsStringOneOfOrContained("dev", "dev", "prod", "test") => true
*/
func checkIsStringOneOfOrContained(target string, checkAgainst ...string) bool {
	if len(checkAgainst) < 1 {
		return false
	}

	var sb strings.Builder
	for _, env := range checkAgainst {
		sb.WriteString(fmt.Sprintf("%s|", env))
	}
	reStr := sb.String()
	return regexp.MustCompile(reStr[:len(reStr)-1]).MatchString(target)
}