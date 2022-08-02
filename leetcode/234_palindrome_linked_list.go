package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ln1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	ln2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	ln3 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	ln4 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  1,
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
						Val: 1,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}

	fmt.Println("Palindrome", isPalindrome(&ln1)) // true
	fmt.Println("Palindrome", isPalindrome(&ln2)) // false
	fmt.Println("Palindrome", isPalindrome(&ln3)) // true
	fmt.Println("Palindrome", isPalindrome(&ln4)) // false
	fmt.Println("Palindrome", isPalindrome(&ln5)) // false
}

func isPalindrome(head *ListNode) bool {
	var prev *ListNode
	var current = head
	var next *ListNode
	count := 0
	for {
		if current == nil {
			break
		}
		current = current.Next
		count++
	}
	if count == 1 {
		return true
	}
	middleLength := count / 2
	current = head
	for {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
		middleLength--
		if current == nil || middleLength == 0 {
			break
		}
	}
	if count%2 == 1 {
		current = current.Next
	}
	for {
		if prev == nil {
			break
		}
		if prev.Val != current.Val {
			return false
		}
		prev = prev.Next
		current = current.Next
	}
	return true
}
