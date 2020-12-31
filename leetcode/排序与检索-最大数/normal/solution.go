package normal

import (
    "sort"
    "strconv"
)

func LargestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        numI := strconv.Itoa(nums[i])
        numJ := strconv.Itoa(nums[j])

        lenI := len(numI)
        lenJ := len(numJ)

        if lenI > lenJ {
            numJ += numI[0:lenI-lenJ+1]
            numI += numJ[0:1]
        } else if lenI < lenJ {
            numI += numJ[0:lenJ-lenI+1]
            numJ += numI[0:1]
        }

        for k := 0; k < len(numI); k ++ {
            if numI[k] > numJ[k] {
                return true
            } else if numI[k] < numJ[k] {
                return false
            }
        }

        return true
    })

    ans := ""
    for _, num := range nums {
        if ans == "" && num == 0 {
            continue
        }
        ans += strconv.Itoa(num)
    }
    if ans == "" {
        return "0"
    }
    return ans
}