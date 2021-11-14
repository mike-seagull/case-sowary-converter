package main

import "testing"

func TestSnakeCase(t *testing.T) {
	correct_strings := [][]string{
		{"a", "String"},
		{"A", "String"},
		{"a", "string"},
	}
	for _, test_string := range correct_strings {
		s := SnakeCase(test_string)
		if s != "a_string" {
			t.Error()
		}
	}
	if KebabCase([]string{""}) != "" {
		t.Error()
	}
	if KebabCase([]string{" "}) != " " {
		t.Error()
	}
	if SnakeCase([]string{"astring"}) != "astring" {
		t.Error()
	}
	if SnakeCase([]string{"ASTRING"}) != "astring" {
		t.Error()
	}
	if SnakeCase([]string{"a-String"}) != "a-string" {
		t.Error()
	}
}
func TestKebabCase(t *testing.T) {
	correct_strings := [][]string{
		{"a", "String"},
		{"A", "String"},
		{"a", "string"},
	}
	for _, test_string := range correct_strings {
		s := KebabCase(test_string)
		if s != "a-string" {
			t.Error()
		}
	}
	if KebabCase([]string{""}) != "" {
		t.Error()
	}
	if KebabCase([]string{" "}) != " " {
		t.Error()
	}
	if KebabCase([]string{"astring"}) != "astring" {
		t.Error()
	}
	if KebabCase([]string{"ASTRING"}) != "astring" {
		t.Error()
	}
	if KebabCase([]string{"a-String"}) != "a-string" {
		t.Error()
	}
}
