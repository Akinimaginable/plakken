package secret

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"git.gnous.eu/gnouseu/plakken/internal/constant"
	"golang.org/x/crypto/argon2"
)

func TestPasswordFormat(t *testing.T) {
	regex := fmt.Sprintf("\\$argon2id\\$v=%d\\$m=%d,t=%d,p=%d\\$[A-Za-z0-9+/]*\\$[A-Za-z0-9+/]*$", argon2.Version, constant.ArgonMemory, constant.ArgonIterations, constant.ArgonThreads)

	got, err := Password("Password!")
	if err != nil {
		t.Fatal(err)
	}

	result, _ := regexp.MatchString(regex, got)
	if !result {
		t.Fatal("Error in Password, format is not valid "+": ", got)
	}
}

func TestVerifyPassword(t *testing.T) {
	result, err := VerifyPassword("Password!", "$argon2id$v=19$m=65536,t=2,p=4$A+t5YGpyy1BHCbvk/LP1xQ$eNuUj6B2ZqXlGi6KEqep39a7N4nysUIojuPXye+Ypp0")
	if err != nil {
		t.Fatal(err)
	}

	if !result {
		t.Fatal("Error in VerifyPassword, got:", result, "want: ", true)
	}
}

func TestVerifyPasswordInvalid(t *testing.T) {
	result, err := VerifyPassword("notsamepassword", "$argon2id$v=19$m=65536,t=2,p=4$A+t5YGpyy1BHCbvk/LP1xQ$eNuUj6B2ZqXlGi6KEqep39a7N4nysUIojuPXye+Ypp0")
	if err != nil {
		t.Fatal(err)
	}

	if result {
		t.Fatal("Error in VerifyPassword, got:", result, "want: ", false)
	}
}

func TestParseHash(t *testing.T) {
	_, config, err := parseHash("$argon2id$v=19$m=65536,t=2,p=4$A+t5YGpyy1BHCbvk/LP1xQ$eNuUj6B2ZqXlGi6KEqep39a7N4nysUIojuPXye+Ypp0")
	if err != nil {
		t.Fatal(err)
	}
	if config.saltLength != constant.ArgonSaltSize {
		t.Fatal("Error in VerifyPassword: config.saltLength are not correct, go: ", config.saltLength, "want: ", constant.ArgonSaltSize)
	}

	if config.keyLength != constant.ArgonKeyLength {
		t.Fatal("Error in VerifyPassword: config.keyLength are not correct, go: ", config.saltLength, "want: ", constant.ArgonKeyLength)
	}

	if config.threads != constant.ArgonThreads {
		t.Fatal("Error in VerifyPassword: config.threads are not correct, go: ", config.threads, "want: ", constant.ArgonThreads)
	}

	if config.memory != constant.ArgonMemory {
		t.Fatal("Error in VerifyPassword: config.memory are not correct, go: ", config.memory, "want: ", constant.ArgonMemory)
	}

	if config.iterations != constant.ArgonIterations {
		t.Fatal("Error in VerifyPassword: config.iterations are not correct, go: ", config.iterations, "want: ", constant.ArgonIterations)
	}
}

func TestParseBadHashAlgo(t *testing.T) {
	_, _, err := parseHash("$notvalid$v=19$m=65536,t=2,p=4$A+t5YGpyy1BHCbvk/LP1xQ$eNuUj6B2ZqXlGi6KEqep39a7N4nysUIojuPXye+Ypp0")
	want := &parseError{message: "Algorithm is not valid"}
	if !errors.As(err, &want) {
		t.Fatal("Error in parseHash, want :", want, "got: ", err)
	}
}
