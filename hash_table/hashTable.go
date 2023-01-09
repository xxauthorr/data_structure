package Hashtable

import "fmt"

const arraySize = 50

type Hashtable struct {
	arr [arraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	value string
	key   *bucketNode
}

func (ht *Hashtable) HashTable() {
	var temp int
again:
	fmt.Println("\nSelect the operation: ")
	fmt.Println("1.Insert \n2.Delete \n3.Search \n4.Traverse \n5.Back")
	fmt.Scan(&temp)
	switch temp {
	case 1:
		ht.Insert()
		goto again
	default:
		fmt.Println("Entered a wrong choice !!")
		goto again
	}
}

// function initialize each Hashtable with buckets
func (ht *Hashtable) initHashtable(pos int) {
	ht.arr[pos] = &bucket{}
	fmt.Println("Initialized", pos)
}

// function get the value and returns its hash code using hash algorithm
func hashing(val string) int {
	var count int32
	for i := range val {
		count = rune(val[i]) + count
	}
	return int(count % 100)
}

// function insert the value into hash table
func (ht *Hashtable) Insert() {
	var (
		val      string
		hashCode int
	)

	fmt.Print("Enter the value to insert :")
	fmt.Scan(&val)
	hashCode = hashing(val)
	if hashCode >= arraySize {
		fmt.Println("Exceeded :", hashCode)
		return
	}
	if ht.arr[hashCode] == nil {
		ht.initHashtable(hashCode)
		ht.arr[hashCode].head = &bucketNode{value: val}
		fmt.Println("Value inserted at the head successfully.")
		return
	}
	ht.insertion(*ht.arr[hashCode].head, val, hashCode)
}

func (ht *Hashtable) insertion(b bucketNode, val string, hashCode int) {
	if b.key == nil {
		b.key = &bucketNode{value: val}
		fmt.Println("Value inserted successfully !! ")
		ht.HashTable()
		return
	}
	ht.insertion(*ht.arr[hashCode].head, val, hashCode)
}
