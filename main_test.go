package main

import (
	"strings"
	"testing"
)

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
	if SnakeCase([]string{""}) != "" {
		t.Error()
	}
	if SnakeCase([]string{" "}) != " " {
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
func TestPascalCase(t *testing.T) {
	correct_strings := [][]string{
		{"a", "String"},
		{"A", "String"},
		{"a", "string"},
	}
	for _, test_string := range correct_strings {
		s := PascalCase(test_string)
		if s != "AString" {
			t.Error()
		}
	}
	if PascalCase([]string{""}) != "" {
		t.Error()
	}
	if PascalCase([]string{" "}) != " " {
		t.Error()
	}
	if PascalCase([]string{"astring"}) != "Astring" {
		t.Error()
	}
	if PascalCase([]string{"ASTRING"}) != "ASTRING" {
		t.Error()
	}
	if PascalCase([]string{"a-String"}) != "A-String" {
		t.Error()
	}
}
func TestCamalCase(t *testing.T) {
	correct_strings := [][]string{
		{"a", "String"},
		{"A", "String"},
		{"a", "string"},
	}
	for _, test_string := range correct_strings {
		s := CamalCase(test_string)
		if s != "aString" {
			t.Error()
		}
	}
	if CamalCase([]string{""}) != "" {
		t.Error()
	}
	if CamalCase([]string{" "}) != " " {
		t.Error()
	}
	if CamalCase([]string{"astring"}) != "astring" {
		t.Error()
	}
	if CamalCase([]string{"ASTRING"}) != "astring" {
		t.Error()
	}
	if CamalCase([]string{"a-String"}) != "a-string" {
		t.Error()
	}
}
func TestIdentifyCase(t *testing.T) {
	correct_strings := []string{
		"my_string",
		"my-string",
		"MyString",
		"myString",
	}
	for _, test_string := range correct_strings {
		result, err := IdentifyCase(test_string)
		if len(result) < 2 {
			t.Errorf("Not enough results. Length=%d. Got: %v at \"%s\"", len(result), result, test_string)
		} else if strings.ToLower(result[0]) != "my" && strings.ToLower(result[1]) != "string" {
			t.Errorf("Got: %v at \"%s\"", result, test_string)
		} else if err != nil {
			t.Error(err)
		}
	}
}
