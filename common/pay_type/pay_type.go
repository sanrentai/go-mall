package pay_type

type payType int

const (
	DEFAULT payType = iota - 1
	NOT_PAY
	ALI_PAY
	WEIXIN_PAY
)

func PayType(i int) payType {
	if i > int(WEIXIN_PAY) || i < int(DEFAULT) {
		return DEFAULT
	}
	return payType(i)
}

func (s payType) Name() string {
	return [...]string{"ERROR", "无", "支付宝", "微信支付"}[s+1]
}
