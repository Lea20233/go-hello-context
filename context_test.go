package go_hello_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {

	//this is background context
	background := context.Background()
	fmt.Println(background)

	//this is TODO context
	todo := context.TODO()
	fmt.Println(todo)

}

//context with value parent-child

func TestContext_WithValue(t *testing.T) {

	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	//to get context value from parent
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))

	fmt.Println(contextA.Value("b"))
}
