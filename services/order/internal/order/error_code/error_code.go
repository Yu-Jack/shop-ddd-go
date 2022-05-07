package error_code

var (
	errMap map[string]string = map[string]string{
		"400": "Request body is invalid.",
		"401": "Request uri is invalid.",
		"500": "Order is not available to checkout",
	}

	REQUEST_BODY_IS_INVALID = "400"
	REQUEST_URI_IS_INVALID  = "401"
	ORDER_IS_NOT_AVAILABLE  = "500"
)

func New(code string) map[string]string {
	v := errMap[code]
	return map[string]string{
		"code":    code,
		"message": v,
	}
}
