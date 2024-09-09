package helper

// IsEmpty はポインタ型 T を受け取り、nil または空文字列の場合に true を返します。
func IsEmpty[T any](val *T) bool {
	if val == nil {
		return true
	}

	switch v := any(*val).(type) {
	case string:
		return v == ""
	default:
		return false
	}
}
