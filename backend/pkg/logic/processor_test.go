package logic

import (
	"reflect"
	"testing"

	"github.com/Childebrand94/takeHomePhoneNumber/pkg/models"
)

// The file path must be changed to be relative form the test directory for
//the init function opening the data.json file or test will fail

func TestCreateMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []models.PrefixInfo
		expected map[string]models.PrefixInfo
	}{
		{
			name: "single item",
			input: []models.PrefixInfo{
				{Prefix: 123,
					Operator:    "",
					Region:      "",
					CountryCode: 1,
					Country:     "Canada"},
			},
			expected: map[string]models.PrefixInfo{
				"123": {
					Prefix:      123,
					Operator:    "",
					Region:      "",
					CountryCode: 1,
					Country:     "Canada"},
			},
		},
		{
			name:     "empty input",
			input:    []models.PrefixInfo{},
			expected: map[string]models.PrefixInfo{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result, err := createMap(tc.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Assert result length
			if len(result) != len(tc.expected) {
				t.Errorf("Expected map length %d, got %d", len(tc.expected), len(result))
			}

			// Assert individual items
			for key, expectedValue := range tc.expected {
				if actualValue, exists := result[key]; !exists || actualValue != expectedValue {
					t.Errorf("Expected key %s to have value %+v, got %+v", key, expectedValue, actualValue)
				}
			}
		})
	}
}

func TestGetPrefixInfo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected models.PrefixInfo
	}{
		{
			name:  "Empty String Input",
			input: "",
			expected: models.PrefixInfo{
				Prefix:      0,
				Operator:    "",
				Region:      "",
				CountryCode: 0,
				Country:     "",
			},
		},
		{
			name:  "Match longest prefix",
			input: "1437329",
			expected: models.PrefixInfo{
				Prefix:      1437329,
				Operator:    "Lucky Mobile",
				Region:      "Ontario",
				CountryCode: 1,
				Country:     "Canada",
			},
		},
		{
			name:  "Not a Match",
			input: "56893",
			expected: models.PrefixInfo{
				Prefix:      0,
				Operator:    "",
				Region:      "",
				CountryCode: 0,
				Country:     "",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := getPrefixInfo(tc.input)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("getPrefixInfo(%s) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestProcessMessage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No URLs",
			input:    "This is a test message without URLs.",
			expected: "This is a test message without URLs.",
		},
		{
			name:     "Single URL",
			input:    "Check out this link: http://example.com",
			expected: "Check out this link: <a href='http://example.com'>http://example.com</a>",
		},
		{
			name:     "Multiple URLs",
			input:    "Multiple links: http://example.com and https://example.org",
			expected: "Multiple links: <a href='http://example.com'>http://example.com</a> and <a href='https://example.org'>https://example.org</a>",
		},
		{
			name:     "Invalid URL",
			input:    "This is not a URL: http://",
			expected: "This is not a URL: http://",
		},
		{
			name:     "Mixed content",
			input:    "Text with a URL: http://example.com in the middle.",
			expected: "Text with a URL: <a href='http://example.com'>http://example.com</a> in the middle.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := processMessage(tt.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestProcessPhoneNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No phone number",
			input:    "",
			expected: "",
		},
		{
			name:     "No matching numbers",
			input:    "00000",
			expected: "",
		},
		{
			name:     "Matching number",
			input:    "143",
			expected: "143",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := processPhoneNumber(tt.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
