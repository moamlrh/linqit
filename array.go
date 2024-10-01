package linqit

import (
	"reflect"
	"sort"
	"time"
)

func Array[T any](arr []T) Linqit[T] {
	return &linqit[T]{v: append([]T{}, arr...)}
}

func (l *linqit[T]) Where(predicate func(T) bool) Linqit[T] {
	result := make([]T, 0, len(l.v))
	for _, item := range l.v {
		if predicate(item) {
			result = append(result, item)
		}
	}
	l.v = result
	return l
}

func (l *linqit[T]) Select(selector func(T) T) Linqit[T] {
	result := make([]T, len(l.v))
	for i, item := range l.v {
		result[i] = selector(item)
	}
	l.v = result
	return l
}

func (l *linqit[T]) First(predicate func(T) bool) (T, bool) {
	for _, item := range l.v {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

func (l *linqit[T]) FirstOrDefault(predicate func(T) bool) T {
	for _, item := range l.v {
		if predicate(item) {
			return item
		}
	}

	var zero T
	v := reflect.ValueOf(zero)
	if !v.IsValid() {
		return zero
	}
	if v.Kind() == reflect.Ptr {
		return zero
	}

	switch any(zero).(type) {
	case time.Time:
		return any(time.Time{}).(T)
	case []int, []string, []bool, []float64, []interface{}:
		sliceValue := reflect.MakeSlice(reflect.TypeOf(zero), 0, 0)
		return sliceValue.Interface().(T)
	case map[string]interface{}, map[string]string, map[int]string:
		mapValue := reflect.MakeMap(reflect.TypeOf(zero))
		return mapValue.Interface().(T)
	}

	return zero
}

func (l *linqit[T]) Any(predicate func(T) bool) bool {
	for _, item := range l.v {
		if predicate(item) {
			return true
		}
	}
	return false
}

func (l *linqit[T]) All(predicate func(T) bool) bool {
	for _, item := range l.v {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (l *linqit[T]) Count(predicate func(T) bool) int {
	count := 0
	for _, item := range l.v {
		if predicate(item) {
			count++
		}
	}
	return count
}

func (l *linqit[T]) ToSlice() []T {
	return append([]T{}, l.v...)
}

func (l *linqit[T]) Distinct(equals func(T, T) bool) Linqit[T] {
	result := make([]T, 0, len(l.v))
outer:
	for _, item := range l.v {
		for _, existing := range result {
			if equals(item, existing) {
				continue outer
			}
		}
		result = append(result, item)
	}
	l.v = result
	return l
}

func (l *linqit[T]) OrderBy(less func(T, T) bool) Linqit[T] {
	result := append([]T{}, l.v...)
	sort.Slice(l.v, func(i, j int) bool {
		return less(result[i], result[j])
	})
	l.v = result
	return l
}
