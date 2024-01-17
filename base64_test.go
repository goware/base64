package base64_test

import (
	"crypto/rand"
	"testing"

	"github.com/goware/base64"
	"github.com/stretchr/testify/require"
)

func TestEncoding(t *testing.T) {
	{
		in := "hello world"
		enc := base64.Base64UrlEncode([]byte(in))

		dec, err := base64.Base64UrlDecode(enc)
		require.NoError(t, err)
		require.Equal(t, in, string(dec))
	}

	{
		in := make([]byte, 30)
		rand.Read(in)
		enc := base64.Base64UrlEncode(in)

		dec, err := base64.Base64UrlDecode(enc)
		require.NoError(t, err)
		require.Equal(t, in, dec)
	}
}
