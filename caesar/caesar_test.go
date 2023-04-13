package caesar

import "testing"

func TestCaesarCipher(t *testing.T) {
	t.Run("Test an empty message to cipher", func(t *testing.T) {
		test := struct {
			Message string
			Key     int
			Want    string
		}{
			Message: "",
			Key:     1,
			Want:    "",
		}

		if got, _ := CaesarCipher(test.Message, test.Key); string(got) != test.Want {
			t.Fatalf("Wanted %s, but got %s\n", test.Want, got)
		}
	})

	t.Run("Test failure with a negative cipher key value", func(t *testing.T) {
		test := struct {
			Message string
			Key     int
			Want    string
		}{
			Message: "abcdefghijk",
			Key:     -1,
			Want:    "cipher key cannot be lower than 0: -1",
		}

		if _, err := CaesarCipher(test.Message, test.Key); err.Error() != test.Want {
			t.Fatalf("Expected %s, but got %s\n", test.Want, err.Error())
		}
	})

	t.Run("Test some non empty messages", func(t *testing.T) {
		tt := []struct {
			Message string
			Key     int
			Want    string
		}{
			{
				Message: "abc",
				Key:     0,
				Want:    "abc",
			},
			{
				Message: "abc",
				Key:     1,
				Want:    "bcd",
			},
			{
				Message: "rst",
				Key:     5,
				Want:    "wxy",
			},
			{
				Message: "*A*b--cd*e-F**G**",
				Key:     1,
				Want:    "*B*c--de*f-G**H**",
			},
		}

		for _, test := range tt {
			if got, _ := CaesarCipher(test.Message, test.Key); string(got) != test.Want {
				t.Fatalf("Wanted %s, but got %s\n", test.Want, got)
			}
		}
	})

}
