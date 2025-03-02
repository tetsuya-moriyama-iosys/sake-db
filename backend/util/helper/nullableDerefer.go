package helper

// NullableDerefer nilの可能性があるものをデリファレンスする
func NullableDerefer[T any](ptr *T) *T {
	if ptr == nil {
		return nil
	}
	return ptr
}
