package utils

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(1, "admin")
	if err != nil {
		t.Error(err)
	}
	//t.Log(token)
	fmt.Println(token)
	claims, err := ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	//t.Log(claims)
	fmt.Println(claims)
}
