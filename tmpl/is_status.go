package tmpl

func IsStatus(status, is, value string, params ...string) string {
	if status == is {
		return value
	}
	if len(params) > 0 {
		return params[0]
	}
	return ""
}
