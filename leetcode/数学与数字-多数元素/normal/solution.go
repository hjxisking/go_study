package normal

func MajorityElement(nums []int) int {
    var candidate int
    vote := 0

    for _, num := range nums {
        if vote == 0 {
            candidate = num
            vote = 1
        } else {
            if num == candidate {
                vote ++
            } else {
                vote --
            }
        }
    }

    return candidate
}