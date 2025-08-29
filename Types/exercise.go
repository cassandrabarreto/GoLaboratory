package main

import "fmt"

func twoSum(nums []int, target int) []int {
    indexes := []int{}
    for i := 0; i < len(nums)-1; i++ {
        current := nums[i]
        current_index := i
        next := nums[i+1] 
        next_index := i + 1 

        sum := current + next

        fmt.Println(sum)
        fmt.Println(target)


        if (sum == target){
            indexes = append(indexes, current_index, next_index)
            break
        }
        continue
        
    }
    return indexes
}
