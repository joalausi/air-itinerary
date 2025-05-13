package formatter

import (
	"regexp"
	"testing"
)

func TestCodePattern(t *testing.T) {
	// (\*)?      — необязательная «*» в grup1
	// (?:#{1,2}) — # or ##
	// ([A-Za-z0-9]+) — code in grup2
	re := regexp.MustCompile(`(\*)?(?:#{1,2})([A-Za-z0-9]+)`)
	cases := map[string][2]string{
		"#LAX":    {"", "LAX"},
		"##EDDW":  {"", "EDDW"},
		"*#LHR":   {"*", "LHR"},
		"*##EGLL": {"*", "EGLL"},
	}

	for input, want := range cases {
		groups := re.FindStringSubmatch(input)
		if len(groups) < 3 {
			t.Fatalf("pattern didn't match %q at all", input)
		}
		gotStar, gotCode := groups[1], groups[2]
		if gotStar != want[0] || gotCode != want[1] {
			t.Errorf("for %q: expected [%q,%q], got [%q,%q]",
				input, want[0], want[1], gotStar, gotCode)
		}
	}
}
