package container

import (
	"bytes"
	"fmt"
)

func TestSplit() {
	fmt.Println("Demo split")

	a := []byte("a,b,c")
	b := bytes.Split(a, []byte(","))

	fmt.Printf("a before changing b[0][0]: %q\n", a)
	b[0][0] = '*'
	fmt.Printf("a after changing b[0][0]: %q\n", a)

	fmt.Printf("b[1] before appending to b[0]: %q\n", b[1])
	b[0] = append(b[0], 'd', 'e', 'f')
	fmt.Printf("b[0] after appending: %q\n", b[0])
	fmt.Printf("b[1] after appending to b[0]: %q\n", b[1])

	fmt.Printf("a after appending to b[0]: %q\n", a)
}

func TestAppend() {
	fmt.Println("Demo append")

	s := make([]int, 2, 4)
	s[0] = 1
	s[1] = 2
	fmt.Printf("Initial values: %p --> %[1]v\n", s)

	s = append(s, 3, 4)
	fmt.Printf("After first append: %p --> %[1]v\n", s)

	s = append(s, 5, 6)
	fmt.Printf("After second append: %p --> %[1]v\n", s)
}
