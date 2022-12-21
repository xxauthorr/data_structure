package linear

import "fmt"


type stack struct {
	items []int
}

var Stack stack

func GoToStack() {
There:
	var ans int

	fmt.Println("\nChoose the operation:")
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
	fmt.Println("Value pushed successfully")
}
func (s *stack) pop() {
	if len(s.items) == 0 {
		fmt.Println("The stack is empty !")
		return
	}
	toRemove := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	fmt.Println("pop ", toRemove)
	fmt.Println("Value poped successfully")
}
func (s *stack) display() {
	fmt.Println("Stack: ", s.items)
}
