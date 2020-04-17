package shaparak

import (
	"testing"
	"time"
)

type Transaction struct {
}

func (t Transaction) GetCallbackUrl() string {
	return "http://test.com"
}

func (t Transaction) GetGatewayOrderId() int64 {
	return time.Now().UnixNano()
}

func (t Transaction) GetPayableAmount() int64 {
	return 1000
}

func (t Transaction) GetDescription() string {
	return "desc"
}

func (t Transaction) SetGatewayToken() bool {
	panic("implement me")
}

func (t Transaction) SetReferenceId() bool {
	panic("implement me")
}

func (t Transaction) SetVerified() bool {
	panic("implement me")
}

func (t Transaction) SetSettled() bool {
	panic("implement me")
}

func (t Transaction) SetAccomplished() bool {
	panic("implement me")
}

func (t Transaction) SetRefunded() bool {
	panic("implement me")
}

func (t Transaction) SetCardNumber() bool {
	panic("implement me")
}

func (t Transaction) SetCallBackParameters() bool {
	panic("implement me")
}

func (t Transaction) IsReadyForTokenRequest() bool {
	panic("implement me")
}

func (t Transaction) IsReadyForVerify() bool {
	panic("implement me")
}

func (t Transaction) IsReadyForInquiry() bool {
	panic("implement me")
}

func (t Transaction) IsReadyForSettle() bool {
	panic("implement me")
}

func (t Transaction) IsReadyForRefund() bool {
	panic("implement me")
}

func (t Transaction) AddExtra() bool {
	panic("implement me")
}

func TestParsianInit(t *testing.T) {
	var c map[string]interface{}
	c = make(map[string]interface{})
	c["MerchantID"] = "1u1KRHFvYkHV3TLcgAyv"
	tr := Transaction{}
	_, err := NewParsian(tr, c, true)

	if err != nil {
		t.Error("could not initialize Parsian instance")
	}
}

func TestSalePaymentRequest(t *testing.T) {
	var c map[string]interface{}
	c = make(map[string]interface{})
	c["MerchantID"] = "1u1KRHFvYkHV3TLcgAyv"
	tr := Transaction{}
	p, err := NewParsian(tr, c, true)

	if err != nil {
		t.Error("could not initialize Parsian instance")
	}

	p.NewTokenRequest(tr)
}
