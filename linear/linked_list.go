package linear

import "fmt"

type node struct {
	val  int
	next *node
}

type linkedList struct {
	head *node
	len  int
}

var lList linkedList

func GoToLinkedList() {
There:
	var ans int

	fmt.Println("\nChoose the operation:")
	fmt.Println("1.Prepend \n2.Delete \n3.Display the head \n4.traversal \n5.Back")
	fmt.Scan(&ans)
	switch ans {
	case 1:
		var val int
		fmt.Println("Enter the value to prepend")
		fmt.Scan(&val)
		lList.prepend(node{val: val})
		goto There
	case 2:
		var val int
		fmt.Println("Enter the value to delete:")
		fmt.Scan(&val)
		lList.delete(val, *lList.head)
		goto There
	case 3:
		lList.display()
		goto There
	case 4:
		lList.traversal(*lList.head)
		goto There
	case 5:
		return
	default:
		fmt.Println("Invalid Choice")
		goto There
	}

}

func (l *linkedList) prepend(n node) {
	temp := l.head
	l.head = &n
	l.head.next = temp
	l.len++
}

func (l *linkedList) delete(val int, n node) {
	if l.len == 0 {
		fmt.Println("The linked list is empty")
		return
	}
	if l.head.val == val {
		l.head = l.head.next
		l.len--
		fmt.Println("Value deleted successfully")
		return
	}
	if l.head.next != nil {
		l.delRepeat(l.head, val)
	} else {
		fmt.Println("The value doesn't exist")
	}
}

func (l *linkedList) delRepeat(n *node, val int) {
	if n.next.val == val {
		n.next = n.next.next
		l.len--
		fmt.Println("Value deleted successfully")
		return
	}
	if n.next.next != nil {
		l.delRepeat(n.next, val)
	} else {
		fmt.Println("The value doesn't exist")
	}
}

func (l *linkedList) traversal(n node) {
	if l.head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	fmt.Println()
	fmt.Print(l.head.val, " ")
	if l.head.next != nil {
		l.addRepeat(l.head.next)
		fmt.Println()
	}
}
func (l *linkedList) addRepeat(n *node) {
	fmt.Print(n.val, " ")
	if n.next != nil {
		l.addRepeat(n.next)
	}
}

func (l *linkedList) display() {
	if l.head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	fmt.Println("Head value:", l.head.val)
	if l.head.next == nil {
		fmt.Print("Next Address :nil")
		return
	}
	fmt.Println("Next Address :", &l.head.next)
}
