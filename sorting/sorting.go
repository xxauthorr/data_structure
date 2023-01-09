package sorting

import "fmt"

func GetArray() {
	var arrLen, choice, temp int
	var arr []int
	fmt.Println("Enter the length of the array :")
	fmt.Scan(&arrLen)
	fmt.Println("Enter the array elements : ")
	for i := 0; i < arrLen; i++ {
		fmt.Scan(&temp)
		arr = append(arr, temp)
	}
	fmt.Println("Choose the sorting algorithm to work :  \n1.Bubble Sort\n2.Insertion Sort\n3.Selection Sort")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		BubbleSort(arr)
	}
}

func BubbleSort(arr []int) {
	fmt.Println("before: ", arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println("after: ", arr)
	fmt.Println("Enter zero to go back..")
	fmt.Scan(&arr[0])
	GetArray()
}

func InsertionSort(arr []int){

}

func SelectionSort(arr []int){
	
}
