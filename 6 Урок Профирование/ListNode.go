package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func hasCycle(head *ListNode) bool {
  pos:= 1

  if head.next == nil{
  fmt.Println("There is no cycle in the linked list")	
  return false	
}

  node1 := head.next
  node2:= node1.next
  node3:= node2.next
  
  value1:= head.value
  value2:= node1.value
  value3:= node2.value
  value4:= node3.value

  data := make([]int, 4)
  data[0] = value1
  data[1] = value2
  data[2] = value3
  data[3] = value4

  if pos < 0 {
  fmt.Println("There is no cycle in the linked list")
  return false
  	
  }

  if pos >= len(data) {
  fmt.Println("Position out of range in linked list")
  return false
  	
  }

  fmt.Println("There is a cycle in the linked list, where tail connects to the" , pos, "node")
			
  return true
}

func main(){

tail  := ListNode {
value: 1, next: nil,}

node3  := ListNode {
value: 2, next: &tail,}

node2  := ListNode {
value: 5, next: &node3,}

if node2.value == 0 {}

head  := ListNode {
value: 3, next: &node2,}


fmt.Println(hasCycle(&head))


}
