package utils

type MapFn[T any, R any] func(T) R

type MapWithErrorFn[T any, R any] func(T) (R, error)

func MapItems[T any, R any](items []T, f MapFn[T, R]) []R {
	mappedItems := make([]R, len(items))
	for i, item := range items {
		mappedItems[i] = f(item)
	}
	return mappedItems
}

func MapItemsWithError[T any, R any](items []T, f MapWithErrorFn[T, R]) ([]R, error) {
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
