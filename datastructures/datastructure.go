package datastructures

import "fmt"

// Datastructure is the common interface that all the datastrucuture should implement
type Datastructure interface {
	Empty() bool
	Size() int
	Elements() []fmt.Stringer
}
