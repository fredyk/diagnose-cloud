package utils

func MapItems[T any, R any](items []T, f func(T) (R, error)) ([]R, error) {
	mappedItems := make([]R, len(items))
	for i, item := range items {
		t, err := f(item)
		if err != nil {
			return nil, err
		}
		mappedItems[i] = t
	}
	return mappedItems, nil
}
