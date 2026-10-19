package main

import (
	"os"
	"reflect"
	"testing"
)

func TestClassify(t *testing.T) {
	tests := []struct {
		name    string
		patch   string
		kind    prKind
		added   []string
		removed []string
	}{
		{
			name: "new package added",
			patch: `@@ -1,2 +1,3 @@
 - [alpha](https://github.com/org/alpha) - Alpha.
+- [beta](https://github.com/org/beta) - Beta.
 - [gamma](https://github.com/org/gamma) - Gamma.`,
			kind:  kindPackage,
			added: []string{"https://github.com/org/beta"},
		},
		{
			name: "package removed",
			patch: `@@ -1,3 +1,2 @@
 - [alpha](https://github.com/org/alpha) - Alpha.
-- [beta](https://github.com/org/beta) - Beta.
 - [gamma](https://github.com/org/gamma) - Gamma.`,
			kind:    kindRemoval,
			removed: []string{"https://github.com/org/beta"},
		},
		{
			name: "formatting only",
			patch: `@@ -1,4 +1,4 @@
-* [alpha](https://github.com/org/alpha) - Alpha. 
+- [alpha](https://github.com/org/alpha) - Alpha.
-- [beta](https://github.com/org/beta)  - Beta
+- [beta](https://github.com/org/beta/) - Beta.
 - [gamma](https://github.com/org/gamma) - Gamma.
-
`,
			kind: kindFormatting,
		},
		{
			name: "entry moved to its alphabetical place",
			patch: `@@ -1,3 +1,3 @@
+- [Wisp](https://github.com/wisp-trading/wisp) - Trading.
 - [ymsdk](https://github.com/rekurt/ymsdk) - SDK.
-   - [Wisp](https://github.com/wisp-trading/wisp) - Trading.`,
			kind: kindFormatting,
		},
		{
			name: "mixed add and remove",
			patch: `@@ -1,3 +1,3 @@
 - [alpha](https://github.com/org/alpha) - Alpha.
-- [beta](https://github.com/org/beta) - Beta.
+- [delta](https://github.com/org/delta) - Delta.
 - [gamma](https://github.com/org/gamma) - Gamma.`,
			kind:    kindPackage,
			added:   []string{"https://github.com/org/delta"},
			removed: []string{"https://github.com/org/beta"},
		},
		{
			name: "repository moved to a new URL",
			patch: `@@ -1 +1 @@
-- [finder](https://github.com/old/finder) - Picker.
+- [finder](https://github.com/new/finder) - Picker.`,
			kind:    kindPackage,
			added:   []string{"https://github.com/new/finder"},
			removed: []string{"https://github.com/old/finder"},
		},
		{
			name: "new link without a list bullet still counts",
			patch: `@@ -1 +1,2 @@
 - [alpha](https://github.com/org/alpha) - Alpha.
+++ [sneaky](https://github.com/org/sneaky) - Sneaky.`,
			kind:  kindPackage,
			added: []string{"https://github.com/org/sneaky"},
		},
		{
			name: "table of contents anchors are not links",
			patch: `@@ -1 +1,2 @@
 - [Actor Model](#actor-model)
+- [New Category](#new-category)`,
			kind: kindFormatting,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := classify(tt.patch)
			if got.Kind != tt.kind {
				t.Fatalf("kind = %s, want %s (added %v, removed %v)", got.Kind, tt.kind, got.Added, got.Removed)
			}
			if !reflect.DeepEqual(got.Added, tt.added) {
				t.Errorf("added = %v, want %v", got.Added, tt.added)
			}
			if !reflect.DeepEqual(got.Removed, tt.removed) {
				t.Errorf("removed = %v, want %v", got.Removed, tt.removed)
			}
		})
	}
}

// The testdata patches are the README.md patches GitHub returned for real PRs.
func TestClassifyRealPRs(t *testing.T) {
	tests := []struct {
		file string
		kind prKind
	}{
		{"testdata/pr6365.patch", kindPackage},    // add rummage
		{"testdata/pr6720.patch", kindPackage},    // add traftest
		{"testdata/pr6766.patch", kindRemoval},    // remove JSON-to-Go and JSON-to-Proto
		{"testdata/pr6777.patch", kindFormatting}, // whitespace, bullets, Wisp moved
	}
	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			data, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			if got := classify(string(data)); got.Kind != tt.kind {
				t.Fatalf("kind = %s, want %s (added %v, removed %v)", got.Kind, tt.kind, got.Added, got.Removed)
			}
		})
	}
}

func TestDecide(t *testing.T) {
	addPatch := "@@ -1 +1,2 @@\n a\n+- [beta](https://github.com/org/beta) - Beta.\n"
	removePatch := "@@ -1,2 +1 @@\n a\n-- [beta](https://github.com/org/beta) - Beta.\n"
	tests := []struct {
		name  string
		files []prFile
		want  bool
	}{
		{"no README change", []prFile{{Filename: "CONTRIBUTING.md", Patch: addPatch}}, false},
		{"README adds a package", []prFile{{Filename: "README.md", Patch: addPatch}}, true},
		{"README removes a package", []prFile{{Filename: "README.md", Patch: removePatch}}, false},
		{"README without a patch runs the checks", []prFile{{Filename: "README.md"}}, true},
		{"README add next to other files", []prFile{{Filename: "go.mod"}, {Filename: "README.md", Patch: addPatch}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, reason := decide(tt.files); got != tt.want {
				t.Fatalf("decide = %t (%s), want %t", got, reason, tt.want)
			}
		})
	}
}
