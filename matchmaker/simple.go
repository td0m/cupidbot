package matchmaker

import (
	"github.com/td0m/cupidbot"
)

type SimpleMatcher struct {
}

func NewSimpleMatcher() *SimpleMatcher {
	return &SimpleMatcher{}
}

func (s *SimpleMatcher) Match(users cupidbot.Users) []cupidbot.Match {
	matches := []cupidbot.Match{}
	for id := range users {
		for id2 := range users {
			if id != id2 && !matchesContain(matches, id) && !matchesContain(matches, id2) {
				matches = append(matches, cupidbot.Match{A: id, B: id2})
			}
		}
	}
	return matches
}
