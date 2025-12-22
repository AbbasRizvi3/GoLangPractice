package main

import (
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid name", "Gladys", false},
		{"empty name", "", true},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			msg, err := Hello(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hello(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if tt.input != "" && !strings.Contains(msg, tt.input) {
				t.Errorf("Hello(%q) = %q, want message containing %q", tt.input, msg, tt.input)
			}
		})
	}
}
