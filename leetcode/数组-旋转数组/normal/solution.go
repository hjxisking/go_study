package normal

func Rotate(nums []int, k int)  {
    n := len(nums)
    if n > 1 {
        start := 0
        from := 0
        exchange := nums[0]
        for i := 0; i < n; i ++ {
            to := (from + k) % n
            tmp := nums[to]
            nums[to] = exchange

            if to == start {    // nums偶数时是会产生循环的
                from = (from + 1) % n
                exchange = nums[from]
                start = from
            } else {
                from = to
                exchange = tmp
            }
        }
    }
}
