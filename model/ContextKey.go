package model

type ContextKey string

func GetDbKey() (key ContextKey) {
	key = "db"
	return key
}

func GetUserKey() (key ContextKey) {
	key = "user"
	return key
}

func GetLocationKey() (key ContextKey) {
	key = "location"
	return key
}
