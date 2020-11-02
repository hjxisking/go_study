package normal

import "sort"

/**
思路：
    1. 大方向：3层循环
    2. 为了避免重复，首先需要排序，其次无论first或是second，都不能和上次重复，比如三元组如果是1开头的，那么下个三元组就不能1开头，不然肯定重复，同理三元组如果是2，3开头的，那么second就不能再是3了，不然third也必然重复
    3. 关于third的循环其实可以和second同步，也就是在second自增的同时，third先递减到最接近second的位置，同时三元组之和大于目标值，在second下次自增的时候third继续接近且不需要重置到尾部来接近（second大了，third就要减小）
 */

func ThreeSum(nums []int) [][]int {
    sort.Sort(sort.IntSlice(nums))
    n := len(nums)
    ans := make([][]int, 0)

    for first := 0; first < n; first ++ {
        if first > 0 && nums[first] == nums[first-1] {      // 如果以相同的数字开头，那么结果一定会重复，因为最后的三元组是有序的
            continue
        }

        third := n - 1
        for second := first + 1; second < n; second ++ {
            if second > first + 1 && nums[second] == nums[second-1] {   // 理由同上
                continue
            }

            // first、second确定的前提下，third从尾部开始向前走，总和大于目标则继续向前走（越向前越小），直到最接近second，但必须比second大
            // second每多进入一次循环，third是有可能继续向second靠近的，因为second变大了，总和不变，third必然要更小
            for third > second && nums[first] + nums[second] + nums[third] > 0 {
                third --
            }

            // third是不会重置的，当second进入下一个循环的时候，third是不会变的，直到second++追上third而终止
            if second == third {
                break
            }

            if nums[first] + nums[second] + nums[third] == 0 {
                ans = append(ans, append([]int{}, nums[first], nums[second], nums[third]))
            }
        }
    }

    return ans
}
