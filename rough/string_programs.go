// rough package contains some practice programs using the data structure concepts

package rough

import "fmt"

// program to replace the value in a string
func Pg1() {
	var word, temp string
	var pos int
	fmt.Println("Enter the string: ")
	fmt.Scan(&word)
here:
	fmt.Println("Length of the string is :", len(word), "\n Enter the position to change:")
	fmt.Scan(&pos)
	if pos > len(word) {
		fmt.Println("Entered postion exceeded the length of the string !!")
		goto here
	}
	fmt.Println("Enter the letter to replace : ")
	fmt.Scan(&temp)
	str := []rune(word)
	str[pos-1] = rune(temp[0])
	word = string(str)
	fmt.Println(word)

}
