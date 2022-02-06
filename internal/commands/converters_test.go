package commands

import "testing"

func Test_BaseConversions_Valid(t *testing.T) {
	cases := []struct {
		fromBase int
		toBase   int
		from     string
		want     string
	}{
		// To Decimal
		{2, 10, "0000", "0"},
		{2, 10, "1111", "15"},
		{2, 10, "-1111", "-15"},
		{8, 10, "0", "0"},
		{8, 10, "36", "30"},
		{8, 10, "-36", "-30"},
		{16, 10, "0", "0"},
		{16, 10, "1e", "30"},
		{16, 10, "-1e", "-30"},
		// From Decimal
		{10, 2, "15", "1111"},
		{10, 2, "0", "0"},
		{10, 2, "-10", "-1010"},
		{10, 8, "0", "0"},
		{10, 8, "30", "36"},
		{10, 8, "-30", "-36"},
		{10, 16, "0", "0"},
		{10, 16, "30", "1e"},
		{10, 16, "-30", "-1e"},
		// To Binary
		{8, 2, "17", "1111"},
		{8, 2, "0", "0"},
		{8, 2, "-17", "-1111"},
		{16, 2, "0", "0"},
		{16, 2, "f", "1111"},
		{16, 2, "-f", "-1111"},
		// From Binary
		{2, 8, "0", "0"},
		{2, 8, "1111", "17"},
		{2, 8, "-1111", "-17"},
		{2, 16, "0", "0"},
		{2, 16, "1111", "f"},
		{2, 16, "-1111", "-f"},
		// To Octal
		{16, 8, "0", "0"},
		{16, 8, "f", "17"},
		{16, 8, "-f", "-17"},
		// From Octal
		{8, 16, "0", "0"},
		{8, 16, "17", "f"},
		{8, 16, "-17", "-f"},
	}

	for _, c := range cases {
		got, err := convertNumberBase(c.from, c.fromBase, c.toBase)
		if err != nil {
			t.Errorf("Error converting number base: %v", err)
		}

		if c.want != got {
			t.Errorf("Failed to convert number base. Wanted %s, Got %s", c.want, got)
		}
	}
}
