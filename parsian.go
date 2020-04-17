// Package shaparak
package shaparak

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"

	//"bytes"
	//"encoding/json"
	//"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"

	//"io/ioutil"
	"log"
	//"net/http"
	//"strconv"
)

type Parsian struct {
	config          map[string]interface{}
	Transaction     TransactionInterface
	MerchantID      string
	Sandbox         bool
	APIEndpoint     string
	PaymentEndpoint string
}

func NewParsian(transaction TransactionInterface, config map[string]interface{}, sandbox bool) (*Parsian, error) {
	MerchantID, exists := config["MerchantID"]

	if exists == false {
		return nil, errors.New("MerchantID does not exists in passed Parsian config")
	}
	if len(MerchantID.(string)) != 20 {
		return nil, errors.New("parsian MerchantID must be 20 characters")
	}

	apiEndPoint := "https://pec.shaparak.ir/NewIPGServices/"
	paymentEndpoint := "https://pec.shaparak.ir/NewIPG/"

	if sandbox == true {
		apiEndPoint = "http://banktest.ir/gateway/Parsian/"
		paymentEndpoint = "http://banktest.ir/gateway/parsian/NewIPG"
	}
	return &Parsian{
		config:          config,
		Transaction:     transaction,
		MerchantID:      MerchantID.(string),
		Sandbox:         sandbox,
		APIEndpoint:     apiEndPoint,
		PaymentEndpoint: paymentEndpoint,
	}, nil
}

type SalePaymentRequest struct {
	LoginAccount   string
	Amount         int64
	OrderId        int64
	CallbackUrl    string
	AdditionalData string
	Originator     string
}

type SalePaymentRequestResp struct {
	XMLName xml.Name
	Body    struct {
		XMLName                    xml.Name
		SalePaymentRequestResponse struct {
			XMLName                  xml.Name
			SalePaymentRequestResult struct {
				Token   string `xml:"Token"`
				Message string `xml:"Message"`
				Status  int    `xml:"Status"`
			} `xml:"SalePaymentRequestResult"`
		} `xml:"SalePaymentRequestResponse"`
	}
}

func (parsian *Parsian) NewTokenRequest(transaction Transaction) (token string, err error) {

	//paymentRequest := SalePaymentRequest{
	//	LoginAccount:   parsian.MerchantID,
	//	Amount:         parsian.Transaction.GetPayableAmount(),
	//	OrderId:        parsian.Transaction.GetGatewayOrderId(),
	//	CallbackUrl:    parsian.Transaction.GetCallbackUrl(),
	//	AdditionalData: parsian.config.AdditionalData.(string),
	//}

	url := parsian.APIEndpoint + "Sale/SaleService.asmx"
	params := SalePaymentRequest{
		LoginAccount: parsian.MerchantID,
		Amount:       parsian.Transaction.GetPayableAmount(),
		OrderId:      parsian.Transaction.GetGatewayOrderId(),
		CallbackUrl:  parsian.Transaction.GetCallbackUrl(),
	}
	SalePaymentRequestPayload := []byte(strings.TrimSpace(fmt.Sprintf(`
    <Envelope xmlns="http://www.w3.org/2003/05/soap-envelope">
        <Body>
            <SalePaymentRequest xmlns="https://pec.Shaparak.ir/NewIPGServices/Sale/SaleService">
                <requestData>
                    <LoginAccount>%s</LoginAccount>
                    <Amount>%d</Amount>
                    <OrderId>%d</OrderId>
                    <CallBackUrl>%s</CallBackUrl>
                    <AdditionalData>%s</AdditionalData>
                    <Originator>%s</Originator>
                </requestData>
            </SalePaymentRequest>
        </Body>
    </Envelope>`,
		params.LoginAccount,
		params.Amount,
		params.OrderId,
		params.CallbackUrl,
		params.AdditionalData,
		params.Originator)))

	httpMethod := "POST"
	soapAction := "urn:SalePaymentRequest" // The format is `urn:<soap_action>`

	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(SalePaymentRequestPayload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return
	}
	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", soapAction)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return
	}

	result := new(SalePaymentRequestResp)
	err = xml.NewDecoder(res.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return
	}

	finalResult := result.Body.SalePaymentRequestResponse.SalePaymentRequestResult

	if finalResult.Status != 0 {
		return "", errors.New(finalResult.Message + " Status: " + string(finalResult.Status))
	}

	return finalResult.Token, nil
}

//type verifyTransactionBody struct {
//	LoginAccount string
//	Token        string
//}
//
//type verifyTransactionResp struct {
//	Status int
//}
//
//type refundTransactionBody struct {
//	LoginAccount string
//	Token        string
//}
//
//type refundTransactionResp struct {
//	Status int
//}
