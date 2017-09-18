package securepay

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)


type PaymentService struct {
	client *Client
}

type PaymentRequest struct {
	XMLName				xml.Name	`xml:"SecurePayMessage"`
	MessageInfo			*MessageInfo
	MerchantInfo		*MerchantInfo
//	RequestType			string		`xml:"RequestType"`
//	PaymentMessage		*PaymentMessage
//	TxnList				[]PaymentPransactions
}

type MessageInfo struct {
	Timeout				int			`xml:"timeoutValue"`
	ApiVersion			string		`xml:"apiVersion"`
}

type MerchantInfo struct {
	MerchantId			string		`xml:"merchantID"`
	Password			string		`xml:"password"`
}

type PaymentResponse struct {
}

func (r *PaymentService) Create(paymentRequest *PaymentRequest) (*PaymentResponse, error) {

	req, err := r.client.NewRequest("POST", "xmlapi/payment", paymentRequest)
	if err != nil {
		return nil, err
	}

	res, err := r.client.Perform(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read the returned data")
	}

	fmt.Println(string(responseData))
	paymentResponse := new(PaymentResponse)
	err = xml.Unmarshal(responseData, paymentResponse)
	if err != nil {
		return nil, fmt.Errorf("could not parse response data")
	}

	return paymentResponse, nil
}
