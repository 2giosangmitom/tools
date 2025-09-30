package main

import "testing"

func TestBuild(t *testing.T) {
	tests := []struct {
		name    string
		config  opts
		wantErr bool
	}{
		{
			name: "Valid config with all options",
			config: opts{
				length:           12,
				includeNumbers:   true,
				includeSymbols:   true,
				includeUppercase: true,
				includeLowercase: true,
				includeAlpha:     true,
			},
			wantErr: false,
		},
		{
			name: "Valid config with only lowercase",
			config: opts{
				length:           8,
				includeNumbers:   false,
				includeSymbols:   false,
				includeUppercase: false,
				includeLowercase: true,
				includeAlpha:     true,
			},
			wantErr: false,
		},
		{
			name: "Invalid config with no character sets",
			config: opts{
				length:           10,
				includeNumbers:   false,
				includeSymbols:   false,
				includeUppercase: false,
				includeLowercase: false,
				includeAlpha:     false,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := build(&tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("build() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
