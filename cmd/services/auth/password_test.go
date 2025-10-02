package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password %v", err)
	}

	if hash == "" {
		t.Error("Error, hashed password cannot be emptyh")
	}
	
	if hash == "password"{
		t.Error("Hashed password cannot be same as plain password")
	}
}

func TestComparePassword(t *testing.T){
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("error hashing password %v", err)
	}

	same := ComparePasswords(hash, []byte(password))

	if !same {
		t.Error("Expect passwords have to match")
	}

	notsame := ComparePasswords(hash, []byte("notpassword"))

	if notsame {
		t.Error("Expect passwords should not match")
	}
}