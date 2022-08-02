package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ln3 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	ln5 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	// fmt.Println("Palindrome", isPalindrome(&ln1)) // true
	// fmt.Println("Palindrome", isPalindrome(&ln2)) // false
	fmt.Println("Middle", middleNode(&ln3).Val) // true
	// fmt.Println("Palindrome", isPalindrome(&ln4)) // false
	fmt.Println("Middle", middleNode(&ln5).Val) // false
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
}
