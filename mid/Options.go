package mid

type Options struct {
	CORS      bool
	Preflight bool
	Auth      bool
}

func MidOptionsGuest() (options Options) {
	options = Options{
		CORS:      true,
		Preflight: true,
	}
	return options
}

func MidOptionsUser() (options Options) {
	options = Options{
		CORS:      true,
		Preflight: true,
		Auth:      true,
	}
	return options
}