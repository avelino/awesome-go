package main

import (
	"fmt"
	"testing"
)

// TestTagSystemWithSampleData tests the tag system with realistic sample data
func TestTagSystemWithSampleData(t *testing.T) {
	testCases := []struct {
		name        string
		description string
		expectTags  int
		expectTypes []string
	}{
		{
			name:        "Cobra library with active status",
			description: "[lib] [active] A Commander for modern Go CLI interactions.",
			expectTags:  2,
			expectTypes: []string{"lib", "active"},
		},
		{
			name:        "Hugo application with active status", 
			description: "[app] [active] Fast and Modern Static Website Engine.",
			expectTags:  2,
			expectTypes: []string{"app", "active"},
		},
		{
			name:        "Gox application with stalled status",
			description: "[app] [stalled] Dead simple, no frills Go cross compile tool.",
			expectTags:  2,
			expectTypes: []string{"app", "stalled"},
		},
		{
			name:        "Library with unmaintained status",
			description: "[lib] [unmaintained] Legacy library no longer maintained.",
			expectTags:  2,
			expectTypes: []string{"lib", "unmaintained"},
		},
		{
			name:        "Project with no tags",
			description: "Standard project description without any tags.",
			expectTags:  0,
			expectTypes: []string{},
		},
		{
			name:        "Project with only type tag",
			description: "[lib] Library without status indication.",
			expectTags:  1,
			expectTypes: []string{"lib"},
		},
		{
			name:        "Project with only status tag",
			description: "[active] Project without type indication.",
			expectTags:  1,
			expectTypes: []string{"active"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tags, cleanDesc := parseTagsFromDescription(tc.description)
			
			// Check number of tags
			if len(tags) != tc.expectTags {
				t.Errorf("Expected %d tags, got %d", tc.expectTags, len(tags))
			}

			// Collect all tag values (both type and status)
			var foundTypes []string
			for _, tag := range tags {
				if tag.Type != "" {
					foundTypes = append(foundTypes, tag.Type)
				}
				if tag.Status != "" {
					foundTypes = append(foundTypes, tag.Status)
				}
			}

			// Check expected tag types/statuses
			if len(foundTypes) != len(tc.expectTypes) {
				t.Errorf("Expected %v, got %v", tc.expectTypes, foundTypes)
			}

			// Verify clean description doesn't contain tags
			for _, expectedType := range tc.expectTypes {
				tagPattern := fmt.Sprintf("[%s]", expectedType)
				if contains(cleanDesc, tagPattern) {
					t.Errorf("Clean description should not contain %s, got: %s", tagPattern, cleanDesc)
				}
			}

			fmt.Printf("✓ %s: %d tags found, clean desc: '%s'\n", tc.name, len(tags), cleanDesc)
		})
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr || 
		   (len(s) > len(substr) && contains(s[1:], substr))
}