package streamable

// Credentials holds user credentials used for authenticated requests.
type Credentials struct {
	Username string `json:username`
	Password string `json:password`
}
