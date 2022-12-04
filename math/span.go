package math

import "fmt"

// represents the span between two boundaries (boundary values are inclusive)
type Span[T Number] struct {
	From T
	To   T
}

func (s Span[T]) Length() T {
	return Abs(s.To-s.From) + 1
}

func (s Span[T]) Direction() int {
	if s.To >= s.From {
		return 1
	} else {
		return -1
	}
}

func (s Span[T]) GoesForward() bool {
	return s.Direction() == 1
}

// iterate over this span
//
// for v := range s.Iter()
func (s Span[T]) Iter() chan T {
	ch := make(chan T)

	go (func() {
		// direction of the span, as the span is able to go backwards
		forward := s.GoesForward()

		// iterate over span values
		v := s.From
		for (forward && v <= s.To) || (!forward && v >= s.To) {
			ch <- v

			// if-check for addition/subtraction required,
			// as T (Number) could be unsigned and unable to be negative
			if forward {
				v += 1
			} else {
				v -= 1
			}
		}

		close(ch)
	})()

	return ch
}

func (s Span[T]) Contains(n T) bool {
	return s.From <= n && s.To >= n
}

func (s Span[T]) ContainsSpan(o Span[T]) bool {
	return s.Contains(o.From) && s.Contains(o.To)
}

func (s Span[T]) Overlaps(o Span[T]) bool {
	return s.Contains(o.From) || s.Contains(o.To) || o.Contains(s.From) || o.Contains(s.To)
}

func (s Span[T]) Array() []T {
	arr := make([]T, int(s.Length()))

	i := 0
	for v := range s.Iter() {
		arr[i] = v
		i += 1
	}

	return arr
}

func (s Span[T]) String() string {
	return fmt.Sprintf("Span{%v..%v}", s.From, s.To)
}
