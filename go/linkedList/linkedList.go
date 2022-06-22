package main

import (
	"fmt"
	"log"
	"strconv"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) printList() {
	if ll.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	itr := ll.head
	llstring := ""

	for itr != nil {
		llstring += strconv.Itoa(itr.data) + " ---> "
		itr = itr.next
	}

	fmt.Println(llstring)
}

func (ll *LinkedList) getLength() int {
	count := 0

	itr := ll.head
	for itr != nil {
		count += 1
		itr = itr.next
	}
	return count
}

func (ll *LinkedList) insertAtBeginning(data int) {
	temp := &Node{data: data, next: ll.head}
	ll.head = temp
}

func (ll *LinkedList) insertAtEnd(data int) {
	if ll.head == nil {
		ll.head = &Node{data: data, next: nil}
		return
	}

	itr := ll.head
	for itr.next != nil {
		itr = itr.next
	}

	itr.next = &Node{data: data, next: nil}
}

func (ll *LinkedList) insertAt(index int, data int) {
	if index < 0 || index > ll.getLength() {
		log.Fatal("Invalid index")
	}

	if index == 0 {
		temp := &Node{data: data, next: ll.head}
		ll.head = temp
		return
	}

	itr := ll.head
	count := 0

	for itr != nil {
		if count == index-1 {
			temp := &Node{data: data, next: itr.next}
			itr.next = temp
			break
		}
	}
}

func (ll *LinkedList) removeAt(index int) {
	if index < 0 || index >= ll.getLength() {
		log.Fatal("Invalid index")
	}

	if index == 0 {
		ll.head = ll.head.next
		return
	}

	itr := ll.head
	count := 0

	for itr != nil {
		if count == index-1 {
			itr.next = itr.next.next
			break
		}
	}
}

func (ll *LinkedList) insertList(dataList []int) {
	ll.head = nil
	for _, data := range dataList {
		ll.insertAtEnd(data)
	}
}
func main() {
	ll := LinkedList{}
	ll.insertAtBeginning(84)
	ll.insertAtBeginning(100)
	ll.insertAtBeginning(101)
	ll.printList()
	fmt.Println(ll.getLength())
	ll.insertAtEnd(77)
	ll.printList()
	ll.insertAt(1, 22)
	ll.printList()
	ll.removeAt(1)
	ll.printList()
	ll.insertList([]int{1, 2, 3, 4, 5})
	ll.printList()

}
