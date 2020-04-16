package main

import "fmt"


func twoSum(nums []int, target int)(f string){
	sum:= [50]int{}
	for j:= 1; j < len(nums); j++ {
	for i:=0; i<len(nums); i++{
	sum[i] = nums[i] + nums[j]
	if sum[i] == target {
	a1 := i
	a2 := j
	fmt.Println("Введенная сумма масива найдена, индексы дающие сумму в массиве это ",a1, a2)
	return ""
}
}
}

return "Введенная сумма в массиве не найдена"
}

func main() {
	
	nums:= []int{2, 5, 4, 7}
	target := 9
	
	fmt.Println(twoSum(nums, target))
}

