// {NOLINT}

package main

import (
	"regexp"
)

var LineLinters = []LineLinter{
RegexLinter{
	LinterDesc{
		"style",
		"trailing-whitespace",
		"Trailing whitespace found"},
	`[ \t]$`},
RegexLinter{
	LinterDesc{
		"style",
		"semicolon",
		"Unnecessary semicolon use is discouraged"},
	`;$`},
RegexLinter{
	LinterDesc{
		"style",
		"comma-whitespace",
		"Whitespace should follow a comma"},
	`,[^ \t]`},
RegexLinter{
	LinterDesc{
		"style",
		"tabs-only",
		"Only tabs should be used for indentation"},
	`^\t* +`},
RegexLinter{
	LinterDesc{
		"style",
		"embedded-tabs",
		"Tabs should only be used for indentation"},
	`[^ \t]\t`},
SimpleLineLinter{
	LinterDesc{
		"style",
		"deep-indent",
		"More than 5 levels of indentation"},
	func (line string) (bool, string) {
	indent := 0
	for _, c := range line {
		if c == '\t' {
			indent += 1
		} else {
			break
		}
	}
	return indent > 5, ""
}},
SimpleLineLinter{
	LinterDesc{
		"style",
		"line-length",
		"Line length should not exceed 80 characters"},
	func (line string) (bool, string) {
	ll := 0
	for _, c := range line {
		if c == '\t' {
			ll += 8
		} else {
			ll += 1
		}
	}
	return ll > 80, ""
}},
SimpleLineLinter{
	LinterDesc{
		"misc",
		"todo",
		"TODO notice"},
	func (line string) (bool, string) {
	r := regexp.MustCompile("(//+|/\\*) *((TODO|FIXME|XXX)( (.*))?)$")
	ms := r.FindStringSubmatch(line)
	if ms != nil {
		return true, ms[2]
	}
	return false, ""
}},
}

var ParsingLinters = [...]ParsingLinter{
OverlappingImportsLinter{},
}
