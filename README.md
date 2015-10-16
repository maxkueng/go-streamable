go-streamable
=============

[![Build Status](https://travis-ci.org/maxkueng/go-streamable.svg)](https://travis-ci.org/maxkueng/go-streamable)
[![GoDoc](https://godoc.org/github.com/maxkueng/go-streamable?status.svg)](https://godoc.org/github.com/maxkueng/go-streamable)

A Go client library for [streamable.com](https://streamable.com/).

WIP!

Examples:

```go
import (
  "fmt"
  "github.com/maxkueng/go-streamable"
)

func main() {
  res, _ := streamable.UploadVideo("selfie.mp4")

  fmt.Printf("%s\n", res.Shortcode);
}
```

```go
import (
  "fmt"
  "github.com/maxkueng/go-streamable"
)

func main() {
  creds := streamable.Credentials{
    Username: "user4",
    Password: "secret",
  }

  res, _ := streamable.UploadVideoAuthenticated(creds, "selfie.mp4")

  fmt.Printf("%s\n", res.Shortcode);
}
```
