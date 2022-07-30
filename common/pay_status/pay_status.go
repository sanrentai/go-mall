package pay_status

type payStatus int

const (
	DEFAULT payStatus = iota - 1
	PAY_ING
	PAY_SUCCESS
)

func PayStatus(i int) payStatus {
	if i > int(PAY_SUCCESS) || i < int(DEFAULT) {
		return DEFAULT
	}
	return payStatus(i)
}

func (s payStatus) Name() string {
	return [...]string{"支付失败", "支付中", "支付成功"}[s+1]
}
