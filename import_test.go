package streamable

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ImportVideo(t *testing.T) {
	videoUrl := "https://archive.org/download/Windows7WildlifeSampleVideo/Wildlife.wmv"
	res, err := ImportVideoFromUrl(videoUrl)

	assert.Nil(t, err)
	assert.Equal(t, 1, res.Status)
	assert.NotEqual(t, "", res.Shortcode)
	fmt.Printf("%v\n", res)
}

func Test_ImportVideoFromUrlAuthenticated_UsernamePassword(t *testing.T) {
	if os.Getenv("STREAMABLE_USERNAME") == "" || os.Getenv("STREAMABLE_PASSWORD") == "" {
		t.Skip("skipping test; $STREAMABLE_USERNAME or $STREAMABLE_PASSWORD not set")
	}

	creds := Credentials{
		Username: os.Getenv("STREAMABLE_USERNAME"),
		Password: os.Getenv("STREAMABLE_PASSWORD"),
	}

	videoUrl := "https://archive.org/download/Windows7WildlifeSampleVideo/Wildlife.wmv"
	res, err := ImportVideoFromUrlAuthenticated(creds, videoUrl)

	assert.Nil(t, err)
	assert.Equal(t, 1, res.Status)
	assert.NotEqual(t, "", res.Shortcode)
	fmt.Printf("%v\n", res)
}
