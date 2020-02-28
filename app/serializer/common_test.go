package serializer

import (
	"github.com/carlosdamazio/Stone-REST-API/app/model"
	"testing"
)

func TestSerializeAccountRequest_BadRequest_NoName(t *testing.T) {
	testAccount := &model.Account{Name: "", Cpf: "123123213"}
	ok, err := SerializeAccountRequest(testAccount)

	if err != "Name cannot be empty or be missing." || ok == true {
		t.Errorf("Function cannot validate if name is missing.")
	}
}

func TestSerializeAccountRequest_BadRequest_NoCpf(t *testing.T) {
	testAccount := &model.Account{Name: "figueira", Cpf: ""}
	ok, err := SerializeAccountRequest(testAccount)

	if err != "CPF cannot be empty or be missing." || ok == true {
		t.Errorf("Function cannot validate if cpf is missing.")
	}
}

func TestSerializeAccountRequest_Serializable(t *testing.T) {
	testAccount := &model.Account{Name: "figueira", Cpf: "000000000"}
	ok, err := SerializeAccountRequest(testAccount)

	if err != "" || ok == false {
		t.Errorf("Should serialize valid account!")
	}
}

func TestSerializeTransferRequest_BadRequest_No_Origin(t *testing.T) {
	testTransfer := &model.Transfer{OriginAccount:"", DestinationAccount:"e213d32fae2231", Amount:30}
	ok, err := SerializeTransferRequest(testTransfer)

	if err != "Transfer must have origin and destination account." || ok != false {
		t.Errorf("Shoudn't serialize if it doesn't have origin account.")
	}
}

func TestSerializeTransferRequest_BadRequest_No_Destination(t *testing.T) {
	testTransfer := &model.Transfer{OriginAccount:"e213f1e2b3f", DestinationAccount:"", Amount:30}
	ok, err := SerializeTransferRequest(testTransfer)

	if err != "Transfer must have origin and destination account." || ok != false {
		t.Errorf("Shoudn't serialize if it doesn't have origin account.")
	}
}

func TestSerializeTransferRequest_BadRequest_Amount_Zero(t *testing.T) {
	testTransfer := &model.Transfer{OriginAccount:"e213f1e2b3f", DestinationAccount:"eef2213e21eba", Amount:0}
	ok, err := SerializeTransferRequest(testTransfer)

	if err != "Amount cannot be 0 or less." || ok != false {
		t.Errorf("Shoudn't serialize if amount in transfer is zero.")
	}
}

func TestSerializeTransferRequest_BadRequest_Amount_Negative(t *testing.T) {
	testTransfer := &model.Transfer{OriginAccount:"e213f1e2b3f", DestinationAccount:"eef2213e21eba", Amount:0}
	ok, err := SerializeTransferRequest(testTransfer)

	if err != "Amount cannot be 0 or less." || ok != false {
		t.Errorf("Shoudn't serialize if amount in transfer is a negative value.")
	}
}

func TestSerializeTransferRequest_Serializable(t *testing.T) {
	testTransfer := &model.Transfer{OriginAccount:"e213f1e2b3f", DestinationAccount:"eef2213e21eba", Amount:15}
	ok, err := SerializeTransferRequest(testTransfer)

	if err != "" || ok != true {
		t.Errorf("This should be serializable.")
	}
}

