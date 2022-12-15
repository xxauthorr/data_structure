package main

import (
	"data_structure/linear"
	nonLinear "data_structure/non_linear"
	"fmt"
	"os"
)

func main() {

	// choose the data structure
There:
	res := ChooseStructure()
	switch res {
	case 1:
		lin := TakeLinear()
		switch lin {
		case 1:
			linear.GoToStack()
			goto There
		case 2:
			linear.GoToQueue()
			goto There
		case 3:
			linear.GoToLinkedList()
			goto There
		case 4:
			goto There
		default:
			fmt.Println("Invalid choice !")
			goto There
		}
	case 2:
		nonLin := TakeNonLinear()
		switch nonLin {
		case 1:
			nonLinear.GoToTree()
			goto There
		case 2:
			nonLinear.GoToGraph()
			goto There
		case 4:
			goto There
		default:
			fmt.Println("Invalid choice !")
			goto There
		}
	case 3:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice !")
		goto There
	}

}

func ChooseStructure() int {
	var ans int
	fmt.Println("Choose the data structure to work:")
	fmt.Println("1.Linear \n2.Non-Linear\n3.Exit")
	fmt.Scan(&ans)
	return ans
}

func TakeLinear() int {
	var ans int
	fmt.Println("Choose your module:")
	fmt.Println("1.Stack \n2.Queue \n3.Linked List \n4.Go Back")
	fmt.Scan(&ans)
	return ans
}

func TakeNonLinear() int {
	var ans int
	fmt.Println("Choose your module:")
	fmt.Println("1.Tree \n2.Graph \n3.Go Back ")
	fmt.Scan(&ans)
	return ans
}
