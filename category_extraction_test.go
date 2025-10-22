package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/avelino/awesome-go/pkg/markdown"
)

func TestCategoryExtractionDoesNotIncludeSubcategories(t *testing.T) {
	testMarkdown := `
## Test Category

This is a test category description.

- [Link A](https://example.com/a) - Description A.
- [Link B](https://example.com/b) - Description B.

### Subcategory One

This is subcategory one.

- [Link C](https://example.com/c) - Description C.
- [Link D](https://example.com/d) - Description D.

### Subcategory Two

This is subcategory two.

- [Link E](https://example.com/e) - Description E.

## Another Category

Another category description.

- [Link F](https://example.com/f) - Description F.
`

	html, err := markdown.ToHTML([]byte(testMarkdown))
	if err != nil {
		t.Fatalf("Failed to convert markdown to HTML: %v", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		t.Fatalf("Failed to create goquery document: %v", err)
	}

	category, err := extractCategory(doc, "#test-category")
	if err != nil {
		t.Fatalf("Failed to extract category: %v", err)
	}

	if category.Title != "Test Category" {
		t.Errorf("Expected category title 'Test Category', got '%s'", category.Title)
	}

	t.Logf("Found %d links in category", len(category.Links))
	for i, link := range category.Links {
		t.Logf("  Link %d: %s - %s", i+1, link.Title, link.URL)
	}

	if len(category.Links) != 2 {
		t.Errorf("Expected 2 links (A, B) but got %d links. Bug: including subcategory links!", len(category.Links))
	}

	expectedLinks := map[string]string{
		"Link A": "https://example.com/a",
		"Link B": "https://example.com/b",
	}

	for _, link := range category.Links {
		if expectedURL, ok := expectedLinks[link.Title]; ok {
			if link.URL != expectedURL {
				t.Errorf("Expected URL %s for link %s, got %s", expectedURL, link.Title, link.URL)
			}
		} else {
			t.Errorf("Unexpected link found: %s (URL: %s). This is from a subcategory!", link.Title, link.URL)
		}
	}
}

func TestCategoryWithoutSubcategories(t *testing.T) {
	testMarkdown := `
## Simple Category

Category with direct links only, no subcategories.

- [Tool A](https://example.com/a) - Tool A description.
- [Tool B](https://example.com/b) - Tool B description.
- [Tool C](https://example.com/c) - Tool C description.
`

	html, err := markdown.ToHTML([]byte(testMarkdown))
	if err != nil {
		t.Fatalf("Failed to convert markdown to HTML: %v", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		t.Fatalf("Failed to create goquery document: %v", err)
	}

	category, err := extractCategory(doc, "#simple-category")
	if err != nil {
		t.Fatalf("Failed to extract simple category: %v", err)
	}

	if category.Title != "Simple Category" {
		t.Errorf("Expected category title 'Simple Category', got '%s'", category.Title)
	}

	if len(category.Links) != 3 {
		t.Errorf("Expected 3 links but got %d", len(category.Links))
	}

	expectedLinks := []string{"Tool A", "Tool B", "Tool C"}
	for i, link := range category.Links {
		if link.Title != expectedLinks[i] {
			t.Errorf("Expected link %d to be '%s', got '%s'", i, expectedLinks[i], link.Title)
		}
	}
}

func TestEmptyCategoryReturnsError(t *testing.T) {
	testMarkdown := `
## Empty Category

This category has no links.

## Next Category

- [Link X](https://example.com/x) - Description X.
`

	html, err := markdown.ToHTML([]byte(testMarkdown))
	if err != nil {
		t.Fatalf("Failed to convert markdown to HTML: %v", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		t.Fatalf("Failed to create goquery document: %v", err)
	}

	_, err = extractCategory(doc, "#empty-category")
	if err == nil {
		t.Error("Expected error for empty category, got nil")
	}

	if !errors.Is(err, errCategoryNoLinks) {
		t.Errorf("Expected errCategoryNoLinks, got: %v", err)
	}
}

func TestCategoryOnlyWithSubcategoriesReturnsError(t *testing.T) {
	testMarkdown := `
## Parent Only Category

This category only has subcategories, no direct links.

### Subcategory A

- [Link A](https://example.com/a) - Description A.

### Subcategory B

- [Link B](https://example.com/b) - Description B.

## Next Category

- [Link C](https://example.com/c) - Description C.
`

	html, err := markdown.ToHTML([]byte(testMarkdown))
	if err != nil {
		t.Fatalf("Failed to convert markdown to HTML: %v", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		t.Fatalf("Failed to create goquery document: %v", err)
	}

	_, err = extractCategory(doc, "#parent-only-category")
	if err == nil {
		t.Error("Expected error for category with only subcategories, got nil")
	}

	if !errors.Is(err, errCategoryNoLinks) {
		t.Errorf("Expected errCategoryNoLinks, got: %v", err)
	}
}

func TestExtractCategoriesWithTOC(t *testing.T) {
	testMarkdown := `
## Contents

<details>
<summary>Expand contents</summary>

- [Test Categories](#test-categories)
  - [Category A](#category-a)
  - [Category B](#category-b)

</details>

## Category A

Description A.

- [Link A1](https://example.com/a1) - Description A1.
- [Link A2](https://example.com/a2) - Description A2.

## Category B

Description B.

- [Link B1](https://example.com/b1) - Description B1.
`

	html, err := markdown.ToHTML([]byte(testMarkdown))
	if err != nil {
		t.Fatalf("Failed to convert markdown to HTML: %v", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		t.Fatalf("Failed to create goquery document: %v", err)
	}

	categories, err := extractCategories(doc)
	if err != nil {
		t.Fatalf("Failed to extract categories: %v", err)
	}

	if len(categories) != 2 {
		t.Errorf("Expected 2 categories, got %d", len(categories))
	}

	catA, ok := categories["#category-a"]
	if !ok {
		t.Error("Category A not found")
	} else {
		if catA.Title != "Category A" {
			t.Errorf("Expected title 'Category A', got '%s'", catA.Title)
		}
		if len(catA.Links) != 2 {
			t.Errorf("Expected 2 links in Category A, got %d", len(catA.Links))
		}
	}

	catB, ok := categories["#category-b"]
	if !ok {
		t.Error("Category B not found")
	} else {
		if len(catB.Links) != 1 {
			t.Errorf("Expected 1 link in Category B, got %d", len(catB.Links))
		}
	}
}

func TestExtractCategoriesSkipsEmptyCategories(t *testing.T) {
	testMarkdown := `
## Contents

<details>
<summary>Expand contents</summary>

- [Categories](#categories)
  - [Good Category](#good-category)
  - [Empty Category](#empty-category)
  - [Another Good Category](#another-good-category)

</details>

## Good Category

Has links.

- [Link 1](https://example.com/1) - Description 1.

## Empty Category

No links here.

## Another Good Category

Also has links.

- [Link 2](https://example.com/2) - Description 2.
`

	html, err := markdown.ToHTML([]byte(testMarkdown))
	if err != nil {
		t.Fatalf("Failed to convert markdown to HTML: %v", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		t.Fatalf("Failed to create goquery document: %v", err)
	}

	categories, err := extractCategories(doc)
	if err != nil {
		t.Fatalf("Failed to extract categories: %v", err)
	}

	// Should only have 2 categories (empty one skipped)
	if len(categories) != 2 {
		t.Errorf("Expected 2 categories (empty skipped), got %d", len(categories))
	}

	if _, ok := categories["#good-category"]; !ok {
		t.Error("Good Category should be present")
	}

	if _, ok := categories["#another-good-category"]; !ok {
		t.Error("Another Good Category should be present")
	}

	if _, ok := categories["#empty-category"]; ok {
		t.Error("Empty Category should be skipped")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
