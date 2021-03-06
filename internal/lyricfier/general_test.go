package lyricfier

import (
	"testing"
)

var titlesTables = []struct {
	in  string
	out string
}{
	{"Cuando Pase El Temblor - Remasterizado 2007", "Cuando Pase El Temblor"},
	{"Girl - Remastered 2009", "Girl"},
	{"Girl - Bonus Track", "Girl"},
	{"Girl - bonus track", "Girl"},
	{"Girl - Live", "Girl"},
	{"Mary Jane - 2015 Remaster", "Mary Jane"},
}

func TestFlagParser(t *testing.T) {
	for _, tt := range titlesTables {
		t.Run(tt.in, func(t *testing.T) {
			s := normalizeTitle(tt.in)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
