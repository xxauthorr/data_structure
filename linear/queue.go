package linear

import "fmt"

type Queue struct {
	items []int
}

var queue Queue

func GoToQueue() {
There:
	var ans int

	fmt.Println("\nChoose the operation:")
	fmt.Println("1.Enqueue \n2.Dequeue \n3.Display \n4.Back")
	fmt.Scan(&ans)
	switch ans {
	case 1:
		var val int
		fmt.Println("Enter the value to enqueue")
		fmt.Scan(&val)
		queue.Enqueue(val)
		goto There
	case 2:
		queue.Dequeue()
		goto There
	case 3:
		queue.display()
		goto There
	case 4:
		return
	default:
		fmt.Println("Invalid Choice")
		goto There
	}
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
	fmt.Println("Value enqued successfully")
}

func (q *Queue) Dequeue() {
	if len(q.items) == 0 {
		fmt.Println("The queue is empty")
		return
	}
	toRemove := q.items[0]
	q.items = q.items[1:]
	fmt.Println("Dequeued ", toRemove)
	fmt.Println("value dequeued successfully")
}

func (q *Queue) display() {
	fmt.Println("Queue: ", q.items)
}
