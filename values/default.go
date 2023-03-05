// Package values handles helpers around managing/setting/getting
// values consistently
package values

// IfNotSet will return the value or the default if the value is the value's default setting
func IfNotSet[T comparable](val, defaultTo T) T {
	if val == *new(T) {
		return defaultTo
	}
	return val
}
