package matchmaker

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/td0m/cupidbot"
)

func values(users cupidbot.Users) []cupidbot.UserInfo {
	infos := []cupidbot.UserInfo{}
	for _, v := range users {
		infos = append(infos, v)
	}
	return infos
}

func lookingFor(users []cupidbot.UserInfo, g cupidbot.Gender) []cupidbot.UserInfo {
	filtered := []cupidbot.UserInfo{}
	for _, u := range users {
		if u.LookingFor == g {
			filtered = append(filtered, u)
		}
	}
	return filtered
}
func gender(users []cupidbot.UserInfo, g cupidbot.Gender) []cupidbot.UserInfo {
	filtered := []cupidbot.UserInfo{}
	for _, u := range users {
		if u.Gender == g {
			filtered = append(filtered, u)
		}
	}
	return filtered
}

// you can add more users but please keep these here.
var users = cupidbot.Users{
	"ben": {
		Gender:     cupidbot.Male,
		LookingFor: cupidbot.Female,
	},
	"dom": {
		Gender:     cupidbot.Male,
		LookingFor: cupidbot.Female,
	},
	"manisha": {
		Gender:     cupidbot.Female,
		LookingFor: cupidbot.Male,
	},
	"sarah": {
		Gender:     cupidbot.Female,
		LookingFor: cupidbot.Male,
	},
	"zack": {
		Gender:     cupidbot.Male,
		LookingFor: cupidbot.Undefined,
	},
	"ana": {
		Gender:     cupidbot.Female,
		LookingFor: cupidbot.Male,
	},
	"steve": {
		Gender:     cupidbot.Male,
		LookingFor: cupidbot.Male,
	},
}

func TestAll_Random(t *testing.T) {
	straightMales := lookingFor(gender(values(users), cupidbot.Male), cupidbot.Female)
	if len(straightMales) < 2 {
		t.Fatalf("not enough straight males")
	}
	girlsLookingForGuys := lookingFor(gender(values(users), cupidbot.Female), cupidbot.Male)
	if len(girlsLookingForGuys) < 3 {
		t.Fatalf("not enough girls looking for guys: %d", len(girlsLookingForGuys))
	}
	users := cupidbot.Users{
		"alphamale1": straightMales[0],
		"alphamale2": straightMales[1],
	}
	for i, g := range girlsLookingForGuys {
		users[cupidbot.ID(fmt.Sprintf("girl%d", i))] = g
	}
	male1Dates := []cupidbot.ID{}
	for _, matchmaker := range All() {
		strType := getType(matchmaker)
		iterations := 10
		for i := 0; i < iterations+1; i++ {
			matches := matchmaker.Match(users)
			for _, m := range matches {
				if m.A == "alphamale1" {
					male1Dates = append(male1Dates, m.B)
				}
				if m.B == "alphamale1" {
					male1Dates = append(male1Dates, m.A)
				}
			}
		}
		different := false
		for i := 1; i < iterations; i++ {
			if male1Dates[0] != male1Dates[i] {
				different = true
			}
		}
		if !different {
			t.Errorf("(%s) Matches do not appear to be randomly selected", strType)
		}
	}
}

func TestAll_Matches(t *testing.T) {
	expectedMatchLen := len(users) / 2
	for i := 0; i < 100; i++ {
		for _, matchmaker := range All() {
			matches := matchmaker.Match(users)
			strType := getType(matchmaker)
			if len(matches) != expectedMatchLen {
				t.Errorf("%s expected %d matches, given %d", strType, expectedMatchLen, len(matches))
			}
			for i, m := range matches {
				prevMatched := matches[:i]
				if matchesContain(prevMatched, m.A) {
					t.Errorf("%s matches the same person twice (%s)", strType, m.A)
				}
				if matchesContain(prevMatched, m.B) {
					t.Errorf("%s matches the same person twice (%s)", strType, m.B)
				}
				if m.A == m.B {
					t.Errorf("%s matched %s with themselves", strType, m.A)
				}
			}
		}
	}
}

func getType(myvar interface{}) string {
	t := reflect.TypeOf(myvar)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}
