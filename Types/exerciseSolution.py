class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        indexes = []
        for i in range(len(nums)):
            current_number = nums[i]
            current_index = i
            next_number = nums[i+1]
            next_index = i + 1 

            sum = current_number + next_number

            if sum == target:
                indexes.append(current_index)
                indexes.append(next_index)
                break
            else:
                continue
        return indexes

"""
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

"""