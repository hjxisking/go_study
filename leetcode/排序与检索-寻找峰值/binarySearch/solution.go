package binarySearch

func FindPeakElement(nums []int) int {
    n := len(nums)
    left, right := 0, n - 1
    for left < right {
        mid := left + (right - left + 1) / 2

        if mid == right {
            if nums[mid] < nums[mid - 1] {
                return mid - 1
            } else {
                return mid
            }
        }

        if nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1] {
            return mid
        }

        if nums[mid] < nums[mid + 1] {
            left = mid
        } else {
            right = mid
        }
    }

    return n - 1
}