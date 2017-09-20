package securepay

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type PaymentService struct {
	client *Client
}

type SurchargeInfo struct {
	Amount string `xml:"amount"`
	Rate   string `xml:"rate"`
	Fee    string `xml:"fee"`
}

type CreditCardInfo struct {
	CardNumber     string `xml:"cardNumber"`
	ExpiryDate     string `xml:"expiryDate"`
	Cvv            string `xml:"cvv"`
	CardHolderName string `xml:"cardHolderName"`
	XID            string `xml:"xID"`
	CAVV           string `xml:"CAVV"`
	SLI            string `xml:"SLI"`
	PARes          string `xml:"PARes"`
	VERes          string `xml:"VERes"`
	MpiECI         string `xml:"MpiECI"`
}

type BuyerInfo struct {
	FirstName       string `xml:"firstName"`
	LastName        string `xml:"lastName"`
	ZipCode         string `xml:"zipCode"`
	Town            string `xml:"town"`
	BillingCountry  string `xml:"billingCountry"`
	DeliveryCountry string `xml:"deliveryCountry"`
	EmailAddress    string `xml:"emailAddress"`
	Ip              string `xml:"ip"`
}

type Transaction struct {
	Id              int    `xml:"id,attr"`
	TxnKey          string `xml:"txnKey"`
	TxnType         string `xml:"txnType"`
	TxnSource       string `xml:"txnSource"`
	TxnChannel      string `xml:"txnChannel"`
	Amount          string `xml:"amount"`
	Currency        string `xml:"currency"`
	PurchaseOrderNo string `xml:"purchaseOrderNo"`
	PayorId         string `xml:"payorId"`
	TxnId           string `xml:"txnId"`
	PreauthId       string `xml:"preauthId"`

	SurchargeInfo  *SurchargeInfo  `xml:"SurchargeInfo"`
	CreditCardInfo *CreditCardInfo `xml:"CreditCardInfo"`
	BuyerInfo      *BuyerInfo      `xml:"BuyerInfo"`
}

type TransactionList struct {
	Count       int          `xml:"count,attr"`
	Transaction *Transaction `xml:"Txn"`
}

type PaymentMessage struct {
	Transactions *TransactionList `xml:"TxnList"`
}

type PaymentRequest struct {
	XMLName        xml.Name `xml:"SecurePayMessage"`
	MessageInfo    *MessageInfo
	MerchantInfo   *MerchantInfo
	RequestType    string          `xml:"RequestType"`
	PaymentMessage *PaymentMessage `xml:"Payment"`
}

type MessageInfo struct {
	Timeout    int    `xml:"timeoutValue"`
	ApiVersion string `xml:"apiVersion"`
}

type MerchantInfo struct {
	MerchantId string `xml:"merchantID"`
	Password   string `xml:"password"`
}

type Status struct {
	StatusCode        string `xml:"statusCode"`
	StatusDescription string `xml:"statusDescription"`
}

type ResponseTransaction struct {
	PurchaseOrderNo string `xml:"purchaseOrderNo"`
	Approved        string `xml:"approved"`
	ResponseCode    string `xml:"responseCode"`
	ResponseText    string `xml:"responseText"`
	SettlementDate  string `xml:"settlementDate"`
	TxnID           string `xml:"txnID"`
}

type ResponseTxnList struct {
	Count               int                  `xml:"count,attr"`
	ResponseTransaction *ResponseTransaction `xml:"Txn"`
}

type ResponseDetails struct {
	TxnList *ResponseTxnList `xml:"TxnList"`
}

type PaymentResponse struct {
	Status          *Status          `xml:"Status"`
	ResponseDetails *ResponseDetails `xml:"Payment"`
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

	paymentResponse := new(PaymentResponse)
	err = xml.Unmarshal(responseData, paymentResponse)
	if err != nil {
		return nil, fmt.Errorf("could not parse response data")
	}

	return paymentResponse, nil
}
