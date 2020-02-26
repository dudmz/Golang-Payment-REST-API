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

func SerializeTransferRequest(transfer *model.Transfer) (bool, string) {
	if transfer.DestinationAccount == "" || transfer.OriginAccount == "" {
		return false, "Transfer must have origin and destination account."
	} else if transfer.Amount <= 0 {
		return false, "Amount cannot be 0 or less."
	} else {
		return true, ""
	}
}
