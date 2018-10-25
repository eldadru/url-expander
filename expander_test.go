package url_expander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExpandUrl_ShortenedUrl(t *testing.T) {
	// given
	shortenedUrl := "http://bit.do/eznjc"

	// when
	expandedUrl, err := ExpandUrl(shortenedUrl)

	// then
	assert.Nil(t, err)
	assert.Equal(t, "http://www.google.com", expandedUrl)
}

func TestExpandUrl_UrlAlreadyShortened(t *testing.T) {
	// given
	shortenedUrl := "https://www.google.com"

	// when
	expandedUrl, err := ExpandUrl(shortenedUrl)

	// then
	assert.Nil(t, err)
	assert.Equal(t, shortenedUrl, expandedUrl)
}

func TestExpandUrl_InvalidUrl(t *testing.T) {
	// given
	shortenedUrl := "invalid-url"

	// when
	expandedUrl, err := ExpandUrl(shortenedUrl)

	// then
	assert.NotNil(t, err)
	assert.Equal(t, "", expandedUrl)
}
