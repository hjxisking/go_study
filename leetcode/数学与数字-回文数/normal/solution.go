package normal

func IsPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    if x == 0 {
        return true
    }

    nums := make([]int, 0, 20)

    for x != 0 {
        nums = append(nums, x % 10)
        x = x / 10
    }

    n := len(nums)
    i := 0
    j := n - 1

    for i < j {
        if nums[i] != nums[j] {
            return false
        }

        i ++
        j --
    }

    return true
}