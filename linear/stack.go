package linear

import "fmt"

type stack struct {
	items []int
}

var Stack stack

func GoToStack() {
There:
	var ans int
	fmt.Println("Choose the operation:")
	fmt.Println("1.Push \n2.Pop \n3.Display \n4.Back")
	fmt.Scan(&ans)
	switch ans {
	case 1:
		var val int
		fmt.Println("Enter the value to push")
		fmt.Scan(&val)
		Stack.push(val)
		goto There
	case 2:
		Stack.pop()
		goto There
	case 3:
		Stack.display()
		goto There
	case 4:
		return
	default:
		fmt.Println("Invalid Choice")
		goto There
	}
}

func (s *stack) push(value int) {
	s.items = append(s.items, value)
	fmt.Scan("Enter to continue..")
}
func (s *stack) pop() {

}
func (s *stack) display() {
	fmt.Println(s.items)
}
