package normal

import (
    "container/heap"
    "sort"
)

// 思路：横向扫描
// 将所有左、右坐标都放入一个一维数组中，挨个循环
// 每个点都循环一次buildings，找到是否左坐标和这个点一样的并放入一个大根堆，同时循环大根堆，去除堆顶的building右节点小于当前点的，最后堆顶如果和结果集不一样，代表是一个新的结果，需要加入
func GetSkyline(buildings [][]int) [][]int {
    ans := [][]int{}
    var bh buildingHeap
    points := make([]int, 0, len(buildings)*2)

    for i := 0; i < len(buildings); i ++ {
        points = append(points, buildings[i][0], buildings[i][1])
    }

    sort.Ints(points)
    points = unique(points)

    for _, point := range points {
        for i := 0; i < len(buildings); i ++ {
            if buildings[i][0] == point {
                heap.Push(&bh, building{
                    left:   buildings[i][0],
                    right:  buildings[i][1],
                    height: buildings[i][2],
                })
            }
        }

        for len(bh) > 0 && bh[0].right <= point {
            heap.Pop(&bh)
        }

        maxHeight := 0
        if len(bh) > 0 {
            maxHeight = bh[0].height
        }

        if len(ans) == 0 || ans[len(ans)-1][1] != maxHeight {
            ans = append(ans, []int{point, maxHeight})
        }
    }

    return ans
}

type building struct {
    left int
    right int
    height int
}

type buildingHeap []building

func (bh buildingHeap) Len() int {
    return len(bh)
}

func (bh buildingHeap) Swap(i, j int) {
    bh[i], bh[j] = bh[j], bh[i]
}

func (bh buildingHeap) Less(i, j int) bool {
    return bh[i].height > bh[j].height
}

func (bh *buildingHeap) Push(b interface{}) {
    *bh = append(*bh, b.(building))
}

func (bh *buildingHeap) Pop() interface{} {
    res := (*bh)[len(*bh)-1]
    *bh = (*bh)[0:len(*bh)-1]
    return res
}

// 有序的数组去重
func unique(nums []int) []int {
    if len(nums) == 1 {
        return nums
    }

    i := 0
    for j := 1; j < len(nums); j ++ {
        if nums[i] != nums[j] {
            i ++
            nums[i] = nums[j]
        } else {
            continue
        }
    }

    return nums[:i+1]
}