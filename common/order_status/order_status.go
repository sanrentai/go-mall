package order_status

type orderStatus int

const (
	DEFAULT orderStatus = iota - 9
	_
	_
	_
	_
	_
	ORDER_CLOSED_BY_JUDGE
	ORDER_CLOSED_BY_EXPIRED
	ORDER_CLOSED_BY_MALLUSER
	ORDER_PRE_PAY
	ORDER_PAID
	ORDER_PACKAGED
	ORDER_EXPRESS
	ORDER_SUCCESS
)

func (i orderStatus) Name() string {
	switch i {
	case DEFAULT:
		return "ERROR"
	case ORDER_CLOSED_BY_MALLUSER:
		return "手动关闭"
	case ORDER_CLOSED_BY_EXPIRED:
		return "超时关闭"
	case ORDER_CLOSED_BY_JUDGE:
		return "商家关闭"
	case ORDER_PRE_PAY:
		return "待支付"
	case ORDER_PAID:
		return "已支付"
	case ORDER_PACKAGED:
		return "配货完成"
	case ORDER_EXPRESS:
		return "出库成功"
	case ORDER_SUCCESS:
		return "交易成功"
	}
	return "ERROR"
}

func OrderStatus(i int) orderStatus {
	if i > int(ORDER_SUCCESS) || i < int(ORDER_CLOSED_BY_JUDGE) {
		return DEFAULT
	}
	return orderStatus(i)
}
