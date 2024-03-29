package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"

	"github.com/td0m/cupidbot"
	"github.com/td0m/cupidbot/matchmaker"
)

const (
	sampleCount = 100
	roundCount  = 10
	userCount   = 80
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	scores := map[string][]float64{}
	for i := 0; i < sampleCount; i++ {
		users := cupidbot.Users{}
		for i := 0; i < userCount; i++ {
			users[cupidbot.ID(fmt.Sprintf("user%d", i))] = randomUser()
		}
		for _, mm := range matchmaker.All() {
			matches := mm.Match(users)
			dayScores := []float64{}
			for i := 0; i < roundCount; i++ {
				dayScores = append(dayScores, combinedSatisfaction(users, matches)...)
			}
			t := getType(mm)
			scores[t] = append(scores[t], mean(dayScores))
		}
	}
	// percentage of the people who were satisfied with their match
	for k, vs := range scores {
		fmt.Printf("%s: %0.2f%s\n", k, 100*median(vs), "%")
	}
}
func mean(vs []float64) float64 {
	return sum(vs) / float64(len(vs))
}
func median(vs []float64) float64 {
	sort.Float64s(vs)
	return vs[len(vs)/2]
}

func sum(vs []float64) float64 {
	total := 0.0
	for _, v := range vs {
		total += v
	}

	return total
}

func randomUser() cupidbot.UserInfo {
	gender := randomGender()
	return cupidbot.UserInfo{
		Gender:     gender,
		LookingFor: randomGenderTarget(gender),
	}
}

func randomGenderTarget(g cupidbot.Gender) cupidbot.Gender {
	switch g {
	case cupidbot.Male:
		return random([]randomRow{
			{cupidbot.Male, 2},
			{cupidbot.Female, 96},
			{cupidbot.Undefined, 2},
		}).(cupidbot.Gender)
	case cupidbot.Female:
		return random([]randomRow{
			{cupidbot.Male, 96},
			{cupidbot.Female, 2},
			{cupidbot.Undefined, 2},
		}).(cupidbot.Gender)
	default:
		panic("")
	}
}

type randomRow struct {
	Value       interface{}
	Probability int
}

func random(rows []randomRow) interface{} {
	arr := []interface{}{}
	for _, row := range rows {
		for n := 0; n < row.Probability; n++ {
			arr = append(arr, row.Value)
		}
	}
	return arr[rand.Intn(len(arr))]
}

func randomGender() cupidbot.Gender {
	genders := []cupidbot.Gender{cupidbot.Male, cupidbot.Female}
	return genders[rand.Intn(len(genders))]
}

func combinedSatisfaction(users cupidbot.Users, matches []cupidbot.Match) []float64 {
	all := []float64{}
	for _, m := range matches {
		all = append(all, matchmaker.Satisfaction(users, m))
	}
	return all
}

const MAX = 1

func getType(myvar interface{}) string {
	t := reflect.TypeOf(myvar)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}
