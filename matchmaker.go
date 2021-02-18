package cupidbot

type MatchMaker interface {
	Match(Users) []Match
}
