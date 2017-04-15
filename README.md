go-streamable
=============

[![Build Status](https://travis-ci.org/maxkueng/go-streamable.svg)](https://travis-ci.org/maxkueng/go-streamable)
[![GoDoc](https://godoc.org/github.com/maxkueng/go-streamable?status.svg)](https://godoc.org/github.com/maxkueng/go-streamable)

An unofficial Go client library for [streamable.com](https://streamable.com/).

## Installation

```sh
$ go get github.com/maxkueng/go-streamable
```

## Features

 - Upload videos from a local file
 - Upload videos from a remote URL
 - Retreive information about a video
 - Authenticated requests

## Examples

Upload a video:

```go
func main() {
  client := streamable.New()
  info, err := client.UploadVideo("selfie.mp4")
  if err != nil {
    panic(err)
  }
  
  fmt.Printf("%s\n", info.Shortcode);
}
```

Upload a video with authentication:

```go
func main() {
  client := streamable.New()
  client.SetCredentials("user", "secret")

  info, err := client.UploadVideo("selfie.mp4")
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s\n", info.Shortcode);
}
```

Upload a video from a remote URL:

```go
func main() {
	videoURL := "https://archive.org/download/Windows7WildlifeSampleVideo/Wildlife.wmv"

  client := streamable.New()
  info, err := client.UploadVideoFromURL(videoURL)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s\n", info.Shortcode);
}
```

Receive information about a video:

```go
func main() {
  shortcode := "ifjh"

  client := streamable.New()
  info, err := client.GetVideo(shortcode)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s\n", info.ThumbnailURL)
}
```

## License

MIT
