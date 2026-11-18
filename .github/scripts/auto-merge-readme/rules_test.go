package main

import (
	"errors"
	"testing"
	"time"
)

func TestOnlyReadme(t *testing.T) {
	cases := []struct {
		files []string
		want  bool
	}{
		{[]string{"README.md"}, true},
		{nil, false},
		{[]string{"readme.md"}, false},
		{[]string{"README.md", "go.mod"}, false},
		{[]string{"docs/README.md"}, false},
	}
	for _, c := range cases {
		if got := onlyReadme(c.files); got != c.want {
			t.Errorf("onlyReadme(%v) = %v, want %v", c.files, got, c.want)
		}
	}
}

func TestAddedRepos(t *testing.T) {
	ok := "@@ -10,6 +10,7 @@ Libraries\n - [a](https://github.com/x/a) - A.\n+- [gotermfx](https://github.com/foo/gotermfx) - Terminal effects.\n - [b](https://github.com/x/b) - B.\n"
	got, err := addedRepos(ok)
	if err != nil || len(got) != 1 || got[0].String() != "foo/gotermfx" {
		t.Fatalf("addedRepos(ok) = %v, %v", got, err)
	}
	two := "@@ -1 +1,3 @@\n+- [a](https://github.com/o/a) - A.\n+* [b](https://github.com/o/b.go/) - B.\n"
	if got, err := addedRepos(two); err != nil || len(got) != 2 || got[1].String() != "o/b.go" {
		t.Fatalf("addedRepos(two) = %v, %v", got, err)
	}
	bad := map[string]string{
		"removal":     "@@\n-- [a](https://github.com/o/a) - A.\n+- [a](https://github.com/o/b) - A.\n",
		"no entry":    "@@\n context\n",
		"gitlab":      "+- [a](https://gitlab.com/o/a) - A.\n",
		"subpath":     "+- [a](https://github.com/o/a/tree/main/x) - A.\n",
		"http":        "+- [a](http://github.com/o/a) - A.\n",
		"prose":       "+Some new paragraph.\n",
		"blank":       "+\n",
		"no desc":     "+- [a](https://github.com/o/a)\n",
		"duplicate":   "+- [a](https://github.com/o/a) - A.\n+- [A](https://github.com/O/A) - A.\n",
		"query":       "+- [a](https://github.com/o/a?x=1) - A.\n",
		"header line": "+++ b/README.md\n",
	}
	for name, p := range bad {
		if got, err := addedRepos(p); err == nil {
			t.Errorf("%s: want error, got %v", name, got)
		}
	}
}

func TestOldEnough(t *testing.T) {
	now := time.Date(2026, 10, 12, 12, 0, 0, 0, time.UTC)
	cases := []struct {
		created time.Time
		want    bool
	}{
		{time.Date(2026, 5, 12, 12, 0, 0, 0, time.UTC), true},
		{time.Date(2026, 5, 12, 12, 0, 1, 0, time.UTC), false},
		{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC), false},
		{time.Time{}, false},
		{now.Add(time.Hour), false},
	}
	for _, c := range cases {
		if got := oldEnough(c.created, now); got != c.want {
			t.Errorf("oldEnough(%v) = %v, want %v", c.created, got, c.want)
		}
	}
}

func TestOldEnoughMonthEnd(t *testing.T) {
	d := func(y int, m time.Month, day int) time.Time { return time.Date(y, m, day, 0, 0, 0, 0, time.UTC) }
	cases := []struct {
		created, now time.Time
		want         bool
	}{
		// Plain subtraction from Jul 31 lands on Mar 3 and would accept these.
		{d(2026, 3, 1), d(2026, 7, 31), false},
		{d(2026, 3, 2), d(2026, 7, 31), false},
		{d(2026, 3, 3), d(2026, 7, 31), false},
		{d(2026, 3, 1), d(2026, 8, 1), true},
		{d(2026, 3, 31), d(2026, 8, 30), false},
		{d(2026, 3, 31), d(2026, 8, 31), true},
		// Sep 30 + 5 months clamps to the end of February.
		{d(2025, 9, 30), d(2026, 2, 27), false},
		{d(2025, 9, 30), d(2026, 2, 28), true},
		{d(2027, 9, 30), d(2028, 2, 29), true},
		{d(2027, 9, 30), d(2028, 2, 28), false},
	}
	for _, c := range cases {
		if got := oldEnough(c.created, c.now); got != c.want {
			t.Errorf("oldEnough(%v, %v) = %v, want %v", c.created, c.now, got, c.want)
		}
	}
}

func TestApprovalStatus(t *testing.T) {
	maint := func(u string) (bool, error) { return u == "avelino" || u == "ceriath", nil }
	r := func(u, s, c string) review { return review{u, s, c} }

	if a, err := approvalStatus([]review{r("avelino", "APPROVED", "h")}, "h", maint); err != nil || len(a) != 1 {
		t.Errorf("maintainer approval: %v %v", a, err)
	}
	if a, _ := approvalStatus([]review{r("random", "APPROVED", "h")}, "h", maint); len(a) != 0 {
		t.Errorf("outsider approval counted: %v", a)
	}
	if a, _ := approvalStatus([]review{r("avelino", "APPROVED", "old")}, "h", maint); len(a) != 0 {
		t.Errorf("stale approval counted: %v", a)
	}
	if a, _ := approvalStatus([]review{r("avelino", "APPROVED", "h"), r("avelino", "DISMISSED", "h")}, "h", maint); len(a) != 0 {
		t.Errorf("dismissed approval counted: %v", a)
	}
	if a, _ := approvalStatus([]review{r("avelino", "APPROVED", "h"), r("avelino", "COMMENTED", "h")}, "h", maint); len(a) != 1 {
		t.Errorf("comment should not reset approval: %v", a)
	}
	if _, err := approvalStatus([]review{r("avelino", "APPROVED", "h"), r("ceriath", "CHANGES_REQUESTED", "h")}, "h", maint); err == nil {
		t.Error("changes requested should block")
	}
	fail := func(string) (bool, error) { return false, errors.New("api down") }
	if _, err := approvalStatus([]review{r("avelino", "APPROVED", "h")}, "h", fail); err == nil {
		t.Error("permission lookup error should block")
	}
}

func TestChecksGreen(t *testing.T) {
	req := []string{"Running test"}
	c := func(n, s, cc string) checkRun { return checkRun{n, s, cc} }
	if err := checksGreen([]checkRun{c("Running test", "completed", "success"), c("Skip", "completed", "skipped"), c("Auto merge", "in_progress", "")}, req, "Auto merge"); err != nil {
		t.Errorf("green: %v", err)
	}
	for name, runs := range map[string][]checkRun{
		"missing":   {c("Other", "completed", "success")},
		"failed":    {c("Running test", "completed", "success"), c("Repository quality checks", "completed", "failure")},
		"running":   {c("Running test", "in_progress", "")},
		"cancelled": {c("Running test", "completed", "cancelled")},
	} {
		if err := checksGreen(runs, req, "Auto merge"); err == nil {
			t.Errorf("%s: want error", name)
		}
	}
}
