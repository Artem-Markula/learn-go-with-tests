package main

import "fmt"


func twoSum(nums []int, target int)(a [2]int) {
	sum:= [50]int{}
	
	for j:= 1; j < len(nums); j++ {
	for i:=0; i<len(nums); i++{
	sum[i] = nums[i] + nums[j]
	if sum[i] == target {
	a[0] = nums[i]
	a[1] = nums[j]
}
}
}

return a
}

func main() {
	
	nums:= []int{2, 5, 4, 7}
	target := 9
	
	fmt.Println(twoSum(nums, target))
}

