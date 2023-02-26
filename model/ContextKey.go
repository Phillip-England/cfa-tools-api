package model

type ContextKey string

func GetUserKey() (key ContextKey) {
	key = "user"
	return key
}

func GetLocationKey() (key ContextKey) {
	key = "location"
	return key
}
