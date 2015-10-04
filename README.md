go-streamable
=============

A Go client library for [streamable.com](https://streamable.com/).

WIP!

Examples:

```go
import (
  "fmt"
  "github.com/maxkueng/go-streamable"
)

func main() {
  res, _ := streamable.Upload("selfie.mp4")

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

  res, _ := streamable.UploadAuthenticated(creds, "selfie.mp4")

  fmt.Printf("%s\n", res.Shortcode);
}
```
