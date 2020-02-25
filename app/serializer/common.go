package serializer

import (
	"github.com/carlosdamazio/Stone-REST-API/app/model"
)


func SerializeAccountRequest(account *model.Account) (bool, string) {
	if account.Name == "" {
		return false, "Name cannot be empty or be missing."
	} else if account.Cpf == "" {
		return false, "CPF cannot be empty or be missing."
	} else {
		return true, ""
	}
}

