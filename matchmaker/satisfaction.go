package matchmaker

import "github.com/td0m/cupidbot"

const maxSatisfaction float64 = 1

func Satisfaction(users cupidbot.Users, match cupidbot.Match) float64 {
	total := 0.0
	a := users[match.A]
	b := users[match.B]

	if gendersMatchSearch(a.LookingFor, b.Gender) && gendersMatchSearch(b.LookingFor, a.Gender) {
		total++
	}
	return total * 2 / maxSatisfaction
}

func gendersMatchSearch(seeking cupidbot.Gender, got cupidbot.Gender) bool {
	switch seeking {
	case cupidbot.Male:
		return got == cupidbot.Male
	case cupidbot.Female:
		return got == cupidbot.Female
	default:
		return true
	}
}
