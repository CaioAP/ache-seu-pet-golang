package test

import (
	"testing"

	"github.com/caioap/ache-seu-pet/backend/internal/domain/entity"
)

func TestComparePassword(t *testing.T) {
	password := "senha321"
	hash, err := entity.Hash(password); if err != nil {
		t.Fatalf("Error hashing password")
	}
	result := entity.Compare(hash, password)
	if result == false {
		t.Errorf("ComparePassword(hash, senha321) FAILED. Expected %v, got %v", true, result)
	} else {
		t.Logf("ComparePassword(hash, senha321) PASSED. Expected %v, got %v", true, result)
	}
}