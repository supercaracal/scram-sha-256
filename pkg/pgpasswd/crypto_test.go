package pgpasswd

import (
	"testing"
)

var (
	dummyPassword = []byte("dummyblahfoobarbaztest")
)

func TestEncrypt(t *testing.T) {
	cases := []struct {
		raw []byte
		err error
	}{
		{[]byte("foo"), nil},
		{[]byte(""), nil},
		{[]byte{}, nil},
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

func BenchmarkEncrypt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Encrypt(dummyPassword); err != nil {
			b.Fatal(err)
		}
	}
}
