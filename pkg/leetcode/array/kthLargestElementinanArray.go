package array

import (
	"container/heap"
	"math/rand"
	"sort"
)

// 215
// sort: o(nlogn)
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// priority queue: time o(nlogk) space o(k)
// 第k个大的数字，所以要大小为k的小根堆，也就是顶是最小的，堆中有k个数，顶部正好是第k个最大的
func findKthLargest2(nums []int, k int) int {
	pq := make(ss, 0)
	for _, v := range nums {
		heap.Push(&pq, v)
		if pq.Len() > k {
			heap.Pop(&pq)
		}
	}
	return pq[0]
}

type ss []int

// 小根堆
func (s ss) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s ss) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ss) Len() int {
	return len(s)
}
func (s *ss) Pop() interface{} {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
func (s *ss) Push(ele interface{}) {
	x := ele.(int)
	*s = append(*s, x)
}

// quick selection 快速选择
// 要先对数组随机洗牌，防止最坏情况n^2
func findKthLargest3(nums []int, k int) int {
	// 升序排列后的第k个最大元素等于从头数第len(nums)-k个元素
	k = len(nums) - k
	// to ensure the o(n), need to shuffle the array first
	shuffle(nums)
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		p := partition(nums, lo, hi)
		if p < k {
			lo = p + 1
		} else if p > k {
			hi = p - 1
		} else {
			return nums[p]
		}
	}
	return -1
}

func partition(nums []int, lo, hi int) int {
	// 不检查的话后面for循环需要检查索引
	if lo == hi {
		return lo
	}
	// 因为先加减，所以索引越界
	i, j := lo, hi+1
	pivot := nums[lo]
	for {
		// 从左开始查找pivot的位置，通过for循环找到比pivot大的
		i++
		for nums[i] < pivot {
			if i == hi {
				break
			}
			i++
		}
		// 从右开始查找pivot的位置，通过for循环找到比pivot小的
		j--
		for nums[j] > pivot {
			if j == lo {
				break
			}
			j--
		}
		// 如果有等于pivot的元素会出现i==j，否则会有i>j
		// 两种情况都应打破循环，因为此时j的左边（不包含j）全部小于pivot而右边全部大于pivot
		if i >= j {
			break
		}
		// 如果走到这里，一定有：
		// nums[i] > pivot && nums[j] < pivot 且i在左j在右
		// 所以需要交换 nums[i] 和 nums[j]，
		// 保证 nums[lo..i] < pivot < nums[j..hi]
		nums[i], nums[j] = nums[j], nums[i]
	}
	// partition完成，交换lo和j的元素，此时j左边全部小于等于pivot，右边大于pivot
	// 返回j，可以对j两边进行排序
	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}

func shuffle(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		// gen random index
		idex := rand.Intn(len(nums) - i)
		// swap that index number with last index number
		nums[len(nums)-1-i], nums[idex] = nums[idex], nums[len(nums)-1-i]
	}
}
