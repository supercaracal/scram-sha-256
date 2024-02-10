package main

import (
	"testing"
)

func TestGetRawPassword(t *testing.T) {
	cases := []struct {
		args []string
		want string
		err  error
	}{
		{[]string{}, "", nil},
		{[]string{""}, "", nil},
		{[]string{"", ""}, "", nil},
		{[]string{"", "dummy"}, "dummy", nil},
	}

	for n, c := range cases {
		if got, err := getRawPassword(c.args); c.err == nil && err != nil {
			t.Errorf("%d: %s", n, err)
		} else if c.err != nil && err == nil {
			t.Errorf("%d: no error: %s", n, c.err)
		} else if string(got) != c.want {
			t.Errorf("%d: want: %s, got: %s", n, c.want, got)
		}
	}
}
