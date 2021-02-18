package cupidbot

type ID string

type Users map[ID]UserInfo

type UserInfo struct {
	Gender     Gender
	LookingFor Gender
}
