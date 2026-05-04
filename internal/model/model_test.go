package model

import "testing"

func TestValidateGUID(t *testing.T) {
	tests := []struct {
		name     string
		guid     string
		expected bool
	}{
		{
			name:     "valid GUID lowercase",
			guid:     "05024756-765e-41a9-89d7-1407436d9a58",
			expected: true,
		},
		{
			name:     "valid GUID uppercase",
			guid:     "05024756-765E-41A9-89D7-1407436D9A58",
			expected: true,
		},
		{
			name:     "invalid GUID - too short",
			guid:     "short",
			expected: false,
		},
		{
			name:     "invalid GUID - wrong format",
			guid:     "this-is-a-bad-guid",
			expected: false,
		},
		{
			name:     "invalid GUID - missing dashes",
			guid:     "05024756765e41a989d71407436d9a58",
			expected: false,
		},
		{
			name:     "invalid GUID - wrong dash positions",
			guid:     "0502475-6765e-41a9-89d7-1407436d9a58",
			expected: false,
		},
		{
			name:     "invalid GUID - invalid characters",
			guid:     "05024756-765g-41a9-89d7-1407436d9a58",
			expected: false,
		},
		{
			name:     "empty string",
			guid:     "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateGUID(tt.guid)
			if result != tt.expected {
				t.Errorf("ValidateGUID(%q) = %v, want %v", tt.guid, result, tt.expected)
			}
		})
	}
}
