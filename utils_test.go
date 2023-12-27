package main

import "testing"

func TestParseExpirationFull(t *testing.T) { // test parseExpirationFull with all valid separator
	result := parseExpiration("2d1h3m47s")
	correctValue := 176627
	if result != correctValue {
		t.Fatal("Error in parseExpirationFull, want : ", correctValue, "got : ", result)
	}
}

func TestParseExpirationMissing(t *testing.T) { // test parseExpirationFull with all valid separator
	result := parseExpiration("1h47s")
	correctValue := 3647
	if result != correctValue {
		t.Fatal("Error in parseExpirationFull, want : ", correctValue, "got : ", result)
	}
}

func TestParseExpirationNull(t *testing.T) { // test parseExpirationFull with all valid separator
	result := parseExpiration("0")
	correctValue := 0
	if result != correctValue {
		t.Fatal("Error in parseExpirationFull, want : ", correctValue, "got : ", result)
	}
}

func TestParseExpirationNegative(t *testing.T) { // test parseExpirationFull with all valid separator
	result := parseExpiration("-42h1m1d4s")
	correctValue := -1
	if result != correctValue {
		t.Fatal("Error in parseExpirationFull, want : ", correctValue, "got : ", result)
	}
}

func TestParseExpirationInvalid(t *testing.T) { // test parseExpirationFull with all valid separator
	result := parseExpiration("8h42h1m1d4s")
	correctValue := -1
	if result != correctValue {
		t.Fatal("Error in parseExpirationFull, want : ", correctValue, "got : ", result)
	}
}
