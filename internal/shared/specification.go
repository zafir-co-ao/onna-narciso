package shared

type Specification[T any] interface {
	IsSatisfiedBy(T) bool
}

type SpecificationFunc[T any] func(T) bool

func (f SpecificationFunc[T]) IsSatisfiedBy(t T) bool {
	return f(t)
}

func And[T any](specs ...Specification[T]) Specification[T] {
	return AndSpecification[T]{Items: specs}
}

func Or[T any](specs ...Specification[T]) Specification[T] {
	return OrSpecification[T]{Items: specs}
}

type AndSpecification[T any] struct {
	Items []Specification[T]
}

func (a AndSpecification[T]) IsSatisfiedBy(t T) bool {
	for _, spec := range a.Items {
		if !spec.IsSatisfiedBy(t) {
			return false
		}
	}
	return true
}

type OrSpecification[T any] struct {
	Items []Specification[T]
}

func (o OrSpecification[T]) IsSatisfiedBy(t T) bool {
	for _, spec := range o.Items {
		if spec.IsSatisfiedBy(t) {
			return true
		}
	}
	return false
}
