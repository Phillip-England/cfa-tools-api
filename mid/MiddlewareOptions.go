package mid

type MiddlewareOptions struct {
	CORS      bool
	Preflight bool
	Auth      bool
}

func MidOptionsGuest() (options MiddlewareOptions) {
	options = MiddlewareOptions{
		CORS:      true,
		Preflight: true,
	}
	return options
}

func MidOptionsUser() (options MiddlewareOptions) {
	options = MiddlewareOptions{
		CORS:      true,
		Preflight: true,
		Auth:      true,
	}
	return options
}