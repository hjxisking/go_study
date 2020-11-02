package main

import "fmt"

func main() {
    nums := []int{2, 5, 7, 11, 13, 15}
    target := 16

    fmt.Println(twoSum1(nums, target))
    fmt.Println(twoSum2(nums, target))
}

// 这个是暴力法，时间复杂度N^2
func twoSum1(nums []int, target int) []int {
    length := len(nums)
    for key:=0; key<length-1; key++ {
        for keyNext:=key+1; keyNext<length; keyNext++ {
            if nums[key] + nums[keyNext] == target {
                return []int{key, keyNext}
            }
        }
    }
    return nil
}

// hash法，将每个数字作为hash表的key，那么只需要判断target-value是否存在于hash表里就可以了
func twoSum2(nums []int, target int) [] int {
    hash := map[int]int{}
    for key, value := range nums {
        if pre, ok := hash[target-value]; ok {
            return []int{pre, key}
        }
        hash[value] = key
    }

    return nil
}
