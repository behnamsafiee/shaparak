package shaparak

type TransactionInterface interface {
	GetCallbackUrl() string
	GetGatewayOrderId() int64
	GetPayableAmount() int64
	GetDescription() string

	SetGatewayToken() bool
	SetReferenceId() bool
	SetVerified() bool
	SetSettled() bool
	SetAccomplished() bool
	SetRefunded() bool
	SetCardNumber() bool
	SetCallBackParameters() bool

	IsReadyForTokenRequest() bool
	IsReadyForVerify() bool
	IsReadyForInquiry() bool
	IsReadyForSettle() bool
	IsReadyForRefund() bool

	AddExtra() bool
}