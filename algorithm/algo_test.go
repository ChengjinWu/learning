package algorithm

import (
	"fmt"
	"testing"
	"time"
)

func waysToChange(n int) int {
	count := 0
	var coin25, coin10 int
	for coin25 = 0; coin25 <= n; coin25 += 25 {
		for coin10 = coin25; coin10 <= n; coin10 += 10 {
			count += (n-coin10)/5 + 1
		}
	}
	return count
}
func TestCoins(t *testing.T) {
	start := time.Now()
	waysToChange(750)
	fmt.Println(time.Now().Sub(start))
}

/**
 * Definition for singly-linked list.

 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var n, sum int
	var head *ListNode
	var currNode *ListNode
	for true {
		if l1 == nil && l2 == nil {
			if n != 0 {
				currNode.Next = &ListNode{
					Val: n,
				}
			}
			break
		} else if l1 != nil && l2 == nil {
			sum = l1.Val + n
			n = sum / 10
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			sum = l2.Val + n
			n = sum / 10
			l2 = l2.Next
		} else {
			sum = l1.Val + l2.Val + n
			n = sum / 10
			l1 = l1.Next
			l2 = l2.Next
		}
		node := &ListNode{
			Val: sum % 10,
		}
		if currNode == nil {
			head = node
			currNode = node
		} else {
			currNode.Next = node
			currNode = currNode.Next
		}
	}
	return head
}

func TestAddNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	start := time.Now()
	root := addTwoNumbers(l1, l2)
	for root != nil {
		fmt.Println(root.Val)
		root = root.Next
	}
	fmt.Println(time.Now().Sub(start))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}
	fmt.Println(numMap)
	var n int
	for i := 0; i < len(nums); i++ {
		n = target - nums[i]
		if j, ok := numMap[n]; ok && i != j {

			return []int{i, j}
		}
	}
	return nil
}

func TestAddNumsIndex(t *testing.T) {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
