package normal

func Search(nums []int, target int) int {
    ans := -1
    count := len(nums)

    if count == 1 {
        if nums[0] == target {
            return 0
        } else {
            return -1
        }
    }

    from := 0
    to := count - 1

    for from <= to {
        index := (to - from) / 2 + from
        if target == nums[index] {
            ans = index
            break
        }

        if nums[index] >= nums[0] {   // 说明左边都是升序的，右边有拐点
            if target >= nums[0] && target < nums[index] { // 在左边继续寻找
                to = index - 1
            } else {
                from = index + 1
            }
        } else {    // 说明过了拐点，右边都是升序的
            if target <= nums[count-1] && target > nums[index] {
                from = index + 1
            } else {
                to = index - 1
            }
        }
    }

    return ans
}