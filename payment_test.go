package securepay

import (
	"testing"
)

func TestPaymentMakeOnceOffSuccess(t *testing.T) {

	client, err := NewClient(nil, "https://test.api.securepay.com.au/")
	if err != nil {
		t.Fail()
	}
	paymentRequest, _ := client.Payment.BuildPaymentRequest("ABC0001", "abc123",
		"spnv0026.corporate-test", "100", "4444333322221111", "01", "20")

	paymentResponse, err := client.Payment.Create(paymentRequest)
	if err != nil {
		t.Fail()
	}
	response := *paymentResponse.ResponseDetails.TxnList.ResponseTransaction
	if response.PurchaseOrderNo != "spnv0026.corporate-test" {
		t.Error("Purchase Order Number does not match")
	}
	if response.Approved != "Yes" {
		t.Error("Transaction is not approved")
	}
	if response.ResponseCode != "00" {
		t.Error("Response code is not successful")
	}
	if response.ResponseText != "Approved" {
		t.Error("Response text is not approved")
	}
}
