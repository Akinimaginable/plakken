package utils

import (
	"errors"
	"testing"
)

func TestCheckCharNotRedundantTrue(t *testing.T) { // Test checkCharRedundant with redundant char
	want := true
	got := checkCharRedundant("2d1h3md4h7s", "h")
	if got != want {
		t.Fatal("Error in parseExpirationFull, want : ", want, "got : ", got)
	}
}

func TestCheckCharNotRedundantFalse(t *testing.T) { // Test checkCharRedundant with not redundant char
	want := false
	got := checkCharRedundant("2d1h3m47s", "h")
	if got != want {
		t.Fatal("Error in parseExpirationFull, want : ", want, "got : ", got)
	}
}

func TestParseExpirationFull(t *testing.T) { // test parseExpirationFull with all valid separator
	result, _ := ParseExpiration("2d1h3m47s")
	correctValue := 176627
	if result != correctValue {
		t.Fatal("Error in parseExpirationFull, want : ", correctValue, "got : ", result)
	}
}

func TestParseExpirationMissing(t *testing.T) { // test parseExpirationFull with all valid separator
	result, _ := ParseExpiration("1h47s")
	correctValue := 3647
	if result != correctValue {
		t.Fatal("Error in ParseExpirationFull, want : ", correctValue, "got : ", result)
	}
}

func TestParseExpirationWithCaps(t *testing.T) { // test parseExpirationFull with all valid separator
	result, _ := ParseExpiration("2D1h3M47s")
	correctValue := 176627
	if result != correctValue {
		t.Fatal("Error in parseExpirationFull, want : ", correctValue, "got : ", result)
	}
}

func TestParseExpirationNull(t *testing.T) { // test ParseExpirationFull with all valid separator
	result, _ := ParseExpiration("0")
	correctValue := 0
	if result != correctValue {
		t.Fatal("Error in ParseExpirationFull, want: ", correctValue, "got: ", result)
	}
}

func TestParseExpirationNegative(t *testing.T) { // test ParseExpirationFull with all valid separator
	_, got := ParseExpiration("-42h1m4s")
	want := &ParseExpirationError{}
	if !errors.As(got, &want) {
		t.Fatal("Error in ParseExpirationFull, want : ", want, "got : ", got)
	}
}

func TestParseExpirationInvalid(t *testing.T) { // test ParseExpirationFull with all valid separator
	_, got := ParseExpiration("8h42h1m1d4s")
	want := &ParseExpirationError{}
	if !errors.As(got, &want) {
		t.Fatal("Error in ParseExpirationFull, want : ", want, "got : ", got)
	}

}

func TestParseExpirationInvalidRedundant(t *testing.T) { // test ParseExpirationFull with all valid separator
	_, got := ParseExpiration("8h42h1m1h4s")
	want := &ParseExpirationError{}
	if !errors.As(got, &want) {
		t.Fatal("Error in ParseExpirationFull, want : ", want, "got : ", got)
	}
}

func TestParseExpirationInvalidTooHigh(t *testing.T) { // test ParseExpirationFull with all valid separator
	_, got := ParseExpiration("2d1h3m130s")
	want := &ParseExpirationError{}
	if !errors.As(got, &want) {
		t.Fatal("Error in ParseExpirationFull, want : ", want, "got : ", got)
	}
}

func TestValidKey(t *testing.T) { // test ValidKey with a valid key
	got := ValidKey("ab_a-C42")
	want := true

	if got != want {
		t.Fatal("Error in ValidKey, want : ", want, "got : ", got)
	}
}

func TestInValidKey(t *testing.T) { // test ValidKey with invalid key
	got := ValidKey("ab_?a-C42")
	want := false

	if got != want {
		t.Fatal("Error in ValidKey, want : ", want, "got : ", got)
	}
}
