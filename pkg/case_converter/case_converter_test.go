package case_converter

import (
	"reflect"
	"testing"
)

func Test_basecase(t *testing.T) {
	// base cases cover most of kebab and snake
	type args struct {
		words     []string
		delimiter string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"no case", args{[]string{"MY", "STRING"}, ""}, "mystring"},
		{"one word no delimeter", args{[]string{"STRING"}, ""}, "string"},
		{"one word with delimeter", args{[]string{"STRING"}, "_"}, "string"},
		{"empty list no delimiter", args{[]string{}, ""}, ""},
		{"empty list with delimiter", args{[]string{}, "_"}, ""},
		{"empty string no delimiter", args{[]string{""}, ""}, ""},
		{"empty string with delimiter", args{[]string{""}, "_"}, ""},
		{"space no delimiter", args{[]string{" "}, ""}, " "},
		{"space with delimiter", args{[]string{" "}, "-"}, " "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basecase(tt.args.words, tt.args.delimiter); got != tt.want {
				t.Errorf("basecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSnakeCase(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"correct string 1", args{[]string{"a", "String"}}, "a_string"},
		{"correct string 2", args{[]string{"A", "String"}}, "a_string"},
		{"correct string 3", args{[]string{"a", "string"}}, "a_string"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCase(tt.args.words); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestKebabCase(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"correct string 1", args{[]string{"a", "String"}}, "a-string"},
		{"correct string 2", args{[]string{"A", "String"}}, "a-string"},
		{"correct string 3", args{[]string{"a", "string"}}, "a-string"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KebabCase(tt.args.words); got != tt.want {
				t.Errorf("KebabCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestPascalCase(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"correct string 1", args{[]string{"a", "String"}}, "AString"},
		{"correct string 2", args{[]string{"A", "String"}}, "AString"},
		{"correct string 3", args{[]string{"a", "string"}}, "AString"},
		{"no case", args{[]string{"MY", "STRING"}}, "MyString"},
		{"one word", args{[]string{"STRING"}}, "String"},
		{"empty list", args{[]string{}}, ""},
		{"empty string", args{[]string{""}}, ""},
		{"space", args{[]string{" "}}, " "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PascalCase(tt.args.words); got != tt.want {
				t.Errorf("PascalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCamalCase(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"correct string 1", args{[]string{"a", "String"}}, "aString"},
		{"correct string 2", args{[]string{"A", "String"}}, "aString"},
		{"correct string 3", args{[]string{"a", "string"}}, "aString"},
		{"no case", args{[]string{"MY", "STRING"}}, "myString"},
		{"one word", args{[]string{"STRING"}}, "string"},
		{"empty list", args{[]string{}}, ""},
		{"empty string", args{[]string{""}}, ""},
		{"space", args{[]string{" "}}, " "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamalCase(tt.args.words); got != tt.want {
				t.Errorf("CamalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestIdentifyCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"snake", args{"my_string"}, []string{"my", "string"}, false},
		{"kebab", args{"my-string"}, []string{"my", "string"}, false},
		{"pascal", args{"MyString"}, []string{"My", "String"}, false},
		{"camal", args{"myString"}, []string{"my", "String"}, false},
		{"lowercase", args{"mystring"}, []string{"mystring"}, false},
		{"uppercase", args{"MYSTRING"}, nil, true},
		{"multiple words (space case)", args{"my space"}, nil, true},
		{"multiple words (space case)", args{"my space"}, nil, true},
		{"empty string", args{""}, nil, true},
		{"empty list", args{}, nil, true},
		{"space", args{" "}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IdentifyCase(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdentifyCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdentifyCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
