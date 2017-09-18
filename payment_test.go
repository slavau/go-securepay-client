package securepay

import (
	"testing"
	"fmt"
)

func TestPayment_MakeOnceOff_Success(t *testing.T) {

	client, err := NewClient(nil, "https://test.api.securepay.com.au/")
	if err != nil {
		t.Fail()
	}

	merchantInfo := &MerchantInfo{MerchantId: "ABC0001", Password: "abc123"}
	messageInfo := &MessageInfo{ApiVersion: "ml-4.2", Timeout: 60}

	paymentRequest := &PaymentRequest{
		MerchantInfo: merchantInfo,
		MessageInfo:  messageInfo,
	}

	paymentResponse, err := client.Payment.Create(paymentRequest)
	if err != nil {
		t.Fail()
	}
	fmt.Printf("%v+++", paymentResponse)
}
