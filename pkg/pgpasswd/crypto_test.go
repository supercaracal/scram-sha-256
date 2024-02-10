package pgpasswd

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	cases := []struct {
		raw []byte
		err error
	}{
		{[]byte("foo"), nil},
		{[]byte(""), nil},
		{nil, nil},
	}

	for n, c := range cases {
		if _, err := Encrypt(c.raw); err != nil && c.err == nil {
			t.Errorf("%d: %s", n, err)
		} else if err == nil && c.err != nil {
			t.Errorf("%d: no error", n)
		}
	}
}
