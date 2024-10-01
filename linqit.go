package linqit

type Linqit[T any] interface {
	Where(predicate func(T) bool) Linqit[T]
	Select(selector func(T) T) Linqit[T]
	First(predicate func(T) bool) (T, bool)
	FirstOrDefault(predicate func(T) bool) T
	Any(predicate func(T) bool) bool
	All(predicate func(T) bool) bool
	Count(predicate func(T) bool) int
	Distinct(equals func(T, T) bool) Linqit[T]
	OrderBy(less func(T, T) bool) Linqit[T]
	ToSlice() []T
}

type linqit[T any] struct {
	v []T
}
