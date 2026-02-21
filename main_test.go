package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isTCPNetworkAddress(t *testing.T) {
	t.Run("given valid/allowed host arg then true is returned", func(t *testing.T) {
		tcs := []string{
			"https://test.com:443",
			"http://test.com:80",
			"ssh://test.com:22",
			"test.com:443",
			"test.com:80",
			"test.com:22",
			"https://test.com",
			"http://test.com",
			"ssh://test.com",
		}
		for _, tc := range tcs {
			assert.True(t, isTCPNetworkAddress(tc), tc)
		}
	})

	t.Run("given not valid host arg then false is returned", func(t *testing.T) {
		tcs := []string{
			"ftp://test.com",
			"test.com",
			"/tmp/certs.pem",
		}
		for _, tc := range tcs {
			assert.False(t, isTCPNetworkAddress(tc), tc)
		}
	})
}

func Test_toTCPNetworkAddress(t *testing.T) {
	t.Run("given valid/allowed host arg then scheme is stripped and port added", func(t *testing.T) {
		tcs := map[string]string{
			"https://test.com:443": "test.com:443",
			"http://test.com:80":   "test.com:80",
			"ssh://test.com:22":    "test.com:22",
			// override port map
			"https://test.com:5443": "test.com:5443",
			"http://test.com:8080":  "test.com:8080",
			"ssh://test.com:2222":   "test.com:2222",
			"test.com:443":          "test.com:443",
			"test.com:80":           "test.com:80",
			"test.com:22":           "test.com:22",
			"https://test.com":      "test.com:443",
			"http://test.com":       "test.com:80",
			"ssh://test.com":        "test.com:22",
		}
		for in, expected := range tcs {
			assert.Equal(t, expected, toTCPNetworkAddress(in))
		}
	})

	t.Run("given not valid host arg then the input is returned", func(t *testing.T) {
		tcs := []string{
			"ftp://test.com",
			"test.com",
			"/tmp/certs.pem",
		}
		for _, tc := range tcs {
			assert.Equal(t, tc, toTCPNetworkAddress(tc))
		}
	})
}
