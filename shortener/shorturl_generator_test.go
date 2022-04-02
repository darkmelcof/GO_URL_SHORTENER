package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink := GenerateShortLink(initialLink, UserId)

	assert.Equal(t, shortLink, "jTa4L57P")
}

func TestShortLinkGenerator2(t *testing.T) {
	initialLink := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink := GenerateShortLink(initialLink, UserId)

	assert.Equal(t, shortLink, "d66yfx7N")
}

func TestShortLinkGenerator3(t *testing.T) {
	initialLink := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink := GenerateShortLink(initialLink, UserId)

	assert.Equal(t, shortLink, "dhZTayYQ")
}

func TestShortLinkGenerator4(t *testing.T) {
	initialLink := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink := GenerateShortLink(initialLink, "e0dba740-124b-4977-872c-d360239e6b1a")

	assert.Equal(t, shortLink, "4A9YEd8s")
}
