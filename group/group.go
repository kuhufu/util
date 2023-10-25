package group

type Group struct {
	GroupKey string
	Value
	NextLevels []Group
}
