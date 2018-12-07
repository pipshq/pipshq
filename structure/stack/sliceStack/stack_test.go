package sliceStack_test

import (
	"fmt"

	"github.com/pipshq/structure/stack/sliceStack"
)

func ExampleRemove() {
	stack := sliceStack.New()
	stack.Push(5)
	fmt.Println(stack.Length())
	fmt.Println(stack.Top())
	stack.Push(10)
	stack.Push(1)
	fmt.Println(stack.Items())
	stack.Remove(1)
	fmt.Println(stack.Items())
	// Output:
	// 1
	// 5
	// [5 10 1]
	// [5 1]
}
