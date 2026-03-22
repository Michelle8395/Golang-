package main

import "fmt"

func main() {
	/*
		[1,2,3],[4,5,6] = [3,6,2,5,1,4]
		[1,2,3],[] = [3,2,1]
		[],[4,5,6] = [6,5,4]
		[3,5,7,8,9],[2,4,6] = [9,8,7,6,5,4,3]
	*/
	

	fmt.Println(RevConcatSlice([]int{1,2,3},[]int{4,5,6}))
	fmt.Println(RevConcatSlice([]int{1,2,3},[]int{}))
	fmt.Println(RevConcatSlice([]int{},[]int{4,5,6}))
	fmt.Println(RevConcatSlice([]int{3,5,7,8,9},[]int{2,4,6}))

}

func RevConcatSlice(s1, s2 []int) (result []int) {
	if len (s2) > len(s1){
		s1,s2 = s2,s1
	}
	for i := len(s1) - 1; i >= 0; i-- {
		result = append(result, s1[i])
		for j := i; j < len(s2); j-- {
			result = append(result, s2[j])
			break
		}
	}
	return
}
