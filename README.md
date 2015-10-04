go-streamable
=============

A Go client library for [streamable.com](https://streamable.com/).

WIP!

```go
import (
  "fmt"
  "github.com/maxkueng/go-streamable"
)

func main() {
  res, _ := streamable.Upload("sekspron.mp4")

  fmt.Printf("%s\n", res.Shortcode);
}
```
