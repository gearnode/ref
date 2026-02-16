package ref

func Ref[T any](v T) *T {
	return &v
}

func Unref[T any](v *T) T {
	return *v
}

func RefOrNil[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}
	return &v
}

func UnrefOrZero[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}
