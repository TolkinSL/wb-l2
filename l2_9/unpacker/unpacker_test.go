package unpacker

import "testing"

func TestUnpack(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected string
		wantErr bool
	}{
		{"Simple", "a4bc2d5e", "aaaabccddddde", false},
		{"No digits", "abcd", "abcd", false},
		{"Invalid (starts with digit)", "45", "", true},
		{"Empty string", "", "", false},
		{"Escaped digit", `qwe\4\5`, "qwe45", false},
		{"Escaped digit with repeat", `qwe\45`, "qwe44444", false},
		{"Escaped backslash", `qwe\\5`, `qwe\\\\\`, false},
		{"Escaped last", `qwe\\5\`, ``, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T)  {
			res, err := Unpack(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error: %v, got: %v", tc.wantErr, err)
			}
			if res != tc.expected {
				t.Errorf("expected: %s, got: %s", tc.expected, res)
			}
		})
	}
}