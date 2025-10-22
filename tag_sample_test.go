package main

import (
	"fmt"
	"strings"
	"testing"
)

// TestTagSystemWithSampleData tests the tag system with realistic sample data
func TestTagSystemWithSampleData(t *testing.T) {
	testCases := []struct {
		name         string
		description  string
		expectTags   int
		expectTypes  []string
		expectStatus []string
	}{
		{
			name:         "Cobra library with active status",
			description:  "[lib] [active] A Commander for modern Go CLI interactions.",
			expectTags:   2,
			expectTypes:  []string{"lib"},
			expectStatus: []string{"active"},
		},
		{
			name:         "Hugo application with active status",
			description:  "[app] [active] Fast and Modern Static Website Engine.",
			expectTags:   2,
			expectTypes:  []string{"app"},
			expectStatus: []string{"active"},
		},
		{
			name:         "Gox application with stalled status",
			description:  "[app] [stalled] Dead simple, no frills Go cross compile tool.",
			expectTags:   2,
			expectTypes:  []string{"app"},
			expectStatus: []string{"stalled"},
		},
		{
			name:         "Library with unmaintained status",
			description:  "[lib] [unmaintained] Legacy library no longer maintained.",
			expectTags:   2,
			expectTypes:  []string{"lib"},
			expectStatus: []string{"unmaintained"},
		},
		{
			name:         "Project with no tags",
			description:  "Standard project description without any tags.",
			expectTags:   0,
			expectTypes:  []string{},
			expectStatus: []string{},
		},
		{
			name:         "Project with only type tag",
			description:  "[lib] Library without status indication.",
			expectTags:   1,
			expectTypes:  []string{"lib"},
			expectStatus: []string{},
		},
		{
			name:         "Project with only status tag",
			description:  "[active] Project without type indication.",
			expectTags:   1,
			expectTypes:  []string{},
			expectStatus: []string{"active"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tags, cleanDesc := parseTagsFromDescription(tc.description)

			// Check number of tags
			if len(tags) != tc.expectTags {
				t.Errorf("Expected %d tags, got %d", tc.expectTags, len(tags))
			}

			// Collect actual tag types and statuses
			var foundTypes, foundStatus []string
			for _, tag := range tags {
				if tag.Type != "" {
					foundTypes = append(foundTypes, tag.Type)
				}
				if tag.Status != "" {
					foundStatus = append(foundStatus, tag.Status)
				}
			}

			// Check expected tag types
			if len(foundTypes) != len(tc.expectTypes) {
				t.Errorf("Expected %d type tags %v, got %d type tags %v", 
					len(tc.expectTypes), tc.expectTypes, len(foundTypes), foundTypes)
			} else {
				for i, expectedType := range tc.expectTypes {
					if i >= len(foundTypes) || foundTypes[i] != expectedType {
						t.Errorf("Expected type tag %q at index %d, got %q", 
							expectedType, i, safeIndex(foundTypes, i))
					}
				}
			}

			// Check expected status tags
			if len(foundStatus) != len(tc.expectStatus) {
				t.Errorf("Expected %d status tags %v, got %d status tags %v", 
					len(tc.expectStatus), tc.expectStatus, len(foundStatus), foundStatus)
			} else {
				for i, expectedStatus := range tc.expectStatus {
					if i >= len(foundStatus) || foundStatus[i] != expectedStatus {
						t.Errorf("Expected status tag %q at index %d, got %q", 
							expectedStatus, i, safeIndex(foundStatus, i))
					}
				}
			}

			// Verify clean description doesn't contain tag patterns using standard library
			for _, expectedType := range tc.expectTypes {
				tagPattern := fmt.Sprintf("[%s]", expectedType)
				if strings.Contains(cleanDesc, tagPattern) {
					t.Errorf("Clean description should not contain %s, got: %s", tagPattern, cleanDesc)
				}
			}
			
			for _, expectedStatus := range tc.expectStatus {
				tagPattern := fmt.Sprintf("[%s]", expectedStatus)
				if strings.Contains(cleanDesc, tagPattern) {
					t.Errorf("Clean description should not contain %s, got: %s", tagPattern, cleanDesc)
				}
			}

			fmt.Printf("✓ %s: %d tags found, clean desc: '%s'\n", tc.name, len(tags), cleanDesc)
		})
	}
}

// safeIndex returns the value at index i or "undefined" if index is out of bounds
func safeIndex(slice []string, i int) string {
	if i >= 0 && i < len(slice) {
		return slice[i]
	}
	return "undefined"
}

// TestTagSystemWhitespaceHandling tests proper whitespace normalization
func TestTagSystemWhitespaceHandling(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Multiple spaces after tag removal",
			input:    "Project  [lib]   [active]  description.",
			expected: "Project description.",
		},
		{
			name:     "Leading hyphen removal",
			input:    " - [lib] [active] Description after hyphen.",
			expected: "Description after hyphen.",
		},
		{
			name:     "Spaces before punctuation",
			input:    "Project [lib] description  ,  with  spaces  .",
			expected: "Project description, with spaces.",
		},
		{
			name:     "Complex whitespace normalization",
			input:    "  Multiple   [app]  spaces   [stalled]   everywhere  .  ",
			expected: "Multiple spaces everywhere.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, cleanDesc := parseTagsFromDescription(tc.input)
			if cleanDesc != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, cleanDesc)
			}

			// Verify no tag patterns remain using standard library strings.Contains
			tagPatterns := []string{"[lib]", "[app]", "[active]", "[stalled]", "[unmaintained]"}
			for _, pattern := range tagPatterns {
				if strings.Contains(cleanDesc, pattern) {
					t.Errorf("Clean description should not contain tag pattern %s, got: %s", 
						pattern, cleanDesc)
				}
			}
		})
	}
}