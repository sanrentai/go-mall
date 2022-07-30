package service_result

type ServiceResult string

const (
	ERROR ServiceResult = "error"

	SUCCESS ServiceResult = "success"

	DATA_NOT_EXIST ServiceResult = "未查询到记录！"

	SAME_CATEGORY_EXIST ServiceResult = "已存在同级同名的分类！"

	SAME_LOGIN_NAME_EXIST ServiceResult = "用户名已存在！"

	LOGIN_NAME_NULL ServiceResult = "请输入登录名！"

	LOGIN_PASSWORD_NULL ServiceResult = "请输入密码！"

	LOGIN_VERIFY_CODE_NULL ServiceResult = "请输入验证码！"

	LOGIN_VERIFY_CODE_ERROR ServiceResult = "验证码错误！"

	SAME_INDEX_CONFIG_EXIST ServiceResult = "已存在相同的首页配置项！"

	GOODS_CATEGORY_ERROR ServiceResult = "分类数据异常！"

	SAME_GOODS_EXIST ServiceResult = "已存在相同的商品信息！"

	GOODS_NOT_EXIST ServiceResult = "商品不存在！"

	GOODS_PUT_DOWN ServiceResult = "商品已下架！"

	SHOPPING_CART_ITEM_LIMIT_NUMBER_ERROR ServiceResult = "超出单个商品的最大购买数量！"

	SHOPPING_CART_ITEM_TOTAL_NUMBER_ERROR ServiceResult = "超出购物车最大容量！"

	LOGIN_ERROR ServiceResult = "登录失败！"

	LOGIN_USER_LOCKED ServiceResult = "用户已被禁止登录！"

	ORDER_NOT_EXIST_ERROR ServiceResult = "订单不存在！"

	ORDER_ITEM_NOT_EXIST_ERROR ServiceResult = "订单项不存在！"

	NULL_ADDRESS_ERROR ServiceResult = "地址不能为空！"

	ORDER_PRICE_ERROR ServiceResult = "订单价格异常！"

	ORDER_GENERATE_ERROR ServiceResult = "生成订单异常！"

	SHOPPING_ITEM_ERROR ServiceResult = "购物车数据异常！"

	SHOPPING_ITEM_COUNT_ERROR ServiceResult = "库存不足！"

	ORDER_STATUS_ERROR ServiceResult = "订单状态异常！"

	OPERATE_ERROR ServiceResult = "操作失败！"

	NO_PERMISSION_ERROR ServiceResult = "无权限！"

	DB_ERROR ServiceResult = "database error"
)
