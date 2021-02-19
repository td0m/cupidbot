package cupidbot

// MatchMaker defines an interface that matches users together
type MatchMaker interface {
	Match(Users) []Match
}
