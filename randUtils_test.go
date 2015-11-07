package randutils

import "testing"

func TestAlphaLower(t *testing.T) {
	for i := 0; i < 1000; i++ {
		buff, err := AlphaLower(uint32(i))
		if err != nil {
			t.Fatal(err)
		}
		for _, c := range buff {
			if c < 'a' || c > 'z' {
				t.Fatalf("The value [ %c ] : should be between 'a' and 'z'", c)
			}
		}
		if len(buff) != i {
			t.Fatalf("The length's buffer should be equal to %v != %v", i, len(buff))
		}
	}
}

func TestAlphaUpper(t *testing.T) {
	for i := 0; i < 1000; i++ {
		buff, err := AlphaUpper(uint32(i))
		if err != nil {
			t.Fatal(err)
		}
		for _, c := range buff {
			if c < 'A' || c > 'Z' {
				t.Fatalf("The value [ %c ] : should be between 'A' and 'Z'", c)
			}
		}
		if len(buff) != i {
			t.Fatalf("The length's buffer should be equal to %v != %v", i, len(buff))
		}
	}
}
