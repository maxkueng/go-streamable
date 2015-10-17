package streamable

import "net/http"

func bytesToString(byteArray []byte) string {
	return string(byteArray)
}

func authenticateHTTPRequest(req *http.Request, creds Credentials) {
	if creds.Username != "" && creds.Password != "" {
		req.SetBasicAuth(creds.Username, creds.Password)
	}
}
