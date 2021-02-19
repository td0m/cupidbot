package matchmaker

import "github.com/td0m/cupidbot"

// All creates all match makers
// this is used in benchmarks and unit testing to test all match maker implementations
func All() []cupidbot.MatchMaker {
	return []cupidbot.MatchMaker{
		NewSimpleMatcher(),
		NewPriorityMatcher(),
	}
}

func contains(s []cupidbot.ID, e cupidbot.ID) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func matchesContain(matches []cupidbot.Match, a cupidbot.ID) bool {
	ids := make([]cupidbot.ID, 2*len(matches))
	for i, m := range matches {
		ids[i*2] = m.A
		ids[i*2+1] = m.B
	}
	return contains(ids, a)
}
