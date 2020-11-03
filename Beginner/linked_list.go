package main

import (
	"container/list"
	"fmt"
)

//Node ...
type Node struct {
	Val  int
	Next *Node
}

func main() {
	fmt.Println("Linkedlist Using Library")
	linkedlist := list.New()
	linkedlist.PushBack(1)
	linkedlist.PushBack(2)
	linkedlist.PushBack(3)
	linkedlist.PushBack(4)
	for ele := linkedlist.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele.Value)
	}

	fmt.Println("Linkedlist Using Structs")
	head := Node{1, nil}
	curr := &head
	for i := 1; i < 5; i++ {
		temp := Node{i, nil}
		curr.Next = &temp
		curr := curr.Next
		fmt.Println(curr.Val)
	}
}
