package comma

import "testing"

func TestComma(t *testing.T) {
	var tests = []struct {
		input  string
		expect string
	}{
		{"", ""},
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
		{"123456789", "123,456,789"},
		{"1234567890", "1,234,567,890"},
	}

	for _, test := range tests {
		if actual := Comma(test.input); actual != test.expect {
			t.Errorf("Comma(%q) = %q; Expected: %q", test.input, actual, test.expect)
		}
	}
}
