package model

import (
	"testing"
)

func TestAccount_Save_No_Data(t *testing.T) {
	account := &Account{}

	account.Save()

	if account.Balance != 30 {
		t.Errorf("Account must start with 30 taokeis.")
	}
}

func TestAccount_Save_No_Override_Balance(t *testing.T) {
	value := 20000000.00
	account := &Account{Balance: value}

	account.Save()

	if account.Balance != 30 {
		t.Errorf("Account must not override balance when instanced.")
	}
}
