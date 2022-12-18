package httphandling

import "testing"

func TestValidateURL(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		input string
		want  bool
	}{
		{"https://golang.org", true},
		{"http://golang.org", true},
		{"http:golang.org", false},
		{"http//golang.org", false},
		{"http:/golang.org", false},
		{"golang.org", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ValidateURL(tt.input); got != tt.want {
				t.Errorf("ValidateURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
