package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrinListNode(l *ListNode) {
	i := 0
	for l != nil {
		fmt.Println(i, " = ", l.Val)
		l = l.Next
		i++
	}
}

func ReverseLinkedList(head *ListNode) {
	if head == nil {
		return
	}
	ReverseLinkedList(head.Next)
	fmt.Printf("%d ", head.Val)
}

func makeListNode(arr []int) *ListNode {
	var l *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		l = &ListNode{Val: arr[i], Next: l}
	}
	return l
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		nextTemp := head.Next
		head.Next = prev
		prev = head
		head = nextTemp
	}

	return prev
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := l1
	list2 := l2

	ost := 0

	var ln *ListNode

	for list1 != nil || list2 != nil {

		sum := 0

		if list1 != nil {
			sum += list1.Val
			list1 = list1.Next
		}

		if list2 != nil {
			sum += list2.Val
			list2 = list2.Next
		}

		val := sum + ost

		if val >= 10 {
			ln = &ListNode{Val: val % 10, Next: ln}
			ost = 1
		} else {
			ln = &ListNode{Val: val, Next: ln}
			ost = 0
		}
	}

	if ost > 0 {
		ln = &ListNode{Val: ost, Next: ln}
	}

	ln = reverseList(ln)
	return ln
}

func main() {

	numArr1 := []int{9, 9, 9, 9, 9, 9, 9}
	numArr2 := []int{9, 9, 9, 9}

	// numArr1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	// numArr2 := []int{5, 6, 4}

	var l1 *ListNode
	var l2 *ListNode

	l1 = makeListNode(numArr1)
	l2 = makeListNode(numArr2)

	PrinListNode(addTwoNumbers(l1, l2))
}
