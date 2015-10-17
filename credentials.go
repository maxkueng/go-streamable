package streamable

// Credentials holds user credentials used form athenticated requests.
type Credentials struct {
	Username string `json:username`
	Password string `json:password`
}
