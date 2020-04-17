package shaparak

type DriverInterface interface {
	SetParameters()
	GetParameters()
	RefundSupport() bool
	GetForm() string
	GetFormParameters() map[string]string
	GetTransaction()
	VerifyTransaction() bool
	SettleTransaction() bool
	RefundTransaction() bool
	GetGatewayReferenceId() string
	GetUrlFor() string
	CanContinueWithCallbackParameters() bool
	checkRequiredActionParameters()
}

//func GetDriver(driverName string, t Transaction, c ) (driver, err error) {
//	switch driverName {
//	case "parsian":
//		return &parsian{}
//	default:
//		log.Printf("driver undefined")
//		return nil
//	}
//}