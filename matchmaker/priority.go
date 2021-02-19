package matchmaker

import (
	"sort"

	"github.com/td0m/cupidbot"
)

// PriorityMatcher doesn't try to always optimize for maximum satisfaction
// instead, it allows people who were unlucky last time to have higher chances this time
// it also makes sure not to match the same people again if we don't have to
type PriorityMatcher struct {
	lastSatisfaction map[cupidbot.ID]float64
}

// NewPriorityMatcher creates a new priority-based match maker
func NewPriorityMatcher() *PriorityMatcher {
	return &PriorityMatcher{
		lastSatisfaction: map[cupidbot.ID]float64{},
	}
}

// Match implementation
func (m *PriorityMatcher) Match(users cupidbot.Users) []cupidbot.Match {
	matches := []cupidbot.Match{}
	// most recently dissatisfied users will have priority
	ids := []cupidbot.ID{}
	for id := range users {
		ids = append(ids, id)
	}

	// least satisfied users go first
	sort.Slice(ids, func(i, j int) bool {
		return m.lastSatisfaction[ids[i]] < m.lastSatisfaction[ids[i]]
	})

	for _, id := range ids {
		possibleMatches := []cupidbot.Match{}
		// get all possible matches
		for id2 := range users {
			if id != id2 && !matchesContain(matches, id2) && !matchesContain(matches, id) {
				possibleMatches = append(possibleMatches, cupidbot.Match{A: id, B: id2})
			}
		}

		// if possible match found we add it
		if len(possibleMatches) > 0 {
			bestMatch := possibleMatches[0]
			satisfaction := Satisfaction(users, bestMatch)
			// get the best match
			for _, m := range possibleMatches {
				ms := Satisfaction(users, m)
				if ms >= satisfaction {
					bestMatch = m
					satisfaction = ms
				}
			}
			m.lastSatisfaction[id] = satisfaction
			matches = append(matches, bestMatch)
		}
	}

	return matches
}
