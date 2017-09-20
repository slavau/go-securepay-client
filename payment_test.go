package securepay

import (
	"testing"
)

func TestPaymentMakeOnceOffSuccess(t *testing.T) {

	client, err := NewClient(nil, "https://test.api.securepay.com.au/")
	if err != nil {
		t.Fail()
	}

	merchantInfo := &MerchantInfo{MerchantId: "ABC0001", Password: "abc123"}
	messageInfo := &MessageInfo{ApiVersion: "ml-4.2", Timeout: 60}

	creditCardInfo := &CreditCardInfo{
		CardNumber:     "4444333322221111",
		ExpiryDate:     "01/20",
		Cvv:            "123",
		CardHolderName: "mmm",
	}

	buyerInfo := &BuyerInfo{
		FirstName:       "James",
		LastName:        "Michel",
		ZipCode:         "3000",
		Town:            "Melbourne",
		BillingCountry:  "Australia",
		DeliveryCountry: "Australia",
		EmailAddress:    "jamesk@securepay.com.au",
		Ip:              "150.101.153.111",
	}

	transaction := &Transaction{
		Id:              1,
		TxnType:         "0",
		TxnSource:       "23",
		TxnChannel:      "0",
		Amount:          "100",
		Currency:        "AUD",
		PurchaseOrderNo: "spnv0026.corporate-test",

		CreditCardInfo: creditCardInfo,
		BuyerInfo:      buyerInfo,
	}

	transactionList := &TransactionList{
		Count:       1,
		Transaction: transaction,
	}

	paymentMessage := &PaymentMessage{
		Transactions: transactionList,
	}

	paymentRequest := &PaymentRequest{
		RequestType:    "Payment",
		MerchantInfo:   merchantInfo,
		MessageInfo:    messageInfo,
		PaymentMessage: paymentMessage,
	}

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
